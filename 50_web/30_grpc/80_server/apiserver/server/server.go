package server

import (
	"fmt"
	ginInstance "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/server/gin"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/gin"
	"golang.org/x/sync/errgroup"

	grpcInstance "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/server/grpc"
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/pkg/grpc"
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

	ginInstance.Init(s.ginServer.Engine)
	grpcInstance.Init(s.grpcServer.Server)

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