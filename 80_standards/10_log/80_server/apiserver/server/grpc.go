package server

import (
	"github.com/rebirthmonkey/go/pkg/log"
	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/80_standards/10_log/80_server/apiserver/user/controller/grpc/v1"
	userRepoMysql "github.com/rebirthmonkey/go/80_standards/10_log/80_server/apiserver/user/repo/mysql"
)

func InitGrpc(server *grpc.Server) {
	log.Info("[GrpcServer] registry userController")

	//userRepoClient, _ := userRepoFake.Repo()
	userRepoClient, _ := userRepoMysql.Repo(config.CompletedMysqlConfig)

	userController := v1.NewController(userRepoClient)
	v1.RegisterUserServer(server, userController)
}
