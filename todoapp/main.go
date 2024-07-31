package main

import (
	"todoapp/database"
	"todoapp/routers"

	"github.com/gin-gonic/gin"
)

func main() {
    database.Init()
    router := gin.Default()
    routers.SetupRouter(router)
    router.Run(":8080")
}
