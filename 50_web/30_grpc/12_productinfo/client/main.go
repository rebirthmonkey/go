package main

import (
	"context"
	pb "github.com/rebirthmonkey/pkg/grpc/productinfo"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "127.0.0.1:50051"
)

// 添加一个测试的商品
func AddProduct(ctx context.Context, client pb.ProductInfoClient) (id string) {
	aMac := &pb.Product{Name: "Mac Book Pro 2019", Description: "From Apple Inc."}
	productId, err := client.AddProduct(ctx, aMac)
	if err != nil {
		log.Println("add product fail.", err)
		return
	}
	log.Println("add product success, id = ", productId.Value)
	return productId.Value
}

// 获取一个商品
func GetProduct(ctx context.Context, client pb.ProductInfoClient, id string) {
	p, err := client.GetProduct(ctx, &pb.ProductId{Value: id})
	if err != nil {
		log.Println("get product err.", err)
		return
	}
	log.Printf("get prodcut success : %+v\n", p)
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Println("did not connect.", err)
		return
	}
	defer conn.Close()

	client := pb.NewProductInfoClient(conn)
	ctx := context.Background()

	id := AddProduct(ctx, client)
	GetProduct(ctx, client, id)
}