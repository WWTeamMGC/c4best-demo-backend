package service

import (
	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	cfg   *config.Config
	db    *gorm.DB
	rds   *redis.Client
	kafka *sarama.Consumer
}

var (
	service *Service
	once    sync.Once
)

func New(cfg *config.Config, db *gorm.DB, rds *redis.Client, kafka *sarama.Consumer) *Service {
	once.Do(func() {
		service = &Service{
			cfg:   cfg,
			db:    db,
			rds:   rds,
			kafka: kafka,
		}
	})
	go WatchRedis()
	return service
}
