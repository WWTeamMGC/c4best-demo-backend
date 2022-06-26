package model

import "gorm.io/gorm"

type IpInfo struct {
	ID      uint64 `gorm:"primaryKey"`
	IpAddr  string //  唯一  index
	Count   uint64
	Address string
	ApiList []Api
}
type Api struct {
	gorm.Model
	Router     string
	Count      uint64
	IpInfoList []IpInfo
}
