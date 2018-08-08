package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Log(ctx *gin.Context)  {
	fmt.Println("I am a middleware!!!!!")
	ctx.Next()
}
