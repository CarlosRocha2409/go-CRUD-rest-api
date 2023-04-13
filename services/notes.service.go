package services

import (
	"github.com/CarlosRocha2409/go-rest-api/models"
	"github.com/CarlosRocha2409/go-rest-api/repos"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteService struct {
	repo *repos.NoteRepo
}

func NewNoteService(client *mongo.Client) *NoteService {
	return &NoteService{
		repo: repos.NewNoteRepo(client),
	}
}

func (r *NoteService) GetAll(limit int64, page int64) (*[]models.Note, error) {
	notes, err := r.repo.GetAll(&limit, &page)
	return notes, err
}

func (r *NoteService) GetById(noteId *primitive.ObjectID) (*models.Note, error) {
	note, err := r.repo.GetById(noteId)
	return note, err
}

func (r *NoteService) Create(note *models.Note) (*interface{}, error) {
	created, err := r.repo.Create(note)
	return created, err
}

func (r *NoteService) Update(id *primitive.ObjectID, note *models.Note) (*mongo.UpdateResult, error) {
	updated, err := r.repo.Update(id, note)
	return updated, err
}

func (r *NoteService) Delete(id *primitive.ObjectID) (*mongo.DeleteResult, error) {
	deleted, err := r.repo.Delete(id)
	return deleted, err
}
