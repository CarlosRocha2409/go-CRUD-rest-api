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

func MakePaginatedResponse(c *gin.Context, status int, msg string, data interface{}, limit int64, page int64) {
	c.JSON(status,
		responses.PaginationResponse{
			Limit: limit,
			Page:  page,
			GenericResponse: responses.GenericResponse{
				Status:  status,
				Message: msg,
				Data:    data,
			},
		},
	)

}
