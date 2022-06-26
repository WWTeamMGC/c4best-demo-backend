package model

import "gorm.io/gorm"

type IpInfo struct {
	ID      uint64 `gorm:"primaryKey"`
	IpAddr  string `gorm:"index"`
	Count   uint64 `gorm:"type:int"`
	Address string `gorm:"size:256"`
	ApiList []Api  `gorm:"many2many:api"`
}
type Api struct {
	gorm.Model
	Router     string   `gorm:"size:256"`
	Count      uint64   `gorm:"type:int"`
	IpInfoList []IpInfo `gorm:"many2many:ipinfo"`
}
