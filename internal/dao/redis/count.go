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
func GetCountByTime() (res []string, err error) {
	ctx := rdb.Context()
	if res, err = rdb.LRange(ctx, "countByTime", 0, 24*60).Result(); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
func PullCountAndTime(res string) (err error) {
	ctx := rdb.Context()
	timestr := time.Now().String()
	req := fmt.Sprintf("%s %s", res, timestr)
	if err = rdb.LPush(ctx, "countByTime", req).Err(); err != nil {
		return err
	}
	return nil
}
