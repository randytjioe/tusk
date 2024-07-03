package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Tusk API")
	})
	router.Static("/attachments", "./attachments")
	router.Run("10.10.103.254:8080")
}
