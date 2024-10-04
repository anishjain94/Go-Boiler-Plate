package util

import (
	"go-boiler-plate/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHTTPPost[InputDtoType any, OutputDtoType any](serviceFunc func(c *gin.Context, dto *InputDtoType) *OutputDtoType) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto InputDtoType

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, common.ToErrorDto("Invalid Request Body"))
			return
		}

		response := serviceFunc(c, &dto)

		c.JSON(http.StatusOK, common.SuccessDto{
			Meta: common.AckDto{
				Success: true,
				Code:    "SUCCESS",
			},
			Data: response,
		})
	}
}

func HandleHTTPPut[InputDtoType any, OutputDtoType any](serviceFunc func(c *gin.Context, dto *InputDtoType) *OutputDtoType) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dto InputDtoType

		if err := c.ShouldBindJSON(&dto); err != nil {
			c.JSON(http.StatusBadRequest, common.ToErrorDto("Invalid Request Body"))
			return
		}

		response := serviceFunc(c, &dto)

		c.JSON(http.StatusOK, common.SuccessDto{
			Meta: common.AckDto{
				Success: true,
				Code:    "SUCCESS",
			},
			Data: response,
		})
	}
}

func HandleHTTPGet[OutputDtoType any](serviceFunc func(c *gin.Context) *OutputDtoType) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := serviceFunc(c)

		c.JSON(http.StatusOK, common.SuccessDto{
			Meta: common.AckDto{
				Success: true,
				Code:    "SUCCESS",
			},
			Data: response,
		})
	}
}
