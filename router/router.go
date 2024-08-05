package router

import (
	"claudinary-cdn-api/controllers"
	"claudinary-cdn-api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	bucketController *controllers.BucketController,
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
		apiRouter.POST("/uploader", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "success",
			})
		})

		bucketRouter := apiRouter.Group("/buckets")
		{
			bucketRouter.GET("/", bucketController.FindAll)
		}

		filesRouter := apiRouter.Group("/files")
		{
			filesRouter.GET("/:bucket/*path", func(c *gin.Context) {
				bucket := c.Param("bucket")
				path := c.Param("path")

				c.JSON(200, gin.H{
					"message": "success",
					"bucket":  bucket,
					"path":    path,
				})
			})
		}

	}

	return r
}
