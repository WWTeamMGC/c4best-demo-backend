package model

import "gorm.io/gorm"

// BadManInfo 阻止IP
type BadManInfo struct {
	gorm.Model
	IpList []string `gorm:"many2many:iplist"`
}

// BadWord 敏感词
type BadWord struct {
	Word string `gorm:"size:256"`
}
