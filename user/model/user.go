package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	/*gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	*/
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

//密码加密难度
const PassWordCost = 12

//bcrypt 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

//bcrypt 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
