package repository

import "claudinary-cdn-api/models"

type PathRepository interface {
	// FindAll() (*[]models.Path, error)
	Save(path *models.Path) (*models.Path, error)
	FindPathIDByBucketName(bucketName string, pathName string) (id string, err error)
}
