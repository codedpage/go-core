package main

import (
	"log"
	"net"
	"time"

	pb "grpc-go/greet/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var greetWithDeadlineTime time.Duration = 1 * time.Second

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("Listening at %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
