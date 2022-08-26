package grpc

import (
	"golang.org/x/sync/errgroup"
	"net"

	"github.com/rebirthmonkey/go/pkg/log"
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
	log.Info("[GrpcServer] PrepareRun")

	return &PreparedServer{s}
}

func (s *PreparedServer) Run() error {
	log.Info("[GrpcServer] Run")

	listen, err := net.Listen("tcp", s.Address)
	if err != nil {
		log.Infof("failed to listen: %s", err.Error())
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("[GrpcServer] Start to listening on http address: %s", s.Address)

		if err := s.Serve(listen); err != nil {
			log.Infof("failed to start grpc productInfoHandler: %s", err.Error())

			return err
		}

		log.Infof("Server on %s stopped", s.Address)

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}

func (s *Server) init() {
	log.Info("[GrpcServer] Init")
}
