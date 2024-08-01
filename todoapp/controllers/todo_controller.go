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

func GetTodoByPK(c *gin.Context) {
	var todo models.Todo
	pk := c.Param("pk")
	if err := database.DB.First(&todo, pk).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Todo Not Found"})
	}

	database.DB.First(&todo, pk)
	c.JSON(http.StatusOK, todo)
}

func Welcome(c *gin.Context){
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
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
