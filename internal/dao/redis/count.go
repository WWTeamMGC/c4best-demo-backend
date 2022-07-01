package redis

import (
	"encoding/json"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"strconv"
	"time"
)

const Prefix = "singleCount"

type TimeAndCount struct {
	Time  string `json:"time"`
	Count string `json:"value"`
}

func GetToTalCount() (res string, err error) {
	ctx := rdb.Context()

	rdb.Get(ctx, "totalCount")

	if res, err = rdb.Get(ctx, "totalCount").Result(); err != nil {
		return "", err
	} else {
		return res, nil
	}

}

func GetTimeAndCountList() (res []string, err error) {
	ctx := rdb.Context()

	res, err = rdb.LRange(ctx, "timeAndCount", 0, 24*60).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	return res, nil
}

func PullCountAndTime(res string) (err error) {
	ctx := rdb.Context()

	timestr := time.Now()
	reqTime := fmt.Sprintf("%d:%d:%d", timestr.Hour(), timestr.Minute(), timestr.Second())
	var timeAndCount = TimeAndCount{
		Time:  reqTime,
		Count: res,
	}
	tc, _ := json.Marshal(timeAndCount)
	//fmt.Println(string(tc))
	err = rdb.LPush(ctx, "timeAndCount", string(tc)).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func SetTotalCount() error {
	ctx := rdb.Context()
	_, err := rdb.Incr(ctx, "totalCount").Result()
	if err != nil {
		return err
	}
	return nil
}
func SetSingleCount(ipInfo *model.IpInfo) (err error) {
	ctx := rdb.Context()
	pip := rdb.Pipeline()
	key := Prefix + ipInfo.Url
	rdb.Incr(ctx, key)
	countstr, _ := rdb.HGet(ctx, ipInfo.IpAddr, ipInfo.Url).Result()
	count, _ := strconv.Atoi(countstr)
	pip.HSet(ctx, ipInfo.IpAddr, ipInfo.Url, count+1)
	countstr, _ = rdb.HGet(ctx, ipInfo.Url, ipInfo.IpAddr).Result()
	count, _ = strconv.Atoi(countstr)
	pip.HSet(ctx, ipInfo.Url, ipInfo.IpAddr, count+1)
	_, err = pip.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	//

	a, _ := rdb.HGetAll(ctx, ipInfo.IpAddr).Result()
	fmt.Println(a)
	return nil
}
