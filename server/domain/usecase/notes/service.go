package notes

import (
	"errors"
	"log/slog"

	"github.com/fogoid/terminalpad-server/domain/entities"
	"github.com/fogoid/terminalpad-server/domain/utils"
)

var (
	NoteNotFoundError      = errors.New("Note not found")
	NoteAlreadyExistsError = errors.New("Note already exists")
    ErrFetchingRecords = errors.New("Error fetching database records")
)

type NoteService struct {
    repository Repository
	notes map[string]*entities.Note
}

func NewNoteService(repository Repository) *NoteService {
	return &NoteService{
        repository,
	    make(map[string]*entities.Note, 0),
	}
}

func (s *NoteService) GetNotes() ([]*entities.Note, error) {
    all, err := s.repository.FindAllNotes()
    if err != nil {
        slog.Error("Could not fetch records from database: %v", err)
        return nil, ErrFetchingRecords
    }

    allRec := make([]*entities.Note, 0)
    for _, v := range all {
        e := utils.DatamodelToEntity(*v)
        allRec = append(allRec, e)
    }

	return allRec, nil
}

func (s *NoteService) GetNote(id string) (*entities.Note, error) {
    res, err := s.repository.FindNote(id)
    if err != nil {
        slog.Error("Error obtaining note from database: %v", err)
        return nil, ErrFetchingRecords
    }

    entity := utils.DatamodelToEntity(*res)
	return entity, nil
}

func (s *NoteService) CreateNote(note *entities.Note) (string, error) {
    dm := utils.EntityToDatamodel(*note)
    id, err := s.repository.CreateNote(dm)
    if err != nil {
        slog.Error("Error creating note in database: %v", err)
        return "", ErrInsertingDocument
    }

	return id, nil
}

func (s *NoteService) UpdateNote(note *entities.Note) (string, error) {
    dm := utils.EntityToDatamodel(*note)
    id, err := s.repository.UpdateNote(dm)
    if err != nil {
        slog.Error("Error updating note: %v", err)
        return "", ErrUpdatingNote
    }

    return id, nil
}

func (s *NoteService) DeleteNote(id string) (int32, error) {
    count, err := s.repository.DeleteNote(id)
    if err != nil {
        slog.Error("Error deleting note: %v", err)
        return int32(0), ErrDeletingNote
    }

	return count, nil
}
