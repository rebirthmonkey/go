// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"
)

const (
	UsernameKey   = "username"
	AuthzAudience = "authz.iam.rebirthmonkey.com" // AuthzAudience defines the value of jwt audience field.

)

// AuthStrategy defines the set of methods used to do resource authentication.
type AuthStrategy interface {
	AuthFunc() gin.HandlerFunc
}

// AuthOperator used to switch between different authentication strategy.
type AuthOperator struct {
	strategy AuthStrategy
}

// AuthFunc execute resource authentication.
func (operator *AuthOperator) AuthFunc() gin.HandlerFunc {
	return operator.strategy.AuthFunc()
}

// SetStrategy used to set to another authentication strategy.
func (operator *AuthOperator) SetStrategy(strategy AuthStrategy) {
	operator.strategy = strategy
}
