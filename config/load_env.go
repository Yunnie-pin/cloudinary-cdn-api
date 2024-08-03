package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV          string
	PORT         string
	CDNCloudName string
	CDNAPIKey    string
	CDNAPISecret string

	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBURL      string
}

func LoadConfig() (config Config, err error) {
	godotenv.Load()

	config = Config{
		ENV:          os.Getenv("ENV"),
		PORT:         os.Getenv("PORT"),
		CDNCloudName: os.Getenv("CDN_CLOUD_NAME"),
		CDNAPIKey:    os.Getenv("CDN_API_KEY"),
		CDNAPISecret: os.Getenv("CDN_API_SECRET"),
		DBHost:       os.Getenv("DB_HOST"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		DBPort:       os.Getenv("DB_PORT"),
		DBURL:        os.Getenv("DB_URL"),
	}

	return config, err
}
