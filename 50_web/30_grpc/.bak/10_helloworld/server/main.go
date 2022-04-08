package main

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/rebirthmonkey/pkg/grpc/productinfo"
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
	productMap map[string]*productinfo.Product
}

//添加商品
func (s *server) AddProduct(ctx context.Context, req *productinfo.Product) (resp *productinfo.ProductId, err error) {
	resp = &productinfo.ProductId{}
	out, err := uuid.NewV4()
	if err != nil {
		return resp, status.Errorf(codes.Internal, "err while generate the uuid ", err)
	}

	req.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*productinfo.Product)
	}

	s.productMap[req.Id] = req
	resp.Value = req.Id
	return
}

//获取商品
func (s *server) GetProduct(ctx context.Context, req *productinfo.ProductId) (resp *productinfo.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*productinfo.Product)
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
	productinfo.RegisterProductInfoServer(s, &server{})
	log.Println("start gRPC listen on port", address)
	if err := s.Serve(listener); err != nil {
		log.Println("failed to serve...", err)
		return
	}
}

