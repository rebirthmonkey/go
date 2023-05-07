package main

import (
	"context"
	"log"

	"github.com/rebirthmonkey/pkg/grpc/productinfo/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewProductInfoClient(conn)

	macProduct := &pb.Product{Name: "Mac Book Pro 2022", Description: "From Apple Inc."}
	productId, err := client.AddProduct(context.TODO(), macProduct)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success with productId", productId)

	product, err := client.GetProduct(context.TODO(), productId)
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Println("get product success with product:", product)
}
