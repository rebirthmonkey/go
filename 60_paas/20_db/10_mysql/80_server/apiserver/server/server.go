package server

import (
	"fmt"

	"github.com/rebirthmonkey/go/pkg/gin"
	"github.com/rebirthmonkey/go/pkg/grpc"
	"github.com/rebirthmonkey/go/pkg/mysql"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	mysqlServer *mysql.Server
	ginServer   *gin.Server
	grpcServer  *grpc.Server
}

type PreparedServer struct {
	preparedMysqlServer *mysql.PreparedServer
	preparedGinServer   *gin.PreparedServer
	preparedGrpcServer  *grpc.PreparedServer
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
		preparedMysqlServer: s.mysqlServer.PrepareRun(),
		preparedGinServer:   s.ginServer.PrepareRun(),
		preparedGrpcServer:  s.grpcServer.PrepareRun(),
	}
}

func (s PreparedServer) Run() error {
	fmt.Println("[PreparedServer] Run")

	var eg errgroup.Group

	eg.Go(func() error {
		s.preparedMysqlServer.Run()

		return nil
	})

	eg.Go(func() error {
		s.preparedGinServer.Run()

		return nil
	})

	return s.preparedGrpcServer.Run()
}
