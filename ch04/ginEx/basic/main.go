package main

import (
	"github.com/gin-gonic/gin"
	"DistributedGo/ch04/ginEx/basic/api/todos"
)

func main() {
	app := gin.Default()
	app.GET("/api/todos", todos.All)
	app.GET("/api/todos/:id", todos.One)
	app.POST("/api/todos", todos.Create)
	app.PUT("/api/todos/:id", todos.Update)
	app.Run()
}
