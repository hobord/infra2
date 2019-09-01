package main

import (
	"fmt"
	"log"
	"net"
	"os"

	api "github.com/hobord/infra2/api/grpc/redirect"
	apimpl "github.com/hobord/infra2/redirect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":50052"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listen: ", port)

	s := grpc.NewServer()
	reflection.Register(s)

	srv := apimpl.CreateGrpcServer()
	api.RegisterRedirectServiceServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
