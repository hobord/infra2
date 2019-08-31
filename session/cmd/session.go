package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	sessiongrpc "github.com/hobord/infra2/session/grpc"
	redistore "github.com/hobord/infra2/session/redistore"
	api "github.com/hobord/infra2/api/grpc/session"
	 
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":50051"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listen: ", port)

	s := grpc.NewServer()
	reflection.Register(s)

	// redisConnectionPool = redistore.NewRedisPool()
	store := redistore.CreateRedisStore(nil)
	
	rpcServer, err := sessiongrpc.CreateGrpcServer(store)
	if err != nil {
		log.Fatalf("failed create a grpc server: %v", err)
	}

	api.RegisterSessionServiceServer(s, rpcServer)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
