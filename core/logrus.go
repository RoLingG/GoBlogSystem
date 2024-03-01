package core

import (
	"GoRoLingG/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct {
}

// Format 实现fomatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的等级去取不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b = &bytes.Buffer{}
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	//自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcValue := entry.Caller.Function
		fileValue := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileValue, funcValue, entry.Message)
	} else {
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s \n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

// 初始化logger
func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                //新建一个实例
	mLog.SetOutput(os.Stdout)                           //设置一个输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号,ShowLine默认为true，在yaml里写的
	mLog.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	//mLog.SetLevel(logrus.DebugLevel)   //设置最低的等级，level是iota，为uint32类型，所以可以转化
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mLog.SetLevel(level)
	InitDefaultLogger()
	return mLog
}

// 初始化全局logger
func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)                           //设置一个输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{})                  //设置自己定义的Formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的等级
}
