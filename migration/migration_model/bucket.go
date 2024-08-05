package migration_model

type Bucket struct {
	Name string `gorm:"primary_key;type:varchar(255);not null" json:"name"`
}
