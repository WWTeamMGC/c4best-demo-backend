package model

import "gorm.io/gorm"

// BadIp 阻止IP
type BadIp struct {
	gorm.Model
	Ip    string `gorm:"size:256"`
	Count uint64 `gorm:"int"`
}

// BadWords 敏感词
type BadWords struct {
	Word  string `gorm:"size:256"`
	Count uint64 `gorm:"type:int"`
}
