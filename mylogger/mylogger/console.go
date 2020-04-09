package mylogger

// 往终端写日志相关内容

import (
	"fmt"
	"time"
)

// Logger 结构
type ConsoleLogger struct {
	level LogLevel
}

// NewConsoleLogger 构造
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		level,
	}
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now().Format("2006-01-02 15:04:05")
		funcName, fileName, line := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now, getLogString(lv), fileName, funcName, line, msg)
	}
}

func (l ConsoleLogger) enable(level LogLevel) bool {
	return l.level <= level
}

// Debug ...
func (c ConsoleLogger) Debug(format string, a ...interface{}) {
	c.log(DEBUG, format, a)
}

// Info ...
func (c ConsoleLogger) Info(format string, a ...interface{}) {
	c.log(INFO, format, a...)
}

// Warning ...
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error ...
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal ...
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
