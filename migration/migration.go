package migration

import (
	models "claudinary-cdn-api/migration/migration_model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func MigrationDB(db *gorm.DB) (err error) {
	db.Migrator().DropTable(&models.Bucket{})
	db.Migrator().DropTable(&models.Path{})
	db.Migrator().DropTable(&models.File{})

	//create table
	db.Table("buckets").AutoMigrate(&models.Bucket{})
	db.Table("paths").AutoMigrate(&models.Path{})
	db.Table("files").AutoMigrate(&models.File{})

	SendSeeder(db)
	return
}

func SendSeeder(db *gorm.DB) (err error) {
	bucket := []models.Bucket{
		{
			Name: "gulagulali",
		},
		{
			Name: "skincare-telto",
		},
	}

	db.Create(&bucket)

	path := []models.Path{
		{
			ID:         uuid.New(),
			Name:       "photo",
			BucketName: "gulagulali",
		},
		{
			ID:         uuid.New(),
			Name:       "photo",
			BucketName: "skincare-telto",
		},
	}

	db.Create(&path)
	return
}
