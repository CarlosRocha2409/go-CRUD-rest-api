package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Description string             `json:"description,omitempty" validate:"required"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type NoteServices interface {
	GetAll() ([]Note, error)
	GetById(id primitive.ObjectID) (Note, error)
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id primitive.ObjectID) error
}

type NoteRepo interface {
	GetAll() ([]Note, error)
	GetById(id primitive.ObjectID) (Note, error)
	Create(note *Note) (*Note, error)
	Update(note *Note) (*Note, error)
	Delete(id primitive.ObjectID) error
}
