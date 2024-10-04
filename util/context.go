package util

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetContextQueries(c *gin.Context) url.Values {
	return c.Request.URL.Query()
}

func GetContextParams(c *gin.Context) gin.Params {
	return c.Params
}

func GetContextQuery(c *gin.Context, key string) (string, bool) {
	return c.GetQuery(key)
}

func GetContextParam(c *gin.Context, key string) string {
	return c.Param(key)
}
