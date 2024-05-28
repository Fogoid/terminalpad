package main

import (
	"log"
	"net"

	"github.com/fogoid/terminalpad-server/cmd/grpcservers"
	"github.com/fogoid/terminalpad-server/domain/usecase/notes"
	"google.golang.org/grpc"
)


func main() {
    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to create listener: %v", err)
    }

    noteService := notes.NewNoteService()
    s := grpc.NewServer()
    grpcservers.RegisterNoteServer(s, noteService)

    err = s.Serve(lis)
    if (err != nil) {
        log.Fatalf("Failed to serve listener: %v", err)
    }
}
