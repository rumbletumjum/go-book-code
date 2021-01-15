package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "productinfo/client/ecommerce"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "FruitPhone17"
	description := "You don't need it"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added succesfullyv", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())

    product, err = c.GetProduct(ctx, &pb.ProductID{Value: "invalid"})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
}
