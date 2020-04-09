package taillog

import (
	"context"
	"fmt"
	"logagent/kafka"

	"github.com/hpcloud/tail"
)

// 专门从日志文件收集日志的模块

var (
	tailObj *tail.Tail
)

type TailTask struct {
	path       string
	topic      string
	instance   *tail.Tail
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	err := tailObj.init()
	if err != nil {
		panic(err)
	}
	return
}

func (t *TailTask) init() (err error) {
	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	t.instance, err = tail.TailFile(t.path, config)

	go t.run()
	return
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task: %s结束\n", t.path)
			return
		case line := <-t.instance.Lines:
			// kafka.SendToKafka(t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
		}
	}
}
