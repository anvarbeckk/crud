package controllers

import (
	"crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostInput struct {
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}


type UpdatePostInput struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

// create
func CreatePost(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{Title: input.Title, Content: input.Content}
	models.DB.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}
// read
func FindPosts(c *gin.Context) {
	var posts []models.Post


	models.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"data": posts})
}
// read by id
func FindPost(c * gin.Context) {
	var post []models.Post

	if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

	models.DB.Find(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// update
func UpdatePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input UpdatePostInput

    if err := c.ShouldBindJSON(&input); err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedPost := models.Post{Title: input.Title, Content: input.Content}

    models.DB.Model(&post).Updates(&updatedPost)
    c.JSON(http.StatusOK, gin.H{"data": post})
}
// delete
func DeletePost(c *gin.Context) {
    var post models.Post
    if err := models.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
        c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    models.DB.Delete(&post)
    c.JSON(http.StatusOK, gin.H{"data": "success"})
}
