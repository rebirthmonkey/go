package main

import (
	"github.com/gin-gonic/gin"
)

// AuthStrategy defines the set of methods used to do resource authentication.
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// authOperator used to switch between different authentication strategy.
type authOperator struct {
	strategy AuthStrategy
}

// AuthFunc execute resource authentication.
func (operator *authOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}

// SetStrategy used to set to another authentication strategy.
func (operator *authOperator) SetStrategy(strategy AuthStrategy) {
	operator.strategy = strategy
}
