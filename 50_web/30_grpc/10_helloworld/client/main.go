/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"github.com/rebirthmonkey/pkg/grpc/helloworld/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect: ", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	resp, err := client.SayHello(context.TODO(), &pb.HelloRequest{Name: "xxx"})
	if err != nil {
		log.Fatalln("could not greet: ", err)
	}
	log.Println("Greeting: ", resp.GetMessage())

	resp, err = client.SayHelloAgain(context.TODO(), &pb.HelloRequest{Name: "yyy"})
	if err != nil {
		log.Fatalln("could not greet: ", err)
	}
	log.Println("Greeting: ", resp.GetMessage())
}
