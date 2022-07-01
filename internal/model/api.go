package model

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Api struct {
	gorm.Model
	Url        string   `gorm:"size:256"`
	Count      uint64   `gorm:"type:int"`
	IpInfoList []IpInfo `gorm:"many2many:ipinfo"`
}
type ApiModel struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewApiModel(db *gorm.DB, rdb *redis.Client) *ApiModel {
	return &ApiModel{
		db:  db,
		rdb: rdb,
	}
}

// GetApi 根据ID获取Api信息
func (a *ApiModel) GetApi(id uint64) (*Api, error) {
	var api Api
	if err := a.db.First(&api, id).Error; err != nil {
		return nil, err
	}
	return &api, nil
}

// GetApiByName 根据Api.Router查找Api
func (a *ApiModel) GetApiByName(url string) (*Api, error) {
	var api Api
	if err := a.db.Where("url = ?", url).First(&api).Error; err != nil {
		return nil, err
	}
	return &api, nil
}

// CreateApi 创建APiRouter
func (a *ApiModel) CreateApi(api *Api) (id uint64, err error) {
	var u = new(Api)
	a.db.Where("url=?", api.Url).First(u)
	if u.Url != "" {
		return
	}
	err = a.db.Model(&Api{}).Save(api).Error
	return uint64(api.ID), err
}
