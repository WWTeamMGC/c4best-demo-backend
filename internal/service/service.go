package service

import (
	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	kfk "github.com/WWTeamMGC/c4best-demo-backend/internal/dao/kafka"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	cfg     *config.Config
	db      *gorm.DB
	rds     *redis.Client
	kafka   *sarama.Consumer
	KfkChan chan []byte
}

var (
	service *Service
	once    sync.Once
)

func New(cfg *config.Config, db *gorm.DB, rds *redis.Client, kafka *sarama.Consumer) *Service {
	once.Do(func() {
		service = &Service{
			cfg:     cfg,
			db:      db,
			rds:     rds,
			kafka:   kafka,
			KfkChan: make(chan []byte, 1000),
		}
	})
	go WatchRedis()
	//初始化KafkaConsumer
	go kfk.InitKafkaConsumer(cfg, service)
	//启动Http监听
	go service.PhasePackage()
	return service
}
