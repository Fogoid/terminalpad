package utils

import (
	"github.com/fogoid/terminalpad-server/domain/datamodels"
	"github.com/fogoid/terminalpad-server/domain/entities"
)

func DatamodelToEntity(note datamodels.Note) *entities.Note {
    return entities.NewNote(note.NoteID, note.Title, note.Content)
}

func EntityToDatamodel(note entities.Note) *datamodels.Note {
    n := datamodels.NewNote(note.Title, note.Content)
    n.NoteID = note.Id

    return n
}
