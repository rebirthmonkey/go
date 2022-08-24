package grpc

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/controller/grpc/v1"
	userRepoFake "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo/fake"
)

func Init(server *grpc.Server) {
	installController(server)
}

func installController(server *grpc.Server) {
	fmt.Println("[GrpcServer] registry userController")

	userRepoClient, _ := userRepoFake.Repo()
	userController := v1.NewController(userRepoClient)
	v1.RegisterUserServer(server, userController)
}
