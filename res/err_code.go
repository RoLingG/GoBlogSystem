package res

type ErrorCode int

const (
	SettingsError = 1001 //定义settings系统错误
	ArgumentError = 1002 //参数错误
)

// 这里可以从json文件里面读
var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		ArgumentError: "参数错误",
	}
)
