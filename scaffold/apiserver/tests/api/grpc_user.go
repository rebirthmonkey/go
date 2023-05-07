// Copyright 2022 Wukong SUN <rebirthmonkey@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/rebirthmonkey/go/scaffold/apiserver/apis/apiserver/user/controller/grpc/v1/pb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		fmt.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	c := &pb.ListUsersRequest{}
	sl, _ := client.ListUsers(context.TODO(), c)
	fmt.Print("User List is:", sl)
}
