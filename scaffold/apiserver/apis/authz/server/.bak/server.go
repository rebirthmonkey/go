// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/log"
	"golang.org/x/sync/errgroup"
)

// Server is the running structure of the server.
type Server struct {
	ginServer  *gin.Server
	grpcServer *grpc.Server
}

// PreparedServer is the running structure of the initialized server.
type PreparedServer struct {
	preparedGinServer  *gin.PreparedServer
	preparedGrpcServer *grpc.PreparedServer
}

// NewServer creates a running server instance
func NewServer(opts *Options) (*Server, error) {
	config := NewConfig()

	if err := opts.ApplyTo(config); err != nil {
		return nil, err
	}

	serverInstance, err := config.Complete().New()
	if err != nil {
		return nil, err
	}

	return serverInstance, nil
}

// PrepareRun creates a running server instance after complete initialization.
func (s *Server) PrepareRun() PreparedServer {
	log.Info("[Server] PrepareRun")

	InitGin(s.ginServer.Engine)
	InitGrpc(s.grpcServer.Server)

	return PreparedServer{
		preparedGinServer:  s.ginServer.PrepareRun(),
		preparedGrpcServer: s.grpcServer.PrepareRun(),
	}
}

// Run launches the prepared server instance.
func (s PreparedServer) Run() error {
	log.Info("[PreparedServer] Run")

	var eg errgroup.Group

	eg.Go(func() error {
		if err := s.preparedGinServer.Run(); err != nil {
			return err
		}

		return nil
	})

	//eg.Go(func() error {
	//	if err := s.preparedGrpcServer.Run(); err != nil {
	//		return err
	//	}
	//
	//	return nil
	//})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
