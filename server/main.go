package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/20Koen02/GoProto/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterMoneyTransactionServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	fmt.Println("Got amount ", in.Amount)
	fmt.Println("Got from ", in.From)
	fmt.Println("For ", in.To)

	return &pb.TransactionResponse{Confirmation: true}, nil
}
