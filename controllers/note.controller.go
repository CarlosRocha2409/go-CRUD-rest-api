package controllers

import (
	"fmt"
	"net/http"

	"github.com/CarlosRocha2409/go-rest-api/models"
	"github.com/CarlosRocha2409/go-rest-api/services"
	"github.com/CarlosRocha2409/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

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
		noteId := c.Param("noteId")

		id, err := primitive.ObjectIDFromHex(noteId)

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": "Please provide a valid Id",
			})
			return
		}

		result, err := ct.service.GetById(id)

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
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
		var note models.Note

		if err := c.BindJSON(&note); err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
			})
			return
		}

		if vErr := validate.Struct(&note); vErr != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": vErr.Error(),
			})
			return
		}

		result, err := ct.service.Create(&note)

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
		noteId := c.Param("noteId")
		var note models.Note

		id, err := primitive.ObjectIDFromHex(noteId)

		if err != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": "Please provide a valid Id",
			})
			return
		}

		if err := c.BindJSON(&note); err != nil {

			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
			})
			return
		}

		if vErr := validate.Struct(&note); vErr != nil {
			utils.MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": vErr.Error(),
			})
			return
		}

		result, err := ct.service.Update(&id, &note)

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
