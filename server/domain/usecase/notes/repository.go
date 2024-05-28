package notes

import (
	"context"
	"errors"
	"log/slog"

	"github.com/fogoid/terminalpad-server/domain/datamodels"
	generators "github.com/fogoid/terminalpad-server/pkg"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const NOTEID_LENGTH = 8

var (
    ErrObtainingCursor = errors.New("Could not obtain database cursor")
    ErrObtainingDocuments = errors.New("Could not obtain all documents from cursor")
    ErrDecodingResult = errors.New("Could not decode obtained element")
    ErrInsertingDocument = errors.New("Error inserting document in database")
    ErrCountingDocuments = errors.New("Error counting documents")
    ErrUpdatingNote = errors.New("Error updating note")
    ErrDeletingNote = errors.New("Error deleting note")
)

type NoteRepository struct {
    *mongo.Collection
}

func NewNoteRepository(collection *mongo.Collection) *NoteRepository {
    return &NoteRepository { collection }
}

func (r *NoteRepository) FindAllNotes() ([]*datamodels.Note, error) {
    cursor, err := r.Collection.Find(context.Background(), bson.M{})
    if err != nil {
        slog.Error(err.Error())
        return nil, ErrObtainingCursor
    }

    all := make([]*datamodels.Note, 0)
    err = cursor.All(context.Background(), all)
    if err != nil {
        slog.Error(err.Error())
        return nil, ErrObtainingCursor
    }

    return all, nil
}

func (r *NoteRepository) FindNote(id string) (*datamodels.Note, error) {
    document := &datamodels.Note { NoteID: id }
    singleResult := r.Collection.FindOne(context.Background(), document)
    if singleResult.Err() != nil {
        if singleResult.Err() == mongo.ErrNoDocuments {
            slog.Info("Could not find document with id %s", id)
            return nil, nil
        }

        slog.Error(singleResult.Err().Error())
        return nil, ErrObtainingDocuments
    }

    err := singleResult.Decode(document)
    if err != nil {
        slog.Error(err.Error())
        return nil, ErrDecodingResult
    }

    return document, nil
}

func (r *NoteRepository) CreateNote(note *datamodels.Note) (string, error) {
    note.NoteID = generators.GenerateRandomString(8)
    _, err := r.Collection.InsertOne(context.Background(), note)
    if err != nil {
        slog.Error("Could not insert record in database")
        return "", ErrInsertingDocument
    }

    return note.NoteID, nil
}

func (r *NoteRepository) UpdateNote(note *datamodels.Note) (string, error) {
    filter := &datamodels.Note { NoteID: note.NoteID }
    updateRes, err := r.Collection.UpdateOne(context.Background(), filter, note)
    if err != nil {
        slog.Error("Error updating note: %v", err)
        return "", ErrUpdatingNote
    }

    if updateRes.UpsertedCount == 0 {
        slog.Warn("No updates performed for note with id %s. No document found", note.NoteID)
        return "", ErrUpdatingNote
    }

    return note.NoteID, nil
}

func (r *NoteRepository) DeleteNote(id string) (int32, error) {
    filter := &datamodels.Note { NoteID: id }
    deleteRes, err := r.Collection.DeleteOne(context.Background(), filter)
    if err != nil {
        slog.Error("Error found while deleting note: %v", err)
        return 0, ErrDeletingNote
    }

    return int32(deleteRes.DeletedCount), nil
}
