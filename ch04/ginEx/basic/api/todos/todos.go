package todos

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func All(ctx *gin.Context) {
	todoItems := make([]string, 5)
	todoItems[0] = "123"
	todoItems[1] = "456"
	ctx.JSON(http.StatusOK, gin.H{
		"todos": todoItems,
	})
}

func One(ctx *gin.Context)  {

}