// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const authHeaderCount = 2

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
// according `Authorization` header.
type AutoStrategy struct {
	basic AuthStrategy
	jwt   AuthStrategy
}

var _ AuthStrategy = &AutoStrategy{}

// NewAutoStrategy create auto strategy with basic strategy and jwt strategy.
func NewAutoStrategy(basic, jwt AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

// AuthFunc defines auto strategy as the gin authentication middleware.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		operator := AuthOperator{}
		authHeader := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != authHeaderCount {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("authorization header format is wrong")})
			c.Abort()

			return
		}

		switch authHeader[0] {
		case "Basic":
			operator.SetStrategy(a.basic)
		case "Bearer":
			operator.SetStrategy(a.jwt)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("unrecognized Authorization header")})

			c.Abort()

			return
		}
		operator.AuthFunc()(c)
		c.Next()
	}
}
