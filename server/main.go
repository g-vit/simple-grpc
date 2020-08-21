package main

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	lis, err := net.Listen("tcp", ":9292")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	grpcServer.Serve(lis)
}
