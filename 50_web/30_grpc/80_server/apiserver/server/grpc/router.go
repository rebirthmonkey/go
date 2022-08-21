package grpc

import (
	"google.golang.org/grpc"

	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
)

func Init(server *grpc.Server) {

	productInfoHandler := &productInfoHandler{}
	
	pb.RegisterProductInfoServer(server, productInfoHandler)
}
