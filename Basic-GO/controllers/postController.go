package controllers

import (
	"backend-api/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func AllPosts(c *gin.Context) {
	var posts []models.Post
	models.DB.Find(&posts)

	c.JSON(200, gin.H{
		"success": true,
		"message": "List All Posts",
		"data":    posts,
	})
}

func StorePost(c *gin.Context) {
	var input ValidatePostInput
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

	post := models.Post{
		Title:   input.Title,
		Content: input.Content,
	}
	models.DB.Create(&post)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Created Successfully",
		"data":    post,
	})
}

func GetPostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not Found"})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Detail Post with ID : " + c.Param("id"),
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not Found"})
		return
	}

	var input ValidatePostInput
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

	models.DB.Model(&post).Updates(input)

	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Updated Successfully",
		"data":    post,
	})
}

func DeletePostById(c *gin.Context) {
	var post models.Post
	if err := models.DB.Where("id=?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not Found"})
		return
	}

	models.DB.Delete(&post)
	c.JSON(200, gin.H{
		"success": true,
		"message": "Post Deleted Successfully",
	})
}
