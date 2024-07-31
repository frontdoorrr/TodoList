package controllers

import (
	"net/http"
	"todoapp/database"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
    var todos []models.Todo
    database.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&todo)
    c.JSON(http.StatusCreated, todo)
}
