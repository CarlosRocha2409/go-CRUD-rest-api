package repos

import (
	"context"
	"time"

	"github.com/CarlosRocha2409/go-rest-api/configs"
	"github.com/CarlosRocha2409/go-rest-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteRepo struct {
	notes *mongo.Collection
}

func NewNoteRepo(client *mongo.Client) *NoteRepo {
	return &NoteRepo{
		notes: configs.GetCollection(client, "notes"),
	}
}

func (r *NoteRepo) GetAll() (*[]models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	var notes []models.Note
	defer cancel()

	results, err := r.notes.Find(ctx, bson.M{})

	if err != nil {
		return &notes, err
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var note models.Note
		err = results.Decode(&note)
		if err == nil {
			notes = append(notes, note)
		}

	}

	return &notes, err
}

func (r *NoteRepo) GetById(noteId *primitive.ObjectID) (*models.Note, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	var note models.Note
	defer cancel()
	err := r.notes.FindOne(ctx, bson.M{"id": noteId}).Decode(&note)

	return &note, err

}

func (r *NoteRepo) Create(note *models.Note) (*interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	newNote := models.Note{
		ID:          primitive.NewObjectID(),
		Title:       note.Title,
		Description: note.Description,
	}

	result, err := r.notes.InsertOne(ctx, newNote)

	return &result.InsertedID, err

}

func (r *NoteRepo) Update(id *primitive.ObjectID, note *models.Note) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	newNote := bson.M{
		"title":       note.Title,
		"description": note.Description,
	}

	result, err := r.notes.UpdateOne(ctx, bson.M{"id": (*id)}, bson.M{"$set": newNote})

	return result, err
}

func (r *NoteRepo) Delete(id *primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	result, err := r.notes.DeleteOne(ctx, bson.M{"id": (*id)})

	return result, err

}
