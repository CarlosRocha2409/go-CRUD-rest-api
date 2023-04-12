package main

import (
	"github.com/CarlosRocha2409/go-rest-api/configs"
	"github.com/CarlosRocha2409/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mongoClient := configs.ConnectDb()
	routes.NotesRouter(router, mongoClient)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": "Hello world",
		})
	})

	router.Run("localhost:5000")
}
