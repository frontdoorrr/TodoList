package controllers

import (
	"net/http"
	"todoapp/database"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetTodos(c *gin.Context) {
    var todos []models.Todo
    database.DB.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

// CreateTodo godoc
// @Summary Create a todo
// @Description Create a new todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body models.Todo true "Todo"
// @Success 201 {object} models.Todo
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
    var todo models.Todo
    if err := c.BindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    database.DB.Create(&todo)
    c.JSON(http.StatusCreated, todo)
}
