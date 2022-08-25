package server

import (
	"fmt"

	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/controller/grpc/v1"
	userRepoMysql "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/repo/mysql"
)

func InitGrpc(server *grpc.Server) {
	fmt.Println("[GrpcServer] registry userController")

	//userRepoClient, _ := userRepoFake.Repo()
	userRepoClient, _ := userRepoMysql.Repo(config.CompletedMysqlConfig)

	userController := v1.NewController(userRepoClient)
	v1.RegisterUserServer(server, userController)
}
