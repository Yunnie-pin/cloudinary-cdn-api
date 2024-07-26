package main

// Import the required packages for upload and admin.

import (
    "context"
    "log"

    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api"
    "github.com/cloudinary/cloudinary-go/v2/api/admin"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func main() {
	// Add your Cloudinary product environment credentials.
	// cld, _ := cloudinary.NewFromParams("<your-cloud-name>", "<your-api-key>", "<your-api-secret>")
cld, _ := cloudinary.NewFromParams("ddskzoj32", "958337737332423", "1fio5Kx7MgqCEtOmwgmJJg2UHDI")
}
