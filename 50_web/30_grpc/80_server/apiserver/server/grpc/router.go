package grpc

import (
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc"
)

func Init(server *grpc.Server) {

	productInfoHandler := &productInfoHandler{}

	pb.RegisterProductInfoServer(server, productInfoHandler)
}
