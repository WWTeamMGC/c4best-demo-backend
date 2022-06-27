package model

import "gorm.io/gorm"

type BadManInfo struct {
	gorm.Model
	IpList []string `gorm:"many2many:iplist"`
}
type BadWord struct {
	Word string `gorm:"size:256"`
}
