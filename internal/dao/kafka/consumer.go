package kafka

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	consumer     sarama.Consumer
	consumerOnce sync.Once
)

func NewConsumer(cfg *config.Config) sarama.Consumer {
	if !cfg.Kafka.Enable {
		return nil
	}
	consumerOnce.Do(func() {
		var err error
		config := sarama.NewConfig()
		config.Consumer.Return.Errors = true
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
		consumer, err = sarama.NewConsumer(cfg.Kafka.Brokers, config)
		if err != nil {
			panic(err)
		}
	})
	return consumer
}
