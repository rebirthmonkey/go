// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/rebirthmonkey/go/pkg/log"
	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/controller/grpc/v1"
	userRepoMysql "github.com/rebirthmonkey/go/88_topics/10_auth/80_server/apiserver/user/repo/mysql"
)

// InitGrpc initializes the Grpc server
func InitGrpc(server *grpc.Server) {
	log.Info("[GrpcServer] registry userController")

	//userRepoClient, err := userRepoFake.Repo()
	//if err != nil {
	//	log.Fatalf("failed to create fake repo: %s", err.Error())
	//}

	userRepoClient, err := userRepoMysql.Repo(config.CompletedMysqlConfig)
	if err != nil {
		log.Fatalf("failed to create Mysql repo: %s", err.Error())
	}

	userController := v1.NewController(userRepoClient)
	v1.RegisterUserServer(server, userController)
}
