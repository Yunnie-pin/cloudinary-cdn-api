package repository

import "claudinary-cdn-api/models"

type FileRepository interface {
	Save(file *models.File) (*models.File, error)
	FindAllByPathID(pathID string) (*[]models.FileResUpload, error)
}