package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/go/pkg/errcode"
	"github.com/rebirthmonkey/go/pkg/errors"
	"github.com/rebirthmonkey/go/pkg/gin/util"
	"github.com/rebirthmonkey/go/pkg/log"
)

func pingHandler(ctx *gin.Context) {
	log.L(ctx).Info("pingHandler() called.")

	// Simulate an external error.
	err := errors.Errorf("auth error")
	if err != nil {
		util.WriteResponse(ctx, errors.WithCode(errcode.ErrInvalidAuthHeader, err.Error()), nil)
		return
	}
	util.WriteResponse(ctx, nil, "pong")
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/ping", pingHandler)
	ginEngine.Run(":8080")
}
