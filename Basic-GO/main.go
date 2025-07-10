package main

import (
	"backend-api/controllers"
	"backend-api/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	models.ConnectDB()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	router.Static("/public", "./public")
	router.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
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
