package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 专门往kafka写日志的模块

var (
	client sarama.SyncProducer //声明一个全局连接kafka的生产者client
)

// Init 初使化
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)

	return
}

// SendToKafka 发送消息
func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("pid: %v offset: %v\n", pid, offset)
}

func Close() {
	client.Close()
}
