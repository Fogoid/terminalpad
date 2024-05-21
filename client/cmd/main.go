package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/fogoid/terminalpad/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
    addr = flag.String("addr", "localhost:8080", "The server address to connect to")
)

func main() {
    flag.Parse()
    conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Could not connect to gRPC server: %v", err)
    }
    defer conn.Close()
    c := proto.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    request := &proto.HelloRequest { Name: "John" }
    r, err := c.SayHello(ctx, request)
    if err != nil {
        log.Printf("Could not greet properly: %v", err)
    }

    log.Printf("Received response: %s", r.Message)
}
