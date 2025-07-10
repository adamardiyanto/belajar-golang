package controllers

import (
	"backend-api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidateTodoInput struct {
	Title string `json:"title" binding:"required"`
}

func AllTodo(c *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	c.JSON(200, gin.H{
		"success": true,
		"message": "List All Todos",
		"data":    todos,
	})
}

func StoreTodo(c *gin.Context) {
	var input ValidateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	todo := models.Todo{
		Title: input.Title,
	}

	models.DB.Create(&todo)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo Created Successfully",
		"data":    todo,
	})
}

func GetTodoById(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not Found"})
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo with ID : " + c.Param("id"),
		"data":    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "todo not Found"})
		return
	}

	var input ValidateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), GetErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	models.DB.Model(&todo).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo Updated Successfully",
		"data":    todo,
	})
}

func DeleteTodoById(c *gin.Context) {
	var todo models.Todo
	if err := models.DB.Where("id=?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todo not Found"})
		return
	}

	models.DB.Delete(&todo)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Todo Deleted Successfully",
	})
}
