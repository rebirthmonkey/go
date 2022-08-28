// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github.com/rebirthmonkey/go/60_paas/20_db/10_mysql/80_server/apiserver/user/controller/grpc/v1"
)

func main() {
	const address = "127.0.0.1:8081"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	ctx := context.Background()
	c := &pb.ListUsersRequest{}
	sl, _ := client.ListUsers(ctx, c)
	fmt.Print("User List is:", sl)
}
