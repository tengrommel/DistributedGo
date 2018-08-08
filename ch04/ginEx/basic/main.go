package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"DistributedGo/ch04/ginEx/basic/api/todos"
	"DistributedGo/ch04/ginEx/basic/middlewares"
)

func main() {
	app := gin.Default()
	app.GET("/hello/:name", hello)
	todosGroup := app.Group("/api/todos")
	todosGroup.GET("/", middlewares.Log , todos.All)
	todosGroup.GET("/:id")
	todosGroup.PUT("/:id")
	todosGroup.POST("/")
	todosGroup.DELETE("/:id")
	fmt.Println(todos.ABC)
	app.Run(":8081")
}

func hello(ctx *gin.Context)  {
	name := ctx.Param("name")
	page := ctx.Query("page")
	fmt.Println(page)
	ctx.String(http.StatusOK, "hello %s", name)
}