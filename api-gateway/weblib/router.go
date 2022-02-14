package weblib

import (
	"code.huawink.com/beiwanglu/api-gateway/weblib/handlers"
	"code.huawink.com/beiwanglu/api-gateway/weblib/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	//跨域、服务处理和错误中间件
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	//使用cookie和session
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	//网络ping验证
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
	}
	//用户服务
	v1.POST("/user/register", handlers.UserRegister)
	v1.POST("/user/login", handlers.UserLogin)

	authed := v1.Group("/")
	//登录保护，使用jwt中间件
	authed.Use(middleware.JWT())
	{
		//restful路由中间件处理
		authed.GET("tasks", handlers.GetTaskList)
		authed.POST("task", handlers.CreateTask)
		authed.GET("task/:id", handlers.GetTaskDetail) // task_id
		authed.PUT("task/:id", handlers.UpdateTask)    // task_id
		authed.DELETE("task/:id", handlers.DeleteTask) // task_id
	}
	return ginRouter
}
