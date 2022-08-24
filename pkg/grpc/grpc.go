package grpc

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net"

	"github.com/rebirthmonkey/pkg/log"
	"google.golang.org/grpc"
)

type Server struct {
	Address string

	*grpc.Server
}

type PreparedServer struct {
	*Server
}

func (s *Server) PrepareRun() *PreparedServer {
	fmt.Println("[GrpcServer] PrepareRun")

	return &PreparedServer{s}
}

func (s *PreparedServer) Run() error {
	fmt.Println("[GrpcServer] Run")

	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		fmt.Printf("failed to listen: %s", err.Error())
	}

	var eg errgroup.Group

	eg.Go(func() error {
		fmt.Printf("[GrpcServer] Start to listening on http address: %s", s.Address)

		if err := s.Serve(listen); err != nil {
			fmt.Printf("failed to start grpc productInfoHandler: %s", err.Error())

			return err
		}

		fmt.Printf("Server on %s stopped", s.Address)

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *Server) init() {
	fmt.Println("[GrpcServer] Init")
}
