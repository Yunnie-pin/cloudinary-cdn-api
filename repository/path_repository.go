package repository

type PathRepository interface {
	// FindAll() (*[]models.Path, error)
	FindPathIDByBucketName(bucketName string, pathName string) (id string, err error)
}
