package model

import "gorm.io/gorm"

type BadManInfo struct {
	gorm.Model
	IpList []string `gorm:"size:256"`
}
type BadWord struct {
	Word string `gorm:"size:256"`
}
