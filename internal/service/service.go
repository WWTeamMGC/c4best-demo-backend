package service

import (
	"github.com/Shopify/sarama"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"sync"
)

type Service struct {
	cfg      *config.Config
	db       *gorm.DB
	rds      *redis.Client
	kafka    sarama.Consumer
	KfkChan  chan []byte
	BadIp    map[string]uint64
	BadWords map[string]uint64
}

var (
	service *Service
	once    sync.Once
)

func New(cfg *config.Config, db *gorm.DB, rds *redis.Client, kafka sarama.Consumer,

//badIp, badWords map[string]uint64
) *Service {
	once.Do(func() {
		service = &Service{
			cfg:     cfg,
			db:      db,
			rds:     rds,
			kafka:   kafka,
			KfkChan: make(chan []byte, 1000),
			//	BadIp:    badIp,
			//BadWords: badWords,
		}
	})
	go WatchRedis()
	//初始化KafkaConsumer
	go service.InitKafkaConsumer(cfg, kafka)
	//启动Http监听
	go service.PhasePackage()
	return service
}
