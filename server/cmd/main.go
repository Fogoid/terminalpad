package main

import (
	"log"
	"net"

	"github.com/fogoid/terminalpad-server/cmd/grpcservers"
	"github.com/fogoid/terminalpad-server/domain/usecase/notes"
	"github.com/fogoid/terminalpad-server/infrastructure/mongodb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)


func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading environment variables")
    }

    mongoNotesCollection := mongodb.GetDbCollection()
    defer mongodb.Disconnect()

    noteRepository := notes.NewNoteRepository(mongoNotesCollection)
    noteService := notes.NewNoteService(noteRepository)
    s := grpc.NewServer()
    grpcservers.RegisterNoteServer(s, noteService)

    lis, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to create listener: %v", err)
    }

    err = s.Serve(lis)
    if err != nil {
        log.Fatalf("Failed to serve listener: %v", err)
    }
}
