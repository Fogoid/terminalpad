package datamodels

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID      primitive.ObjectID `bson:"_id"`
	NoteID  string             `bson:"restaurant_id"`
	Title   string
	Content string
}

func NewNote(title, content string) *Note {
	return &Note{
		Title:   title,
		Content: content,
    }
}
