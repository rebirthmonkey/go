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

type Server struct {
	ginServer  *gin.Server
	grpcServer *grpc.Server
}

type PreparedServer struct {
	preparedGinServer  *gin.PreparedServer
	preparedGrpcServer *grpc.PreparedServer
}

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

func (s *Server) PrepareRun() PreparedServer {
	log.Info("[Server] PrepareRun")

	InitGin(s.ginServer.Engine)
	InitGrpc(s.grpcServer.Server)

	return PreparedServer{
		preparedGinServer:  s.ginServer.PrepareRun(),
		preparedGrpcServer: s.grpcServer.PrepareRun(),
	}
}

func (s PreparedServer) Run() error {
	log.Info("[PreparedServer] Run")

	var eg errgroup.Group

	eg.Go(func() error {
		if err := s.preparedGinServer.Run(); err != nil {
			return err
		}

		return nil
	})

	eg.Go(func() error {
		if err := s.preparedGrpcServer.Run(); err != nil {
			return err
		}

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
