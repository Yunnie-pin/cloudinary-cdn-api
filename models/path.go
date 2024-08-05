package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Path struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	BucketName string    `gorm:"not null;references:Name" json:"bucket_name"`
}

type PathResponse struct {
	ID         uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	BucketName string    `gorm:"not null;references:Name" json:"bucket_name"`
}

func (u *Path) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}