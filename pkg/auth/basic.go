// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/log"
)

type BasicStrategy struct {
	compare func(username string, password string) bool
}

var _ AuthStrategy = &BasicStrategy{}

func NewBasicStrategy(compare func(username string, password string) bool) BasicStrategy {
	return BasicStrategy{
		compare: compare,
	}
}

func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("[BasicStrategy] Authentication")
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not basic auth")})
			c.Abort()
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("username or password not matched")})
			c.Abort()
			return
		}

		c.Set(UsernameKey, pair[0])
		c.Next()
	}
}
