package main

import (
	"backend-api/controllers"
	"backend-api/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.ConnectDB()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "HELLLOOOO",
		})
	})

	//membuat route get all posts
	router.GET("/api/posts", controllers.AllPosts)

	router.POST("/api/posts", controllers.StorePost)

	router.Run(":3000")
}
