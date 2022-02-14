package model

import "github.com/jinzhu/gorm"

//定义任务结构体，用户信息以及内容，状态
type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    int    `gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64 `gorm:"default:'0'"`
}
