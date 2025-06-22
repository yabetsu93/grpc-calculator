package main

import (
	"context"
	"log"
	"net"

	"github.com/yabetsu93/grpc-calculator/proto"

	"google.golang.org/grpc"
)

type CalculatorServer struct {
	proto.UnimplementedCalculatorServiceServer
}

func(s *CalculatorServer) Add(ctx context.Context, req *protoCalcRequest) (*proto.CalcResponse, error) {
	result := req.Num1 + req.Num2

	return &proto.CalcResponse{Result: result}, nil
} 

func (s *CalculatorServer) Subtract(ctx context.Context, req *proto.CalcRequest) (*proto.CalcResponse, error) {
	result := req.Num1 - req.Num2
	return &proto.CalcResponse{Result: result}, nil
}

func (s *CalculatorServer) Divide(ctx context.Context, req *proto.CalcRequest)(*proto.CalcResponse, error) {
	if req.Num2 == 0 {
		return nil, grpc.Errorf(grpc.Code(3), "Division by zero")
	}
	result := req.Num1 / req.Num2
	return &proto.CalcResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCalcularServiceService(grpcServer, &CalculatorServer{})

	log.Println("gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}