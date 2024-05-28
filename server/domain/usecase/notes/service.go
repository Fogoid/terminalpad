package notes

import (
	"errors"

	"github.com/fogoid/terminalpad-server/domain/entities"
)

var (
	NoteNotFoundError      = errors.New("Note not found")
	NoteAlreadyExistsError = errors.New("Note already exists")
)

type NoteService struct {
	notes map[string]*entities.Note
}

func NewNoteService() *NoteService {
	return &NoteService{
		notes: make(map[string]*entities.Note, 0),
	}
}

func (n *NoteService) GetNotes() ([]*entities.Note, error) {
	all := make([]*entities.Note, 0)
	for _, note := range n.notes {
		all = append(all, note)
	}

	return all, nil
}

func (n *NoteService) GetNote(id string) (*entities.Note, error) {
	if n.notes[id] == nil {
		return nil, NoteNotFoundError
	}

	return n.notes[id], nil
}

func (n *NoteService) CreateNote(note *entities.Note) (*entities.Note, error) {
	if n.notes[note.Title] != nil {
		return nil, NoteAlreadyExistsError
	}

	n.notes[note.Title] = note
	return note, nil
}

func (n *NoteService) UpdateNote(note *entities.Note) (*entities.Note, error) {
	if n.notes[note.Title] == nil {
		return nil, NoteNotFoundError
	}

	n.notes[note.Title] = note
	return note, nil
}

func (n *NoteService) DeleteNote(id string) error {
	n.notes[id] = nil
	return nil
}
