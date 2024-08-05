package controllers

import (
	"claudinary-cdn-api/data"
	"claudinary-cdn-api/helpers"
	"claudinary-cdn-api/models"
	"claudinary-cdn-api/repository"
	"time"

	// "fmt"

	"net/http"

	// "github.com/cloudinary/cloudinary-go/v2/api"
	// "github.com/cloudinary/cloudinary-go/v2/api/admin"

	"github.com/gin-gonic/gin"
)

type UploaderController struct {
	bucketRepository     repository.BucketRepository
	pathRepository       repository.PathRepository
	fileRepository       repository.FileRepository
	cloudinaryRepository repository.CloudinaryRepository
}

func NewUploaderController(
	bucketRepository repository.BucketRepository,
	pathRepository repository.PathRepository,
	fileRepository repository.FileRepository,
	cloudinaryRepository repository.CloudinaryRepository,
) *UploaderController {
	return &UploaderController{
		bucketRepository:     bucketRepository,
		pathRepository:       pathRepository,
		fileRepository:       fileRepository,
		cloudinaryRepository: cloudinaryRepository,
	}
}

func (controller *UploaderController) UploadFile(ctx *gin.Context) {
	bucket := ctx.PostForm("bucket")
	path := ctx.PostForm("path")
	// Multipart form
	form, _ := ctx.MultipartForm()
	files := form.File["files[]"]

	_, err := controller.bucketRepository.FindBucketByName(bucket)
	if err != nil {
		helpers.ErrorNotFound(ctx, err, "Bucket not found")
		return
	}

	pathID, err := controller.pathRepository.FindPathIDByBucketName(bucket, path)
	if err != nil {
		helpers.ErrorNotFound(ctx, err, "Path not found")
		return
	}

	var urls []string

	for _, file := range files {
		url, err := controller.cloudinaryRepository.SendImage(file, path, ctx)
		if err != nil {
			helpers.ErrorResponse(ctx, err, "Error when upload file")
			return
		}

		urls = append(urls, url)
	}

	// Save all file to database
	allFiles := []models.FileResUpload{}
	for _, url := range urls {
		file := models.File{
			Url:       url,
			PathID:    pathID,
			Show:      true,
			CreatedAt: time.Now(),
		}

		newFile, err := controller.fileRepository.Save(&file)
		if err != nil {
			helpers.ErrorResponse(ctx, err, "Error when save file")
			return
		}

		allFiles = append(allFiles, models.FileResUpload{
			ID:        newFile.ID,
			PathID:    newFile.PathID,
			Url:       newFile.Url,
			CreatedAt: newFile.CreatedAt,
		})
	}


	// allFiles, err := controller.fileRepository.FindAllByPathID(pathID)

	// if err != nil {
	// 	helpers.ErrorResponse(ctx, err, "Error when get all files")
	// 	return
	// }

	result := models.ResponseUpload{
		BucketName: bucket,
		PathName:   path,
		File:       allFiles,
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "claudinary-cdn-api",
		Controller: "UploaderController",
		Action:     "UploadFile",
		Result:     result,
	}

	ctx.JSON(200, response)
}
