package models

type Bucket struct {
	Name  string `gorm:"primary_key;type:varchar(255);not null" json:"bucket"`
	Paths []Path `gorm:"foreignkey:BucketName;references:Name;" json:"paths"`
}
