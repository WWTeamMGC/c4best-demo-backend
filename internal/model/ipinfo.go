package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type IpInfo struct {
	ID      uint64 `gorm:"primaryKey"`
	IpAddr  string `gorm:"index"`
	Count   uint64 `gorm:"type:int"`
	Address string `gorm:"size:256"`
	ApiList []Api  `gorm:"many2many:api"`
}

type IpInfoModel struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewIpInfoModel(db *gorm.DB, rdb *redis.Client) *IpInfoModel {
	return &IpInfoModel{
		db:  db,
		rdb: rdb,
	}
}
