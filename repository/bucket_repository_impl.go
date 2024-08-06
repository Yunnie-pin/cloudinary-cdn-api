package repository

import (
	"claudinary-cdn-api/models"

	"gorm.io/gorm"
)

type BucketRepositoryImpl struct {
	Db *gorm.DB
}

func NewBucketRepository(Db *gorm.DB) BucketRepository {
	return &BucketRepositoryImpl{Db: Db}
}

func (b *BucketRepositoryImpl) FindAll() (*[]models.Bucket, error) {
	var buckets []models.Bucket
	if err := b.Db.Preload("Paths").Find(&buckets).Error; err != nil {
		return nil, err
	}
	return &buckets, nil
}

func (b *BucketRepositoryImpl) FindBucketByName(bucketName string) (*models.Bucket, error) {
	var bucket models.Bucket
	if err := b.Db.Preload("Paths").Where("name = ?", bucketName).First(&bucket).Error; err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (b *BucketRepositoryImpl) Save(bucket *models.Bucket) (*models.Bucket, error) {
	if err := b.Db.Create(bucket).Error; err != nil {
		return nil, err
	}
	return bucket, nil
}
