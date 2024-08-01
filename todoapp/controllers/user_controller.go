package controllers

import (
	"net/http"
	"todoapp/database"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)


func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	var user models.User
	pk := c.Param("pk")
	database.DB.First(&user, pk)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&user)
    c.JSON(http.StatusCreated, user)
}
