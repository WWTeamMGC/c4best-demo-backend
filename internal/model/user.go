package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique_index"`
}
type UserModel struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserModel(db *gorm.DB, rdb *redis.Client) *UserModel {
	return &UserModel{
		db:  db,
		rdb: rdb,
	}
}
