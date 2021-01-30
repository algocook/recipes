package main

import (
	"net"
	"recipes/pkg/methods"

	"github.com/algocook/proto/recipes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	user     = "recipes"
	password = "recipes"
)

func main() {
	// Create new gRPC server instance
	s := grpc.NewServer()
	srv := &methods.RecipesMainServer{}

	// Register gRPC server
	recipes.RegisterRecipesServer(s, srv)

	// Listern on 5301
	listener, err := net.Listen("tcp", ":5301")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// Start gRPC server
	if err := s.Serve(listener); err != nil {
		grpclog.Fatalf("failed on starting server: %v", err)
	}
}
