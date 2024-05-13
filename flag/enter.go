package flag

import sysflag "flag"

// 迁移表结构的脚手架
type Option struct {
	DB   bool
	User string //-u admin 创建root用户 -u user 创建普通用户
}

// 解析命令行数
func Parse() Option {
	db := sysflag.Bool("db", false, "初始化数据库")
	user := sysflag.String("u", "", "创建用户")
	//解析命令行参数写入注册的flag里
	sysflag.Parse()
	return Option{
		DB:   *db,
		User: *user,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return true
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" || option.User == "user" {
		CreateUser(option.User)
		return
	}
	sysflag.Usage() //有而外的内容则直接不生效
}
