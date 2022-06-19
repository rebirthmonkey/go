package main

import (
	"context"
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	macProduct := &pb.Product{Name: "Mac Book Pro 2022", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, macProduct)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success with productId", productId)

	product, err := client.GetProduct(ctx, productId)
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Println("get product success with product:", product)
}
