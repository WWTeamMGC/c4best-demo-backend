package service

import (
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"time"
)

//func GetAllRouterAndCount() (res []model.Api, err error) {
//	if res, err = mysql.GetAllRouterAndCount(); err != nil {
//		return nil, err
//	} else {
//		return res, nil
//	}
//
//}
func GetAllRouterAndCount() (res map[string]string, err error) {
	if res, err = redis.GetAllRouterAndCount(); err != nil {
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

//拿到单个API被哪些IP调用或单个IP调用了哪些API   prefix即为需要查看的
func SingleCount(prefix string, req string) (res string, err error) {
	if res, err = redis.GetSingleCount(prefix, req); err != nil {
		fmt.Println(err)
		return
	} else {
		return
	}

}

//func GetCountBytime() (timeList, countList string, err error) {
//
//	if timeList, countList, err = redis.GetTimeAndCountList(); err != nil {
//		fmt.Println(err)
//		return
//	} else {
//
//		return
//	}
//}
func GetCountBytime() (res []string, err error) {

	if res, err = redis.GetTimeAndCountList(); err != nil {
		fmt.Println(err)
		return
	} else {

		return
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
