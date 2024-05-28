package notes

import (
	"github.com/fogoid/terminalpad-server/domain/datamodels"
	"github.com/fogoid/terminalpad-server/domain/entities"
)

type Repository interface {
    FindAllNotes() ([]*datamodels.Note, error)
    FindNote(id string) (*datamodels.Note, error)
    CreateNote(*datamodels.Note) (string, error)
    UpdateNote(*datamodels.Note) (string, error)
    DeleteNote(id string) (int32, error)
}

type Service interface {
    GetNotes() ([]*entities.Note, error)
    GetNote(id string) (*entities.Note, error)
    CreateNote(*entities.Note) (string, error)
    UpdateNote(*entities.Note) (string, error)
    DeleteNote(id string) (int32, error)
}
