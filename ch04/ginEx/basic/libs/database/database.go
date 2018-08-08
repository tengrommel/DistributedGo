package database

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/mosluce/go-toolkits"
	"github.com/mosluce/go-toolkits/database"
	"DistributedGo/ch04/ginEx/basic/models"
)

func Open() *gorm.DB {
	var db *gorm.DB
	var err error
	if gin.Mode() == gin.ReleaseMode {
		db, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:database.SQLITE,
			Filepath: "db.sqlite",
		})
	}else {
		db, err = toolkits.OpenDB(database.ConnectionConfig{
			Dialect:database.SQLITE,
			Filepath: "db.sqlite",
		})
		db.LogMode(true)
	}
	if err != nil{
		panic(err)
	}
	db.AutoMigrate(&models.Todo{})
	return db
}

