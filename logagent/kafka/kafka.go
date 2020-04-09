package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer //声明一个全局连接kafka的生产者client
	logDataChan chan *logData
)

// Init 初使化
func Init(addrs []string, maxSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)

	logDataChan = make(chan *logData, maxSize)

	// 开启后台的goroutine从通道中获取数据发给kafka
	go sendToKafka()
	return
}

// SendToKafka 发送消息
func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{
				Topic: ld.topic,
				Value: sarama.StringEncoder(ld.data),
			}
			pid, offset, err := client.SendMessage(msg)
			if err != nil {
				panic(err)
			}
			fmt.Printf("pid: %v offset: %v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// SendToChan 把数据放到channel中
func SendToChan(topic, data string) {
	msg := &logData{
		topic,
		data,
	}
	logDataChan <- msg
}

func Close() {
	client.Close()
}
