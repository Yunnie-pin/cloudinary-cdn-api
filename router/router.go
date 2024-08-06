package router

import (
	"claudinary-cdn-api/controllers"
	"claudinary-cdn-api/middlewares"

	"github.com/gin-gonic/gin"
)

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
		c.HTML(200, "index.html", nil)
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
