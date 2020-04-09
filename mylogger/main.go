package main

import (
	"mylogger/mylogger"
)

func main() {
	// mylog := mylogger.NewConsoleLogger("info")
	mylog := mylogger.NewFileLogger("Info", "./", "test.log", 10*1024*1024)
	mylog.Debug("这是一条Debug日志")
	mylog.Info("这是一条Info日志")
	mylog.Warning("这是一条Warning日志, %d", 1111)
	mylog.Error("这是一条Error日志")
	mylog.Fatal("这是一条Fatal日志")
}
