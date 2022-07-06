package model

import "gorm.io/gorm"

// BadIp 阻止IP
type BadIp struct {
	gorm.Model
	Ip      string `gorm:"size:256"`
	Address string `gorm:"size:256"`
	Count   uint64 `gorm:"int"`
	PcMp    string `gorm:"size:256"`
}

// BadWords 敏感词
type BadWords struct {
	gorm.Model
	Word  string `gorm:"size:256"`
	Count uint64 `gorm:"type:int"`
}
type BadIpRsp struct {
	Ip      string `json:"ip"`
	PcMp    string `json:"pc_mp"`
	Address string `json:"address"`
}
