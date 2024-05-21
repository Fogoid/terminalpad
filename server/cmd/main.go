package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fogoid/terminalpad/proto"
	"google.golang.org/grpc"
)

type server struct {
    proto.UnimplementedGreeterServer
}

func (s *server) SayHello(context context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
    log.Printf("Received request: %v", request.Name)
    return &proto.HelloReply { Message: fmt.Sprintf("It's a pleasure to meet you, %s", request.Name) }, nil
}
 
func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to create listener: %v", err)
    }

    s := grpc.NewServer()
    proto.RegisterGreeterServer(s, &server{})

    err = s.Serve(lis)
    if (err != nil) {
        log.Fatalf("Failed to serve listener: %v", err)
    }
}
