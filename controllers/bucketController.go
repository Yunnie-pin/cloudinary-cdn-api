package controllers

import (
	"claudinary-cdn-api/data"
	"claudinary-cdn-api/helpers"
	"claudinary-cdn-api/models"
	"claudinary-cdn-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BucketController struct {
	bucketRepository repository.BucketRepository
}

func NewBucketController(repository repository.BucketRepository) *BucketController {
	return &BucketController{bucketRepository: repository}
}

func (controller *BucketController) FindAll(ctx *gin.Context) {
	bucket, err := controller.bucketRepository.FindAll()

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Tidak bisa menemukan data bucket")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "skincare-server",
		Controller: "bucket",
		Action:     "GetListBucket",
		Result:     bucket,
	}

	ctx.JSON(http.StatusOK, response)
}

func (controller *BucketController) CreateBucket(ctx *gin.Context) {
	name := ctx.PostForm("name")

	_, err := controller.bucketRepository.FindBucketByName(name)
	if err == nil {
		helpers.ErrorResponse(ctx, err, "Bucket sudah ada")
		return
	}

	newBucket := models.Bucket{
		Name: name,
	}

	nbucket, err := controller.bucketRepository.Save(&newBucket)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Tidak bisa membuat bucket")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "skincare-server",
		Controller: "bucket",
		Action:     "CreateBucket",
		Result:     nbucket,
	}

	ctx.JSON(http.StatusOK, response)
}
