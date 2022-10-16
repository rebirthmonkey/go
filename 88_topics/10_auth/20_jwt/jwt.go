package main

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const UsernameKey = "username"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

var _ AuthStrategy = &JWTStrategy{}

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

func (strategy JWTStrategy) AuthFunc() gin.HandlerFunc {
	return strategy.MiddlewareFunc()
}
