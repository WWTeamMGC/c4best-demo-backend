package redis

import (
	"fmt"
	"time"
)

func GetToTalCount() (res string, err error) {

	ctx := rdb.Context()

	rdb.Get(ctx, "totalCount")

	if res, err = rdb.Get(ctx, "totalCount").Result(); err != nil {
		return "", err
	} else {
		return res, nil
	}

}

func GetTimeList() (res []string, err error) {
	ctx := rdb.Context()
	if res, err = rdb.LRange(ctx, "countByTime", 0, 24*60).Result(); err != nil {
		if res, err = rdb.LRange(ctx, "timeList", 0, 24*60).Result(); err != nil {
			return nil, err
		} else {
			return res, nil
		}
	}
	//TODO something
	return nil, nil
}
func GetCountList() (res []string, err error) {
	ctx := rdb.Context()
	if res, err = rdb.LRange(ctx, "countList", 0, 24*60).Result(); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
func GetTimeAndCountList() (timeList, countList string, err error) {
	ctx := rdb.Context()
	pip := rdb.Pipeline()
	pip.LRange(ctx, "timeList", 0, 24*60)
	pip.LRange(ctx, "countList", 0, 24*60)

	cm, err := pip.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	//for _, v := range cm {
	//	//fmt.Println(v.String())
	//}

	return cm[1].String()[26:], cm[0].String()[25:], nil
}

func PullCountAndTime(res string) (err error) {
	ctx := rdb.Context()
	timestr := time.Now().String()
	req := fmt.Sprintf("%s %s", res, timestr)
	if err = rdb.LPush(ctx, "countByTime", req).Err(); err != nil {
		pip := rdb.Pipeline()
		timestr := time.Now()
		reqTime := fmt.Sprintf("%d:%d:%d", timestr.Hour(), timestr.Minute(), timestr.Second())
		fmt.Println(reqTime)
		pip.LPush(ctx, "countList", reqTime)
		pip.LPush(ctx, "timeList", res)
		if _, err = pip.Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}
