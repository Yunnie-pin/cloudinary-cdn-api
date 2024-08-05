package models

import (
	"github.com/google/uuid"
)

type Bucket struct {
	ID    uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Name  string    `gorm:"type:varchar(255);not null" json:"name"`
	Paths []Path    `gorm:"foreignkey:BucketID;references:ID;" json:"paths"`
}
