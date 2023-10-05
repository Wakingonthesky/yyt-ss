package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"guanxingtuan_bck/global"
	"os"
	"path"
)

// 颜色
const (
	red    = 31
	yellow = 33
	blue   = 36
	grey   = 37
)

type Logformatter struct{}

func (t *Logformatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level展示不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = grey
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	pre := global.Config.Logger.Prefix
	//自定义日期格式
	timestamp := entry.Time.Format("2023-10-02 09:42:05")
	if entry.HasCaller() {
		//自定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d\n", path.Base(entry.Caller.File), entry.Caller.Line)
		//自定义输出格式
		fmt.Fprintf(b, "%s:[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", pre, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s:[%s] \x1b[%dm[%s]\x1b[0m %s \n", pre, timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mlog := logrus.New()                                //创建实例
	mlog.SetOutput(os.Stdout)                           //设置输出类型
	mlog.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数和哈行号
	mlog.SetFormatter(&Logformatter{})                  //设置自己定义的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	mlog.SetLevel(level) //设置最低level
	InitDefaultLogger()
	return mlog
}

func InitDefaultLogger() {
	//全局log
	logrus.SetOutput(os.Stdout)                           //设置输出类型
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //开启返回函数名和行号
	logrus.SetFormatter(&Logformatter{})                  //设置自己的formatter
	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level) //设置最低的level
}
