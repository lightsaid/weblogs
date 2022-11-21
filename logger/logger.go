package logger

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"lightsaid.com/weblogs/global"
)

// InitLogger 简单初始化日志
func InitLogger() {
	// 设置日志输出和分割
	logFile := &lumberjack.Logger{
		Filename:   global.Config.LogOutput,
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28,    //days
		Compress:   false, // disabled by default
	}

	// 设置多输出
	var mw io.Writer
	if global.Config.Mode == "dev" {
		mw = io.MultiWriter(os.Stdout, logFile)
	} else {
		mw = io.MultiWriter(logFile)
	}

	logrus.SetOutput(mw)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logLevel, err := logrus.ParseLevel(global.Config.LogLevel)
	if err != nil {
		log.Fatal("logrus.ParseLevel error: " + err.Error())
		// logLevel = logrus.InfoLevel
	}
	logrus.SetLevel(logLevel)
}

// TODO:
// func useHookEmail() {
// logrus.AddHook()
// }
