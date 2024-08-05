package controllers

import (
	"claudinary-cdn-api/data"
	"claudinary-cdn-api/helpers"
	"claudinary-cdn-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	bucketRepository repository.BucketRepository
	pathRepository   repository.PathRepository
	fileRepository   repository.FileRepository
}

func NewFileController(bucketRepository repository.BucketRepository,
	pathRepository repository.PathRepository,
	fileRepository repository.FileRepository) *FileController {
	return &FileController{bucketRepository: bucketRepository,
		pathRepository: pathRepository,
		fileRepository: fileRepository}
}

func (controller *FileController) FindAllFiles(ctx *gin.Context) {
	bucket := ctx.Param("bucket")
	path := helpers.TrimmedString(ctx.Param("path"))

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

	files, err := controller.fileRepository.FindAllByPathID(pathID)

	if err != nil {
		helpers.ErrorResponse(ctx, err, "Tidak bisa menemukan data file")
		return
	}

	response := data.ResponseModel{
		Response:   http.StatusOK,
		Error:      "",
		AppID:      "skincare-server",
		Controller: "file",
		Action:     "GetListFile",
		Result:     files,
	}

	ctx.JSON(http.StatusOK, response)
}
