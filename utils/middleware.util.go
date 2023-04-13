package utils

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CheckId(paramName string) gin.HandlerFunc {

	return func(c *gin.Context) {
		noteId := c.Param(paramName)
		id, err := primitive.ObjectIDFromHex(noteId)

		if err != nil {
			MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": "Please provide a valid Id",
			})
			return
		}
		c.Set("id", &id)
		c.Next()
	}
}

func ValidateJson(modelPtr interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		model := reflect.New(reflect.TypeOf(modelPtr)).Interface()

		if err := c.BindJSON(&model); err != nil {
			MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": err.Error(),
			})
			return
		}

		if vErr := validate.Struct(model); vErr != nil {
			MakeResponse(c, http.StatusBadRequest, "error", gin.H{
				"error": vErr.Error(),
			})
			return
		}

		c.Set("body", model)
		c.Next()
	}
}
