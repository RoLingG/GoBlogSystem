package log_stash

import "encoding/json"

type LogLevel int

const (
	DebugLevel   LogLevel = 1
	InfoLevel    LogLevel = 2
	WarningLevel LogLevel = 3
	ErrorLevel   LogLevel = 4
)

// MarshalJSON 将LogLevel类型的MarshalJSON方法覆写
func (level LogLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}

// 将LogLevel类型的String方法覆写
func (level LogLevel) String() string {
	var str string
	switch level {
	case DebugLevel:
		str = "Debug"
	case InfoLevel:
		str = "Info"
	case WarningLevel:
		str = "Warning"
	case ErrorLevel:
		str = "Error"
	default:
		str = "未知日志等级"
	}
	return str
}
