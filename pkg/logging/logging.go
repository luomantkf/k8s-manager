package logging

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)

	//添加获取调用行数的hook
	log.AddHook(NewContextHook())
}

//返回日志{}填充符字符串
func toString(args []string) string {
	if len(args) == 0 {
		return ""
	}

	c := strings.Count(args[0], "{}")
	var i int
	for i = 1; i <= c && i < len(args); i++ {
		args[0] = strings.Replace(args[0], "{}", args[i], 1)
	}

	for j := i; j < len(args); j++ {
		args[0] = args[0] + "," + args[j]
	}

	return args[0]
}

func Trace(args ...string) {
	log.Trace(toString(args))
}

func Debug(args ...string) {
	log.Debug(toString(args))
}

func Print(args ...string) {
	log.Print(toString(args))
}

func Info(args ...string) {
	log.Info(toString(args))
}

func Warn(args ...string) {
	log.Warn(toString(args))
}

func Warning(args ...string) {
	log.Warning(toString(args))
}

func Error(args ...string) {
	log.Error(toString(args))
}

func Fatal(args ...string) {
	log.Fatal(toString(args))
}

func Panic(args ...string) {
	log.Panic(toString(args))
}
