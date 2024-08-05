package helpers

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}

func TrimmedString(str string) string {
	return strings.TrimPrefix(str, "/")
}
