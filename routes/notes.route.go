package routes

import (
	"github.com/CarlosRocha2409/go-rest-api/controllers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NotesRouter(router *gin.Engine, client *mongo.Client) {

	var controler *controllers.NoteController = controllers.NewNoteController(client)

	router.GET("/notes", controler.GetAll())
	router.POST("/notes", controler.Create())
	router.GET("/notes/:noteId", controler.GetById())
	router.PUT("/notes/:noteId", controler.Update())

}
