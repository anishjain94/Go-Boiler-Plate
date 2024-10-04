package middleware

import (
	"go-boiler-plate/common"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			HandleError(c, err)
		}
	}
}

func HandleError(c *gin.Context, err error) {
	statusCode, msg := SeparateCodeFromMsg(err.Error())

	c.JSON(statusCode, common.ToErrorDto(msg))
}

func SeparateCodeFromMsg(msg string) (int, string) {
	statusCode := http.StatusInternalServerError

	parts := strings.SplitN(msg, ":", 2)

	if len(parts) > 1 {
		if code, err := strconv.Atoi(parts[0]); err == nil {
			statusCode = code
			msg = strings.TrimSpace(parts[1])
		}
	}

	zap.L().Error(msg)

	return statusCode, msg
}
