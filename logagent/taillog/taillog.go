package taillog

import (
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail
)

// Init 初使化
func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	tailObj, err = tail.TailFile(fileName, config)

	return
}

// ReadChan 读取日志
func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
