package routes

import (
	"github.com/CarlosRocha2409/go-rest-api/controllers"
	"github.com/CarlosRocha2409/go-rest-api/models"
	"github.com/CarlosRocha2409/go-rest-api/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NotesRouter(router *gin.Engine, client *mongo.Client) {

	var controler *controllers.NoteController = controllers.NewNoteController(client)
	singleNote := router.Group("/notes/:noteId")

	router.GET("/notes", controler.GetAll())
	router.POST("/notes", utils.ValidateJson(models.Note{}), controler.Create())

	singleNote.Use(utils.CheckId("noteId"))
	{
		singleNote.GET("/", controler.GetById())
		singleNote.PUT("/", utils.ValidateJson(models.Note{}), controler.Update())

	}

}
