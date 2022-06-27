package kafka

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	producer     sarama.AsyncProducer
	producerOnce sync.Once
)

func NewProducer(cfg *config.Config) sarama.AsyncProducer {
	if !cfg.Kafka.Enable {
		return nil
	}
	producerOnce.Do(func() {
		config := sarama.NewConfig()
		config.Producer.Return.Successes = true
		config.Producer.Return.Errors = true
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Retry.Max = 3
		config.Producer.Flush.Frequency = 500
		config.Producer.Flush.Messages = 10
		config.Producer.Flush.MaxMessages = 10
		config.Producer.Partitioner = sarama.NewManualPartitioner
		config.Producer.Retry.Backoff = 500
		config.Producer.Compression = sarama.CompressionGZIP
		config.Version = sarama.V2_1_0_0
		var err error
		producer, err = sarama.NewAsyncProducer(cfg.Kafka.Brokers, sarama.NewConfig())
		if err != nil {
			panic(err)
		}
	})
	return producer
}
