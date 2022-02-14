package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func Database(connString string) {
	db, err := gorm.Open("mysql", connString)

	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
}
