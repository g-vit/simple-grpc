package main

import (
	ctx "context"
	"fmt"
	"net"

	pb "github.com/g-vit/simple-grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":9292")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServer(grpcServer, &server{})
	grpcServer.Serve(lis)
}

func (s *server) Do(c ctx.Context, request *pb.Request) (response *pb.Response, err error) {
	msg := request.Message
	response = &pb.Response{
		Message: fmt.Sprintf("Hello %s", msg),
	}
	return response, nil
}
