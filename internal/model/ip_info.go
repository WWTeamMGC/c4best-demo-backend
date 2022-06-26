package model

import "gorm.io/gorm"

type IpInfo struct {
	ID      uint64 `gorm:"primaryKey"`
	IpAddr  string
	Count   int
	Address string
	ApiList []Api
}
type Api struct {
	gorm.Model
	Router     string
	Count      int
	IpInfoList []IpInfo
}

//
