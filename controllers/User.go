package controllers

import (
	"gin-gorm-mysql/business"
	"gin-gorm-mysql/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FetchUsers(c *gin.Context) {
	var users []models.User
	if err := business.FetchUsers(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	if err := business.CreateUser(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func FetchUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	var user models.User
	if err := business.FetchUserById(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	var user models.User
	c.BindJSON(&user)

	if err = business.UpdateUser(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID."})
		return
	}

	var user models.User
	if err = business.DeleteUser(id, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	}
}
