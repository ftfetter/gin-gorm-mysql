package routes

import (
	"gin-gorm-mysql/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("users", GetUsers)
	r.POST("users", CreateUser)
	r.GET("users/:id", GetUserById)
	r.PUT("users/:id", UpdateUser)
	r.DELETE("users/:id", DeleteUser)
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

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, id)
}
