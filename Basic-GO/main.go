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
	router.GET("/api/posts/:id", controllers.GetPostById)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePostById)

	router.GET("/api/todos", controllers.AllTodo)
	router.POST("/api/todos", controllers.StoreTodo)
	router.GET("/api/todos/:id", controllers.GetTodoById)
	router.PUT("/api/todos/:id/done", controllers.MarkDoneTodo)
	router.DELETE("/api/todos/:id", controllers.DeleteTodoById)

	router.Run(":3000")
}
