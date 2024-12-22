package controllers

import (
	"net/http"
	"tusk/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (*UserController) Login(c *gin.Context) {
	user := models.User{}
	errBindJson := c.ShouldBindJSON(&user)
	if errBindJson != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":errBindJson.Error()})
		return;
	}
	password := user.Password
	errDb:=u.DB.Where("email = ?", user.Email).Take(&user).Error
	if errDb != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":errDb.Error()})
		return;
	}
	errCompare := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errCompare != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":errCompare.Error()})
		return;
	}
	c.JSON(http.StatusOK, user)
}