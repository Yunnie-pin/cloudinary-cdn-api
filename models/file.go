package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type File struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
	PathID    string    `gorm:"not null;references:ID" json:"path_id"`
	Show      bool      `gorm:"default:true" json:"show"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type FileResUpload struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
	PathID    string    `gorm:"not null;references:ID" json:"path_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (u *File) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}