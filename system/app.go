package system

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	var root = gin.Default()
	root.POST("/create", CreateToDo)
	root.GET("/read", ReadToDo)
	root.PUT("/update/:id", UpdateToDo)
	root.DELETE("/delete/:id", DeleteToDo)
	root.Run(":5000")
}
