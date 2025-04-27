// package main

// import (
// 	pb "distributed_calculator/proto"
// 	"fmt"
// 	"log"
// 	"net"

// 	"github.com/Knetic/govaluate"
// 	"google.golang.org/grpc"
// )

// type server struct {
// 	pb.UnimplementedCalculatorServer
// }

// func (s *server) Calculate(req *pb.CalculationRequest, stream pb.Calculator_CalculateServer) error {
// 	expression, err := govaluate.NewEvaluableExpression(req.Expression)
// 	if err != nil {
// 		return fmt.Errorf("invalid expression: %w", err)
// 	}
// 	result, err := expression.Evaluate(nil)
// 	if err != nil {
// 		return fmt.Errorf("evaluation error: %w", err)
// 	}
// 	floatResult, ok := result.(float64)
// 	if !ok {
// 		return fmt.Errorf("result is not a float")
// 	}
// 	return stream.SendAndClose(&pb.CalculationResponse{Result: floatResult})
// }

// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterCalculatorServer(s, &server{})
// 	log.Println("gRPC Calculator Service started on :50051")
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
package main

import (
 pb "distributed_calculator/proto/proto"
 "fmt"
 "log"
 "net"
 "context"
 "github.com/Knetic/govaluate"
 "google.golang.org/grpc"
)

type server struct {
 pb.UnimplementedCalculatorServer
}

func (s *server) Calculate(ctx context.Context, req *pb.CalculationRequest) (*pb.CalculationResponse, error) {
 expression, err := govaluate.NewEvaluableExpression(req.Expression)
 if err != nil {
  return nil, fmt.Errorf("invalid expression: %w", err)
 }
 result, err := expression.Evaluate(nil)
 if err != nil {
  return nil, fmt.Errorf("evaluation error: %w", err)
 }
 floatResult, ok := result.(float64)
 if !ok {
  return nil, fmt.Errorf("result is not a float")
 }
 return &pb.CalculationResponse{Result: floatResult}, nil
}

func main() {
 lis, err := net.Listen("tcp", ":50051")
 if err != nil {
  log.Fatalf("failed to listen: %v", err)
 }
 s := grpc.NewServer()
 pb.RegisterCalculatorServer(s, &server{})
 log.Println("gRPC Calculator Service started on :50051")
 if err := s.Serve(lis); err != nil {
  log.Fatalf("failed to serve: %v", err)
 }
}