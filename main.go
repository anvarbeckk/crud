package main

import (
	"crud/controllers"
	"crud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()
	router.POST("/posts", controllers.CreatePost)
	router.Run("localhost:8080")
}