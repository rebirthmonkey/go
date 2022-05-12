package main

import (
	"context"
	"github.com/gofrs/uuid"
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

func (s *server) AddProduct(ctx context.Context, req *pb.Product) (resp *pb.ProductId, err error) {
	if req.Id == "" {
		out, _ := uuid.NewV4()
		req.Id = out.String()
	}

	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[req.Id] = req

	resp = &pb.ProductId{}
	resp.Value = req.Id
	err = nil
	return
}

func (s *server) GetProduct(ctx context.Context, req *pb.ProductId) (resp *pb.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	resp = s.productMap[req.Value]
	err = nil
	return
}

func main() {
	s := grpc.NewServer()

	pb.RegisterProductInfoServer(s, &server{})

	listener, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Println("net listen err ", err)
		return
	}
	log.Println("start gRPC listen on port:", "127.0.0.1:50051")
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}
