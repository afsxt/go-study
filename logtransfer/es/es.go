package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

// 初始化ES，准备接收kafka那边发来的数据

// LogData ...
type LogData struct {
	Topic string `json:"topic"`
	Type  string `json:"type"`
	Data  string `json:"data"`
}

var (
	client *elastic.Client
	ch     chan *LogData
)

// Init ...
func Init(address string, chanSize int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	ch = make(chan *LogData, chanSize)

	go sendToES()
	return
}

func SendToChan(msg *LogData) {
	ch <- msg
}

// SendToES 发送数据到ES
func sendToES() {
	for {
		select {
		case msg := <-ch:
			put, err := client.Index().
				Index(msg.Topic).
				Type(msg.Type).
				BodyJson(msg).
				Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("indexed user %s to index %s, type %s\n", put.Id, put.Index, put.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
