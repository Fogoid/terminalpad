package grpcservers

import (
	"context"
	"log/slog"

	"github.com/fogoid/terminalpad-server/domain/entities"
	"github.com/fogoid/terminalpad-server/domain/usecase/notes"
	"github.com/fogoid/terminalpad/proto"
	"google.golang.org/grpc"
)

type noteServer struct {
    notes.Service
    proto.UnimplementedNotepadServer
}

func RegisterNoteServer(s *grpc.Server, noteService notes.Service) {
    noteServer := &noteServer {Service: noteService}
    proto.RegisterNotepadServer(s, noteServer)
}

func (n *noteServer) GetAllNotes(e *proto.Empty, allNotesServer proto.Notepad_GetAllNotesServer) error {
    allNotes, err := n.Service.GetNotes()
    if err != nil {
        slog.Error("Error obtaining all notes: %v", err)
        return err
    }

    for _, note := range allNotes {
        nResponse := &proto.NoteResponse {
            Id: note.Id,
            Title: note.Title,
            Content: note.Content,
        }

        allNotesServer.Send(nResponse)
    }

    return nil

}

func (n *noteServer) GetNote(ctx context.Context, request *proto.NoteGetRequest) (*proto.NoteResponse, error) {
    note, err := n.Service.GetNote(request.Id)
    if err != nil {
        slog.Error("Error obtaining note: %v", err)
        return nil, err
    }

    response := &proto.NoteResponse {
        Id: note.Id,
        Title: note.Title,
        Content: note.Content,
    }
    return response, nil
}

func (n *noteServer) UpdateNote(ctx context.Context, request *proto.UpdateNoteRequest) (*proto.NoteResponse, error) {
    note := entities.NewNote(request.Id, request.Title, request.Content)
    id, err := n.Service.UpdateNote(note)
    if err != nil {
        slog.Error("Error updating note: %v", err)
    }

    response := &proto.NoteResponse {
        Id: id,
    }
    return response, nil
}


func (n *noteServer) CreateNote(ctx context.Context, request *proto.NewNoteRequest) (*proto.NoteResponse, error) {
    note := entities.NewNote("", request.Title, request.Content)
    id, err := n.Service.CreateNote(note)
    if err != nil {
        slog.Error("Error creating note: %v", err)
        return nil, err
    }

    response := &proto.NoteResponse {
        Id: id,
    }
    return response, nil

}


func (n *noteServer) DeleteNote(ctx context.Context, request *proto.DeleteNoteRequest) (*proto.Empty, error) {
    _, err := n.Service.DeleteNote(request.Id)
    if err != nil {
        return nil, err
    }

    return &proto.Empty{}, nil
}
