package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
)

// InitKafkaConsumer 初始化kafkaConsumer
func (s *Service) InitKafkaConsumer(cfg *config.Config, consumer sarama.Consumer) {
	partitionConsumer, err := consumer.ConsumePartition(cfg.Kafka.Topic, 0, sarama.OffsetNewest) // 根据topic取到所有的分区
	if err != nil {
		panic("error get consumer")
	}

	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			go func() {
				//TODO 未处理错误,应该把错误放在logger
				redis.SetTotalCount()
				fmt.Println(msg.Value)
				s.KfkChan <- msg.Value
				fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
			}()
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}
