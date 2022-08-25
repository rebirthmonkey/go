package server

import (
	"fmt"

	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
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
	opts.ApplyTo(config)
	serverInstance, err := config.Complete().New()

	return serverInstance, err
}

func (s *Server) PrepareRun() PreparedServer {
	fmt.Println("[Server] PrepareRun")

	InitGin(s.ginServer.Engine)
	InitGrpc(s.grpcServer.Server)

	return PreparedServer{
		preparedGinServer:  s.ginServer.PrepareRun(),
		preparedGrpcServer: s.grpcServer.PrepareRun(),
	}
}

func (s PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	var eg errgroup.Group

	eg.Go(func() error {
		s.preparedGinServer.Run()

		return nil
	})

	return s.preparedGrpcServer.Run()
}
