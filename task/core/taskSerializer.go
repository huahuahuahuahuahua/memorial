package core

import (
	model "code.huawink.com/beiwanglu/task/model"
	services "code.huawink.com/beiwanglu/task/services"
)

//通过services里面建立的用户模型进行初始化。
func BuildTask(item model.Task) *services.TaskModel {
	taskModel := services.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}
