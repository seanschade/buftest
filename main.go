package main

import (
	"fmt"
	"net"

	pingv1 "go.buf.build/grpc/go/seanschade/buftest"
	"google.golang.org/grpc"
)

func main() {
	listenOn := "127.0.0.1:8000"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		fmt.Errorf("failed to listen on %s: %w", listenOn, err)
	}

	server := grpc.NewServer()
	buftest.RegisterPingServiceServer(server, &pingService{})

	if err := server.Serve(listener); err != nil {
		fmt.Errorf("failed to serve gRPC server: %w", err)
	}
}

type pingService struct {
	buftest.UnimplementedPingServiceServer
}
