package kafka

import (
	"fmt"
	"logtransfer/es"

	"github.com/Shopify/sarama"
)

// 初始化kafka消费者 从kafka取数据发往ES

// Init 初始化
func Init(addres []string, topic string) (err error) {
	consumer, err := sarama.NewConsumer(addres, nil)
	if err != nil {
		return
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		return
	}
	fmt.Println("分区列表：", partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		go func(sarama.PartitionConsumer) {
			defer pc.AsyncClose()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d key:%v value:%v\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				// ld := map[string]interface{}{
				// 	"data": string(msg.Value),
				// }
				ld := &es.LogData{
					Data:  string(msg.Value),
					Topic: topic,
					Type:  "go",
				}
				es.SendToChan(ld)
			}
		}(pc)
	}
	return
}
