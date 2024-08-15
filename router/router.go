package router

import (
	"claudinary-cdn-api/controllers"
	"claudinary-cdn-api/middlewares"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var repoData map[string]interface{} = nil
var contribData []map[string]interface{} = nil

func NewRouter(
	bucketController *controllers.BucketController,
	pathController *controllers.PathController,
	fileController *controllers.FileController,
	uploaderController *controllers.UploaderController,
) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("view/index.html")

	//web
	r.GET("/", func(c *gin.Context) {
		client := &http.Client{}
		method := "GET"

		if repoData == nil {
			url := "https://api.github.com/repos/Yunnie-pin/cloudinary-cdn-api"
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			defer res.Body.Close()
			body, err := io.ReadAll(res.Body)

			repoData = make(map[string]interface{})
			_ = json.Unmarshal(body, &repoData)

			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		}
		if contribData == nil {
			url := "https://api.github.com/repos/Yunnie-pin/cloudinary-cdn-api/contributors"
			req, err := http.NewRequest(method, url, nil)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)

			contribData = make([]map[string]interface{}, 0)
			_ = json.Unmarshal(body, &contribData)

			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		}

		c.HTML(200, "index.html", gin.H{
			"repo":         repoData,
			"contributors": contribData,
		})
	})

	r.Use(middlewares.CORSMiddleware())

	//api
	apiRouter := r.Group("/api")
	{
		apiRouter.Use(middlewares.SimpleAuth())

		apiRouter.POST("/uploader", uploaderController.UploadFile)

		bucketRouter := apiRouter.Group("/buckets")
		{
			bucketRouter.GET("/", bucketController.FindAll)
			bucketRouter.POST("/", bucketController.CreateBucket)
		}

		pathRouter := apiRouter.Group("/path")
		{
			pathRouter.POST("/", pathController.CreatePath)
		}

		filesRouter := apiRouter.Group("/files")
		{
			filesRouter.GET("/:bucket/*path", fileController.FindAllFiles)
			filesRouter.DELETE("/:bucket/*path", fileController.DeleteFile)
		}

	}

	return r
}
