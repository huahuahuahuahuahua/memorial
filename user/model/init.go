package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//数据库连接单例
var DB *gorm.DB

//Database在中间件中初始化mysql连接
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	//设置日志输出
	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	//默认不加复数
	db.SingularTable(true)
	//设置数据池
	//空闲时最高数目
	db.DB().SetMaxIdleConns(20)
	//打开时最高数目
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
