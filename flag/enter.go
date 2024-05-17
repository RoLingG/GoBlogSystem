package flag

import (
	sysflag "flag"
	"github.com/fatih/structs"
)

// 迁移表结构的脚手架
type Option struct {
	DB   bool
	User string //-u admin 创建root用户 -u user 创建普通用户
	ES   string //-es create 创建表索引 -es delete 删除表索引
}

// 解析命令行数
func Parse() Option {
	db := sysflag.Bool("db", false, "初始化数据库")
	user := sysflag.String("u", "", "创建用户")
	es := sysflag.String("es", "", "ES处理")
	//解析命令行参数写入注册的flag里
	sysflag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		ES:   *es,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	//将命令行参数map化
	maps := structs.Map(&option)
	//用for循环maps去检测命令行参数类型，从而进行情况选择
	//(一般来说maps里只有一个参数，因为命令行一般执行的都是单参数操作，用for循环检测多是为了以后多参数也可以用)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val != false {
				f = true
			}
		}
	}
	return f
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
	if option.ES == "create" {
		EsCreateIndex()
		return
	} else if option.ES == "delete" {
		EsRemoveIndex()
		return
	}
	sysflag.Usage() //有而外的内容则直接不生效
}
