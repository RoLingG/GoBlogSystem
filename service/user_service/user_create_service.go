package user_service

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/utils"
	"GoRoLingG/utils/pwd"
	"errors"
)

const Avatar = "/upload/avatar/avatar.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		return errors.New("用户名已存在")
	}
	// 对密码进行hash
	hashedPwd := pwd.HashPwd(password)

	addr := utils.GetAddr(ip)
	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashedPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar, //使用默认头像
		IP:         ip,
		Address:    addr,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
