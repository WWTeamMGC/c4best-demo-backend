package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
)

// InitKafkaConsumer 初始化kafkaConsumer
func InitKafkaConsumer(cfg *config.Config, s *service.Service) {
	partitionConsumer, err := consumer.ConsumePartition(cfg.Kafka.Topic, 0, sarama.OffsetNewest) // 根据topic取到所有的分区
	if err != nil {
		panic("error get consumer")
	}
	//defer consumer.Close()
	//i := reflect.Copy(partitionConsumer)
	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err = consumer.ConsumePartition("logstash_test", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error get partition consumer", err)
	}
	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			go func() {
				// TODO 未处理错误,应该把错误放在logger
				redis.SetTotalCount()
				s.KfkChan <- msg.Value
				fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
			}()
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}
