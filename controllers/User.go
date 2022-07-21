package controllers

import (
	models "gin-gorm-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func GetUsers(c *gin.Context) {
	var users []models.User
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	c.JSON(http.StatusOK, user)
}
