package routers

import (
	"todoapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
    router.GET("/todos", controllers.GetTodos)
	router.GET("/todo/:pk", controllers.GetTodoByPK)
    router.POST("/todos", controllers.CreateTodo)
	router.GET("/welcome", controllers.Welcome)
}
