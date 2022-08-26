// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/component-base/pkg/core"
	"github.com/marmotedu/errors"
	"github.com/rebirthmonkey/iam/internal/pkg/errcode"
	"github.com/rebirthmonkey/iam/pkg/middleware"
)

// Defined errors.
var (
	ErrMissingKID    = errors.New("Invalid token format: missing kid field in claims")
	ErrMissingSecret = errors.New("Can not obtain secret information from cache")
)

// Secret contains the basic information of the secret key.
type Secret struct {
	Username string
	ID       string
	Key      string
	Expires  int64
}

// CacheStrategy defines jwt bearer authentication strategy which called `cache strategy`.
// GetPolicyRepo are obtained through grpc api interface and cached in memory.
type CacheStrategy struct {
	getKeyFunc func(kid string) (Secret, error)
}

var _ AuthStrategy = &CacheStrategy{}

// NewCacheStrategy create cache strategy with function which can list and cache secrets.
func NewCacheStrategy(get func(kid string) (Secret, error)) CacheStrategy {
	return CacheStrategy{get}
}

// AuthFunc defines cache strategy as the gin authentication middleware.
func (cache CacheStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if len(header) == 0 {
			core.WriteResponse(
				c,
				errors.WithCode(errcode.ErrMissingHeader, "Authorization header cannot be empty."),
				nil,
			)
			c.Abort()

			return
		}

		var rawJWT string
		// Parse the header to getKeyFunc the token part.
		fmt.Sscanf(header, "Bearer %s", &rawJWT)

		// Use own validation logic, see below
		var secret Secret

		claims := &jwt.MapClaims{}
		// Verify the token
		parsedT, err := jwt.ParseWithClaims(rawJWT, claims, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is HMAC signature
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, ErrMissingKID
			}

			var err error
			secret, err = cache.getKeyFunc(kid)
			if err != nil {
				return nil, ErrMissingSecret
			}

			return []byte(secret.Key), nil
		}, jwt.WithAudience(AuthzAudience))
		if err != nil || !parsedT.Valid {
			core.WriteResponse(c, errors.WithCode(errcode.ErrSignatureInvalid, err.Error()), nil)
			c.Abort()

			return
		}

		if KeyExpired(secret.Expires) {
			tm := time.Unix(secret.Expires, 0).Format("2006-01-02 15:04:05")
			core.WriteResponse(c, errors.WithCode(errcode.ErrExpired, "expired at: %s", tm), nil)
			c.Abort()

			return
		}

		c.Set(middleware.UsernameKey, secret.Username)
		c.Next()
	}
}

// KeyExpired checks if a key has expired, if the value of user.SessionState.Expires is 0, it will be ignored.
func KeyExpired(expires int64) bool {
	if expires >= 1 {
		return time.Now().After(time.Unix(expires, 0))
	}

	return false
}
