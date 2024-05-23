package notes

import (
    "github.com/fogoid/terminalpad-server/domain/entities"
)

type Service interface {
    GetNotes() ([]*entities.Note, error)
    GetNote(id string) ([]*entities.Note, error)
    CreateNote(*entities.Note) (*entities.Note, error)
    UpdateNote(*entities.Note) (*entities.Note, error)
    DeleteNote(id string) (*entities.Note, error)
}
