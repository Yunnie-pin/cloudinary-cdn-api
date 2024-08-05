package repository

import (
	"claudinary-cdn-api/models"
)

type BucketRepository interface {
	FindAll() (*[]models.Bucket, error)
	FindBucketByName(bucketName string) (*models.Bucket, error)
}