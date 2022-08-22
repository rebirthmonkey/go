package grpc

import (
	"context"

	"github.com/gofrs/uuid"
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productInfoHandler struct {
	productMap map[string]*pb.Product
	pb.UnimplementedProductInfoServer
}

//添加商品
func (s *productInfoHandler) AddProduct(ctx context.Context, req *pb.Product) (resp *pb.ProductId, err error) {
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
func (s *productInfoHandler) GetProduct(ctx context.Context, req *pb.ProductId) (resp *pb.Product, err error) {
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	resp = s.productMap[req.Value]
	return
}
