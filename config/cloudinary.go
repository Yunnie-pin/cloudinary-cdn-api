package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func ConnectCld(config *Config) (cld *cloudinary.Cloudinary, err error) {
	cld, err = cloudinary.NewFromParams(config.CDNCloudName, config.CDNAPIKey, config.CDNAPISecret)

	if err != nil {
		return cld, err
	}

	return cld, err
}
