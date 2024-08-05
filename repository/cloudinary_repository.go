package repository

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type CloudinaryRepository interface {
	SendImage(files *multipart.FileHeader, path string, ctx *gin.Context) (url string, err error)
}
