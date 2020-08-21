package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/g-vit/simple-grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	var name string
	var srvAddr string

	if sa, ok := os.LookupEnv("SERVER_ADDR"); ok && (len(sa) > 0) {
		srvAddr = sa
	} else {
		fmt.Printf("Missing or malformed SERVER_ADDR!\n")
		os.Exit(1)
	}

	if n, ok := os.LookupEnv("NAME"); ok && (len(n) > 0) {
		name = n
	} else {
		fmt.Printf("Missing or malformed env NAME. Default NAME = World\n")
		name = "World"
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(srvAddr, opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewHelloClient(conn)
	for {
		request := &pb.Request{
			Message: name,
		}

		response, err := client.Do(context.Background(), request)

		if err != nil {
			grpclog.Fatalf("fail to dial: %v", err)
		}

		fmt.Println(response.Message)
		time.Sleep(5 * time.Second)
	}
}
