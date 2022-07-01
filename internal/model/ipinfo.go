package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type IpInfo struct {
	gorm.Model
	IpAddr  string `gorm:"index"`
	Count   uint64 `gorm:"type:int"`
	Url     string `gorm:"size:256"`
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

// CreateIpInfo 创建APiRouter
func (a *IpInfoModel) CreateIpInfo(info *IpInfo) (id uint64, err error) {
	u := new(IpInfo)
	a.db.Where("ip_addr=?", info.IpAddr).First(u)
	if u.IpAddr != "" {
		return
	}
	err = a.db.Model(&IpInfo{}).Save(info).Error
	return uint64(info.ID), err
}
