package main

import (
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"golang.org/x/sync/errgroup"
	"net"

	"google.golang.org/grpc"

	"github.com/rebirthmonkey/pkg/log"
)

type GRPCServer struct {
	Address string
	*grpc.Server
	*productInfoHandler
}

func CreateGRPCServer(cfg *grpcConfig) (*GRPCServer, error) {

	grpcServer := grpc.NewServer()
	productInfoHandler := &productInfoHandler{}

	return &GRPCServer{cfg.Address, grpcServer, productInfoHandler}, nil
}

func (grpcServer *GRPCServer) Init () error {
	log.Info("[GRPC Server] registry productInfoHandler")
	pb.RegisterProductInfoServer(grpcServer.Server, grpcServer.productInfoHandler)

	return nil
}

func (grpcServer *GRPCServer) Run() error {
	listen, err := net.Listen("tcp", grpcServer.Address)
	if err != nil {
		log.Fatalf("failed to listen: %s", err.Error())
	}

	var eg errgroup.Group

	eg.Go(func() error {
		log.Infof("[GRPC Server] Start to listening on http address: %s", grpcServer.Address)

		//if err := productInfoHandler.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		if err := grpcServer.Serve(listen); err != nil  {
			log.Fatalf("failed to start grpc productInfoHandler: %s", err.Error())

			return err
		}

		log.Infof("Server on %s stopped", grpcServer.Address)

		return nil
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
