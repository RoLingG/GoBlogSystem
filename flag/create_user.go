package flag

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/utils/pwd"
	"bufio"
	"fmt"
	"os"
)

func CreateUser(permission string) {
	// 创建用户的逻辑
	//用户名 昵称 密码 确认密码 邮箱
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)
	fmt.Printf("请输入用户名：")
	fmt.Scan(&userName)
	fmt.Printf("请输入用户昵称：")
	fmt.Scan(&nickName)
	fmt.Printf("请输入用户密码：")
	fmt.Scan(&password)
	fmt.Printf("请重新确认用户密码：")
	fmt.Scan(&rePassword)
	//fmt.Printf("请输入用户邮箱：")	这里如果直接这样写的话，会发现程序直接输入完rePassword之后就跳过了输入email这一个步骤，这是因为我们为了可以不输入空邮箱就能创建用户使用了Scanln()导致的
	//fmt.Scanln(&email)			Scan()在读取确认密码时，可能会读取到输入中的换行符，这导致Scanln()在尝试读取邮箱时，实际上已经读取到了之前输入的换行符，因此它认为已经读取了一行，并不会等待用户输入。
	// 清除输入缓冲区中的换行符
	bufio.NewReader(os.Stdin).ReadByte()
	fmt.Printf("请输入用户邮箱：") //要用的话，只能清空缓冲区的换行符再使用fmt.Scanln()
	fmt.Scanln(&email)

	//判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Debug().Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		//用户已存在
		global.Log.Error("用户已存在，请修改用户名后重新创建用户")
		return
	}
	if password != rePassword {
		global.Log.Error("两次密码不相同，请重新确认密码")
		return
	}
	//对密码进行加密
	hashedPwd := pwd.HashPwd(password)

	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	//头像解决
	//方案一：默认头像
	userAvatar := "/upload/avatar/avatar.png"
	//方案二：随机选择头像

	//入库
	err = global.DB.Debug().Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashedPwd,
		Email:      email,
		Avatar:     userAvatar,
		IP:         "127.0.0.1",
		Address:    "内网",
		Role:       role,
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", userName)
}
