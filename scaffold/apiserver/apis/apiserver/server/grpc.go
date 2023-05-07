// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/rebirthmonkey/go/pkg/log"
	"google.golang.org/grpc"

	policyv1 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/grpc/v1"
	policypb "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/controller/grpc/v1/pb"
	policyRepoMysql "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/policy/repo/mysql"
	userv1 "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/controller/grpc/v1"
	userpb "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/controller/grpc/v1/pb"
	userRepoMysql "github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/repo/mysql"
)

// InitGrpc initializes the Grpc server
func InitGrpc(server *grpc.Server) {
	log.Info("[GrpcServer] registry userController")

	userRepoClient, err := userRepoMysql.Repo(config.CompletedMysqlConfig)
	if err != nil {
		log.Fatalf("failed to create Mysql repo: %s", err.Error())
	}

	userController := userv1.NewController(userRepoClient)
	userpb.RegisterUserServer(server, userController)

	log.Info("[GrpcServer] registry policyController")

	policyRepoClient, err := policyRepoMysql.Repo(config.CompletedMysqlConfig)
	if err != nil {
		log.Fatalf("failed to create Mysql repo: %s", err.Error())
	}

	policyController := policyv1.NewController(policyRepoClient)
	policypb.RegisterPolicyServer(server, policyController)
}
