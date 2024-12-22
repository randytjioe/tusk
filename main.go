package main

import (
	"net/http"
	"tusk/config"
	"tusk/controllers"
	"tusk/models"

	"github.com/gin-gonic/gin"
)

func main() {
	//Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{},&models.Task{})
	config.CreateOwnerAccount(db)

	// Controller
	userController := controllers.UserController{DB: db}
	//Router
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Tusk API")
	})

	router.POST("/users/login",userController.Login())
	router.Static("/attachments", "./attachments")
	router.Run("10.10.103.255:8080")
}
