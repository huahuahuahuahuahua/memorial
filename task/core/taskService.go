package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"code.huawink.com/beiwanglu/task/model"
	services "code.huawink.com/beiwanglu/task/services"
	"github.com/streadway/amqp"
)

//创建备忘录，将备忘录信息标记生产，放到rabbitMQ消息队列中
func (*TaskService) CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	fmt.Println("接受到了createTask的请求")
	if err != nil {
		err = errors.New("rabbitMQ channel err;" + err.Error())
		return err
	}
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)
	//json格式化
	body, _ := json.Marshal(req)
	//发布消息
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbitMQ publish err:" + err.Error())
		return err
	}
	return nil
}

//实现备忘录服务接口 获取备忘录列表
func (*TaskService) GetTasksList(ctx context.Context, req *services.TaskRequest, resp *services.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count uint32
	//查询数据库 查找备忘录
	err := model.DB.Offset(req.Start).Limit(req.Limit).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find : " + err.Error())
	}
	// 统计数量
	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	// 返回proto中定义的模型
	var taskRes []*services.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = count
	return nil
}

//获取详细的备忘录
func (*TaskService) GetTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	//第一条记录
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}

//修改备忘录
func (*TaskService) UpdateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	//查找用户该信息
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Save(&taskData)
	resp.TaskDetail = BuildTask(taskData)
	return nil
}

//删除备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	err := model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Delete(&model.Task{}).Error
	if err != nil {
		return errors.New("删除失败:err" + err.Error())

	}
	return nil
}
