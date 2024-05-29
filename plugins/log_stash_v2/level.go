package log_stash_v2

import "encoding/json"

type LogLevel int

const (
	Info    LogLevel = 1
	Warning LogLevel = 2
	Error   LogLevel = 3
)

// MarshalJSON 将LogLevel类型的MarshalJSON方法覆写
func (level LogLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(level.String())
}

// 将LogLevel类型的String方法覆写
func (level LogLevel) String() string {
	var str string
	switch level {
	case Info:
		str = "info"
	case Warning:
		str = "warning"
	case Error:
		str = "error"
	default:
		str = ""
	}
	return str
}
