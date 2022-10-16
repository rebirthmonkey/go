package main

import (
	"github.com/gin-gonic/gin"
)

// AuthStrategy defines the set of methods used to do authentication.
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// authOperator used to switch between different authentication strategy.
type authOperator struct {
	strategy AuthStrategy
}

// AuthFunc executes the set authentication.
func (operator *authOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}

// SetStrategy used to set to an authentication strategy.
func (operator *authOperator) SetStrategy(strategy AuthStrategy) {
	operator.strategy = strategy
}
