package handlers

import (
	"context"
	"fmt"
	"net/http"

	"code.huawink.com/beiwanglu/api-gateway/pkg/utils"

	services "code.huawink.com/beiwanglu/api-gateway/services"

	"github.com/gin-gonic/gin"
)

func UserRegister(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
}

func UserLogin(ginCtx *gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	if err != nil {
		fmt.Println(err)
	}
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
