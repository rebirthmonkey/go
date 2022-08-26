// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ AuthStrategy = &JWTStrategy{}

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}
