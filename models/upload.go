package models

type ResponseUpload struct {
	BucketName string `json:"bucket"`
	PathName   string `json:"path"`
	File       any    `json:"file"`
}
