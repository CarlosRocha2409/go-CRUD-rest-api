package utils

import (
	"github.com/CarlosRocha2409/go-rest-api/responses"
	"github.com/gin-gonic/gin"
)

func MakeResponse(c *gin.Context, status int, msg string, data gin.H) {
	c.JSON(status,
		responses.GenericResponse{
			Status:  status,
			Message: msg,
			Data:    data,
		},
	)

}

func MakePaginatedResponse(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(status,
		responses.PaginationResponse{
			Limit: 4,
			Page:  5,
			GenericResponse: responses.GenericResponse{
				Status:  status,
				Message: msg,
				Data:    data,
			},
		},
	)

}
