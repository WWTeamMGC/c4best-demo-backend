package model

import "gorm.io/gorm"

type BadManInfo struct {
	gorm.Model
	IpList []string
}
type BadWord struct {
	Word string
}
