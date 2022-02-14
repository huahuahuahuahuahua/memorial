package core

import (
	"context"
	"errors"

	"code.huawink.com/beiwanglu/user/model"
	services "code.huawink.com/beiwanglu/user/services"
	"github.com/jinzhu/gorm"
)

//通过services里面建立的用户模型进行初始化。
func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	//返回建立好的用户模型
	return &userModel
}

//用户登录验证
//UserService是pkg里面建立的结构体
func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	//查询数据库是否存在用户
	if err := model.DB.Where("user_name=?", req.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	//检验密码
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}
	//返回用户具体信息，用
	resp.UserDetail = BuildUser(user)
	return nil
}

//用户注册
func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	count := 0
	//查询用户名称数目
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	//建立用户
	user := model.User{
		UserName: req.UserName,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	//返回用户具体信息
	resp.UserDetail = BuildUser(user)
	return nil
}
