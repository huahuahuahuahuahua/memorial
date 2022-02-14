package handlers

import (
	"context"
	"strconv"

	"code.huawink.com/beiwanglu/api-gateway/pkg/utils"
	services "code.huawink.com/beiwanglu/api-gateway/services"
	"github.com/gin-gonic/gin"
)

//获取任务清单
func GetTaskList(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	//抛出错误
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin keys中取出用户实例
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("authorization"))
	taskReq.Uid = uint64(claim.Id)
	//调用服务端的函数
	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	//返回json数据，gin.H是map[]string 格式
	ginCtx.JSON(200, gin.H{
		"data": gin.H{
			"task":  taskResp.TaskList,
			"count": taskResp.Count,
		},
	})
}

//创建任务
func CreateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

//获取用户详情
func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.BindUri(&taskReq))
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id")) // 获取task_id
	taskReq.Id = uint64(id)
	productService := ginCtx.Keys["taskService"].(services.TaskService)
	productRes, err := productService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{"data": productRes.TaskDetail})
}

//更新任务
func UpdateTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	taskReq.Uid = uint64(claim.Id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}

//删除任务
func DeleteTask(ginCtx *gin.Context) {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	//从gin.keys取出服务实例
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)
	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	taskRes, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(200, gin.H{"data": taskRes.TaskDetail})
}
