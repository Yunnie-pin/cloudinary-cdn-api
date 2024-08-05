package repository

import (
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryRepositoryImpl struct {
	Cd *cloudinary.Cloudinary
}

func NewCloudinaryRepository(Cd *cloudinary.Cloudinary) CloudinaryRepository {
	return &CloudinaryRepositoryImpl{Cd: Cd}
}

func (c *CloudinaryRepositoryImpl) SendImage(file *multipart.FileHeader, path string,ctx *gin.Context) (url string, err error) {

	uploadParams := uploader.UploadParams{
		Folder: path,
	}

	fileContent, err := file.Open()

	if err != nil {
		return "", err
	}

	defer fileContent.Close()

	uploadResult, err := c.Cd.Upload.Upload(ctx, fileContent, uploadParams)

	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
