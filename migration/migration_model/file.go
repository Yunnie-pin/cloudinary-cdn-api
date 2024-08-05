package migration_model

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;" json:"id"`
	Url       string    `gorm:"type:varchar(255);not null" json:"url"`
	PathID    string    `gorm:"not null;references:ID" json:"path_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
