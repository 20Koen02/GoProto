package main

import (
	"log"

	pb "github.com/20Koen02/GoProto/protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const address = "localhost:8000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while making connection, %v", err)
	}

	c := pb.NewMoneyTransactionClient(conn)

	c.MakeTransaction(
		context.Background(),
		&pb.TransactionRequest{
			From:   "John",
			To:     "Alice",
			Amount: float32(120.15),
		},
	)
}
