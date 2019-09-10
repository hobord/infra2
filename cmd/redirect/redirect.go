package main

import (
	"net"
	"os"

	api "github.com/hobord/infra2/api/grpc/redirect"
	log "github.com/hobord/infra2/log"
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
		log.Logger.Fatalf("failed to listen: %v", err)
	}
	log.Logger.Infoln("Server listen: ", port)

	s := grpc.NewServer()
	reflection.Register(s)

	srv := apimpl.CreateGrpcServer()
	api.RegisterRedirectServiceServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("failed to serve: %v", err)
	}
}
