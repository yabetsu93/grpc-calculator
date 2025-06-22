package main

import (
	"context"
	"log"
	"time"

	"github.com/yabetsu93/grpc-calculator/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewCalculatorServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &proto.CalcRequest{Num1: 12, Num2:4}
	res, err := client.Divide(ctx, req)
	if err != nil {
		log.Fatalf("Error during RPC: %v", err)
	}
	log.Printf("Result: %v", res.Result)
}