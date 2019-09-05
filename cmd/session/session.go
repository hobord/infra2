package main

import (
	"fmt"
	"net"
	"os"

	log "github.com/hobord/infra2/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	api "github.com/hobord/infra2/api/grpc/session"
	sessiongrpc "github.com/hobord/infra2/session/grpc"
	redistore "github.com/hobord/infra2/session/redistore"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":50051"
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Logger.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listen: ", port)

	s := grpc.NewServer()
	reflection.Register(s)

	// If You want to inject the redis connection pool by hand
	// redisConnectionPool = redistore.NewRedisPool()
	// store := redistore.CreateRedisStore(redisConnectionPool)
	// otherwise  the constructor create automatically using the OS ENV
	store := redistore.CreateRedisSessionStore(nil)

	rpcServer, err := sessiongrpc.CreateGrpcServer(store)
	if err != nil {
		log.Logger.Fatalf("failed create a grpc server: %v", err)
	}

	api.RegisterSessionServiceServer(s, rpcServer)

	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("failed to serve: %v", err)
	}
}
