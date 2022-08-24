package grpc

import (
	"github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/controller/grpc/v1"
	"google.golang.org/grpc"

	userRepoFake "github.com/rebirthmonkey/go/50_web/30_grpc/80_server/apiserver/user/repo/fake"
)

func Init(server *grpc.Server) {

	//productInfoHandler := &productInfoHandler{}
	//pb.RegisterProductInfoServer(server, productInfoHandler)

	userRepoClient, _ := userRepoFake.NewRepo()
	userController := v1.NewController(userRepoClient)
	v1.RegisterUserServer(server, userController)
}
