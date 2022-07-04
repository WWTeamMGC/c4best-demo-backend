package redis

import (
	"encoding/json"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
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
	fmt.Printf("%#v", tc)
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
	//key := Prefix + ipInfo.Url
	var urlAndCount = map[string]int{}
	ans, _ := rdb.HGet(ctx, "urlAndCount", ipInfo.IpAddr).Result()
	err = json.Unmarshal([]byte(ans), &urlAndCount)
	pip.HIncrBy(ctx, "singleApiCount", ipInfo.Url, 1)
	pip.HIncrBy(ctx, "singleIpCount", ipInfo.IpAddr, 1)
	if count, ok := urlAndCount[ipInfo.Url]; !ok {
		urlAndCount[ipInfo.Url] = 1
	} else {
		urlAndCount[ipInfo.Url] = count + 1
	}
	field, _ := json.Marshal(urlAndCount)
	pip.HSet(ctx, "urlAndCount", ipInfo.IpAddr, field)

	var ipAndCount = map[string]int{}
	ans2, _ := rdb.HGet(ctx, "ipAndCount", ipInfo.Url).Result()
	err = json.Unmarshal([]byte(ans2), &ipAndCount)

	if count, ok := ipAndCount[ipInfo.IpAddr]; !ok {
		ipAndCount[ipInfo.IpAddr] = 1
	} else {
		ipAndCount[ipInfo.IpAddr] = count + 1
	}

	field2, _ := json.Marshal(ipAndCount)
	fmt.Println(string(field2))
	pip.HSet(ctx, "ipAndCount", ipInfo.Url, field2)

	//pip.HIncrBy(ctx, "singleCount", ipInfo.Url, 1)
	////rdb.Incr(ctx, key)
	////countstr, _ := rdb.HGet(ctx, ipInfo.IpAddr, ipInfo.Url).Result()
	////count, _ := strconv.Atoi(countstr)
	//pip.HIncrBy(ctx, ipInfo.IpAddr, ipInfo.Url, 1)
	////countstr, _ = rdb.HGet(ctx, ipInfo.Url, ipInfo.IpAddr).Result()
	////count, _ = strconv.Atoi(countstr)
	//pip.HIncrBy(ctx, ipInfo.Url, ipInfo.IpAddr, 1)

	_, err = pip.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	//

	//a, _ := rdb.HGetAll(ctx, ipInfo.IpAddr).Result()
	//fmt.Println(a)
	return nil
}

func CountDetail(prefix string) (res map[string]string, err error) {
	ctx := rdb.Context()
	s := fmt.Sprintf("single%sCount", prefix)
	res, err = rdb.HGetAll(ctx, s).Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return res, err
}

//拿到单个API被哪些IP调用或单个IP调用了哪些API   prefix即为需要查看的
func GetSingleCount(prefix, req string) (res string, err error) {
	ctx := rdb.Context()
	fmt.Println(prefix+"AndCount", req)
	if res, err = rdb.HGet(ctx, prefix+"AndCount", req).Result(); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(res)
		return
	}

}
