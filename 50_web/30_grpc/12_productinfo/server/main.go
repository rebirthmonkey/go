package main

import (
	"context"
	"github.com/gofrs/uuid"
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	address = "127.0.0.1:50051"
)

type server struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

//添加商品
func (s *server) AddProduct(ctx context.Context, req *pb.Product) (resp *pb.ProductId, err error) {
	resp = &pb.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *server) GetProduct(ctx context.Context, req *pb.ProductId) (resp *pb.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	resp = s.productMap[req.Value]
	return
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Println("net listen err ", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Println("start gRPC listen on port", address)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

