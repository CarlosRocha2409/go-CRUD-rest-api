package controllers

import (
	"fmt"
	"net/http"

	"github.com/CarlosRocha2409/go-rest-api/models"
	"github.com/CarlosRocha2409/go-rest-api/services"
	"github.com/CarlosRocha2409/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoteController struct {
	service *services.NoteService
}

func NewNoteController(client *mongo.Client) *NoteController {
	return &NoteController{
		service: services.NewNoteService(client),
	}
}

func (ct *NoteController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		notes, err := ct.service.GetAll()

		if err != nil {
			fmt.Println(err.Error())
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": "Error getting notes",
			})
			return
		}

		utils.MakePaginatedResponse(c, http.StatusOK, "Ok", gin.H{
			"notes": notes,
		})
	}
}

func (ct *NoteController) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("id")

		result, err := ct.service.GetById(id.(*primitive.ObjectID))

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": fmt.Sprintf("User with id: %v", id.(*primitive.ObjectID).String()),
			})
			return
		}

		utils.MakeResponse(c, http.StatusOK, "Ok", gin.H{
			"note": result,
		})

	}
}

func (ct *NoteController) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		note, _ := c.Get("body")
		result, err := ct.service.Create(note.(*models.Note))

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
			})
			return
		}

		utils.MakeResponse(c, http.StatusOK, "Ok", gin.H{
			"id": result,
		})

	}
}

func (ct *NoteController) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Get("id")
		note, _ := c.Get("body")

		result, err := ct.service.Update(id.(*primitive.ObjectID), note.(*models.Note))

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
			})
			return
		}

		utils.MakeResponse(c, http.StatusOK, "Ok", gin.H{
			"id": result.MatchedCount,
		})
	}
}
