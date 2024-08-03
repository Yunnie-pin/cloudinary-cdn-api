package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartDB(config *Config) (db *gorm.DB, err error) {
	var dsn string

	if config.ENV == "production" || config.ENV == "PRODUCTION" {
		dsn = config.DBURL
	} else {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
	}

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	fmt.Println("🧊🧊🧊 Successfully connected to the database!")

	return db, err
}
