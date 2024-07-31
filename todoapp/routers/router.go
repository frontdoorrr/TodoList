package routers

import (
	"todoapp/controllers"
	_ "todoapp/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine){
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)
}
