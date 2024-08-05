package models

import (
	"github.com/google/uuid"
)

type Path struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	BucketID string    `gorm:"not null;references:ID" json:"bucket_id"`
}

type PathResponse struct {
	ID       uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	BucketID string    `gorm:"not null;references:ID" json:"bucket_id"`
}
