package repository

import (
	"claudinary-cdn-api/models"

	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	Db *gorm.DB
}

func NewFileRepository(Db *gorm.DB) FileRepository {
	return &FileRepositoryImpl{Db: Db}
}

func (f *FileRepositoryImpl) Save(file *models.File) (*models.File, error) {
	if err := f.Db.Create(&file).Error; err != nil {
		return nil, err
	}
	return file, nil
}

func (f *FileRepositoryImpl) FindAllByPathID(pathID string) (*[]models.FileResUpload, error) {
	var files []models.File
	if err := f.Db.Where("path_id = ?", pathID).Where("show = ?", true).Order("created_at desc").Find(&files).Error; err != nil {
		return nil, err
	}

	var filesRes []models.FileResUpload
	for _, file := range files {
		filesRes = append(filesRes, models.FileResUpload{
			ID:        file.ID,
			PathID:    file.PathID,
			Url:       file.Url,
			CreatedAt: file.CreatedAt,
		})
	}
	return &filesRes, nil
}

func (f *FileRepositoryImpl) FindFileByID(fileID string) (*models.File, error) {
	var file models.File
	if err := f.Db.Where("id = ?", fileID).First(&file).Error; err != nil {
		return nil, err
	}

	return &file, nil
}

func (f *FileRepositoryImpl) HideFileByID(file *models.File) (*models.File, error) {
	if err := f.Db.Model(&file).Update("show", false).Error; err != nil {
		return nil, err
	}
	return file, nil
}
