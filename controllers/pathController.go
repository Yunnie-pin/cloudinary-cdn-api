package controllers

import (
	"claudinary-cdn-api/data"
	"claudinary-cdn-api/helpers"
	"claudinary-cdn-api/models"
	"claudinary-cdn-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PathController struct {
	bucketRepository repository.BucketRepository
	pathRepository   repository.PathRepository
}

func NewPathController(
	bucketRepository repository.BucketRepository,
	pathRepository repository.PathRepository,
) *PathController {
	return &PathController{
		bucketRepository: bucketRepository,
		pathRepository:   pathRepository,
	}
}

func (controller *PathController) CreatePath(ctx *gin.Context) {

	bucket := ctx.PostForm("bucket")
	path := ctx.PostForm("path")

	_, err := controller.bucketRepository.FindBucketByName(bucket)
	if err != nil {
		helpers.ErrorNotFound(ctx, err, "Bucket not found")
		return
	}

	_, err = controller.pathRepository.FindPathIDByBucketName(bucket, path)
	if err == nil {
		helpers.ErrorResponse(ctx, err, "Path already exists")
		return
	}

	newPath := models.Path{
		Name:       path,
		BucketName: bucket,
	}

	npath, err := controller.pathRepository.Save(&newPath)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Cannot create path")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "skincare-server",
		Controller: "path",
		Action:     "CreatePath",
		Result:     npath,
	}

	ctx.JSON(http.StatusOK, response)
}
