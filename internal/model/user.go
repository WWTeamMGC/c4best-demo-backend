package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:32;unique_index"`
	NickName string `gorm:"size:32;unique_index"`
	Password string `gorm:"size:64;unique_index"`
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

// GetUser 通过ID查找用户
func (u *UserModel) GetUser(id uint64) (*User, error) {
	var user User
	if err := u.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName 根据Name查找用户
func (u *UserModel) GetUserByName(username string) (*User, error) {
	var user User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建用户
func (u *UserModel) CreateUser(user *User) (id uint64, err error) {
	err = u.db.Model(&User{}).Save(user).Error

	return uint64(user.ID), err
}
