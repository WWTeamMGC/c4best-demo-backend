package service

import (
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"time"
)

func GetAllRouterAndCount() (res []model.Api, err error) {
	if res, err = mysql.GetAllRouterAndCount(); err != nil {
		return nil, err
	} else {
		return res, nil
	}

}
func GetToTalCount() (res string, err error) {
	if res, err = redis.GetToTalCount(); err != nil {
		return "", err
	} else {
		return res, nil
	}
}
func GetCountBytime() (res []string, err error) {

	if res, err = redis.GetCountByTime(); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func WatchRedis() {

	for {
		res, err := redis.GetToTalCount()
		if err != nil {
			fmt.Println(err)
		}
		err = redis.PullCountAndTime(res)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(60 * time.Second)
	}
}
