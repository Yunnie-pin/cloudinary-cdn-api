package repository

import (
	"claudinary-cdn-api/models"

	"gorm.io/gorm"
)

type PathRepositoryImpl struct {
	Db *gorm.DB
}

func NewPathRepository(Db *gorm.DB) PathRepository {
	return &PathRepositoryImpl{Db: Db}
}

func (p *PathRepositoryImpl) FindPathIDByBucketName(bucketName string, pathName string) (id string, err error) {
	var path models.Path
	if err := p.Db.Where("bucket_name = ? AND name = ?", bucketName, pathName).First(&path).Error; err != nil {
		return "", err
	}
	return string(path.ID.String()), nil
}

func (p *PathRepositoryImpl) Save(path *models.Path) (*models.Path, error) {
	if err := p.Db.Create(path).Error; err != nil {
		return nil, err
	}
	return path, nil
}
