package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//接受服务实例，并存到gin.key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		//将实例保存在key中
		c.Keys = make(map[string]interface{})
		c.Keys["userService"] = service[0]
		c.Keys["taskService"] = service[1]
		c.Next()
	}
}

//错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(200, gin.H{
					"code": "404",
					"msg":  fmt.Sprintf("%s", r),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
