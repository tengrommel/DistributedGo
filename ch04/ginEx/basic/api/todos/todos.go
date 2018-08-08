package todos

import (
	"github.com/gin-gonic/gin"
	"DistributedGo/ch04/ginEx/basic/libs/database"
	"DistributedGo/ch04/ginEx/basic/models"
	"net/http"
	"fmt"
)

func All(ctx *gin.Context) {
	db := database.Open()
	defer db.Close()
	var todoItems []models.Todo
	if err := db.Find(&todoItems).Error; err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, todoItems)
}

func One(ctx *gin.Context)  {
	id := ctx.Param("id")
	db := database.Open()
	defer db.Close()
	var todoItem models.Todo
	if err := db.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err != nil{
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, todoItem)
}

type postCreate struct {
	Title string `json:"title" binding:"required"`

}

func Create(ctx *gin.Context)  {
	var postData postCreate
	if err := ctx.BindJSON(&postData); err != nil{
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db := database.Open()
	defer db.Close()
	todoItem := models.Todo{
		Title: "Hello",
	}
	if err := db.Create(&todoItem).Error; err != nil{
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": todoItem.ID})
}

func Update(ctx *gin.Context)  {
	id := ctx.Param("id")
	db := database.Open()
	defer db.Close()
	var todoItem models.Todo
	if err := db.Model(&todoItem).Where("id = ?", id).First(&todoItem).Error; err != nil{
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	todoItem.Title = "World"
	if err := db.Save(&todoItem).Error; err != nil{
		fmt.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}
	// 另外一种更新
	//if err := db.Model(&todoItem).Where("id = ?", id).Update("title", "direct update yo!").Error; err != nil{
	//	fmt.Println(err)
	//	ctx.Status(http.StatusInternalServerError)
	//	return
	//}
	ctx.Status(http.StatusOK)
}