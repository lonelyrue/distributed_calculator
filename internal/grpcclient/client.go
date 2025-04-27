package grpcclient

import (
 "context"
 pb "distributed_calculator/proto/proto"
 "log"

 "google.golang.org/grpc"
)

type GRPCClient struct {
 conn   *grpc.ClientConn
 client pb.CalculatorClient
}

// Создание нового GRPC клиента
func NewGRPCClient(address string) *GRPCClient {
 conn, err := grpc.Dial(address, grpc.WithInsecure())
 if err != nil {
  log.Fatalf("failed to connect to gRPC server: %v", err)
 }
 client := pb.NewCalculatorClient(conn)
 return &GRPCClient{conn: conn, client: client}
}

// Метод для выполнения расчётов через gRPC
func (c *GRPCClient) Calculate(expression string) (float64, error) {
 resp, err := c.client.Calculate(context.Background(), &pb.CalculationRequest{
  Expression: expression,
 })
 if err != nil {
  return 0, err
 }
 return resp.Result, nil
}