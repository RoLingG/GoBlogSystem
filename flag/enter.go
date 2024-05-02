package flag

import sysflag "flag"

// 迁移表结构的脚手架
type Option struct {
	DB bool
}

// 解析命令行数
func Parse() Option {
	db := sysflag.Bool("db", false, "初始化数据库")
	//解析命令行参数写入注册的flag里
	sysflag.Parse()
	return Option{
		DB: *db,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) bool {
	if option.DB {
		return true
	}
	return false
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
