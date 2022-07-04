package controller

import (
	"encoding/json"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ctl *Controller) CountApiDetailHandler(c *gin.Context) {

	res, err := service.CountDetail("Api")

	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	var r []UrlAndCount
	for k, v := range res {
		var b UrlAndCount
		b.Url = k
		c, _ := strconv.Atoi(v)
		b.Count = c
		r = append(r, b)
	}
	ResponseSuccess(c, r)
	return
}
func (ctl *Controller) CountIpDetailHandler(c *gin.Context) {

	res, err := service.CountDetail("Ip")

	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	var r []UrlAndCount
	for k, v := range res {
		var b UrlAndCount
		b.Url = k
		c, _ := strconv.Atoi(v)
		b.Count = c
		r = append(r, b)
	}
	ResponseSuccess(c, r)
	return
}

func (ctl *Controller) SingleApiCountHandler(c *gin.Context) {
	api := c.Param("api")

	res, err := service.SingleCount("ip", "/"+api)
	var m = make(map[string]int)

	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	json.Unmarshal([]byte(res), &m)
	var r []IPAndCount
	for key, v := range m {
		var b IPAndCount
		b.IP = key
		b.Count = v
		r = append(r, b)
	}
	fmt.Printf("%#v", r)
	ResponseSuccess(c, r)
}

type UrlAndCount struct {
	Url   string `json:"url"`
	Count int    `json:"count"`
}
type IPAndCount struct {
	IP    string `json:"ip"`
	Count int    `json:"count"`
}

func (ctl *Controller) SingleipCountHandler(c *gin.Context) {
	ip := c.Param("ip")

	res, err := service.SingleCount("url", ip)

	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	var m = make(map[string]int)
	json.Unmarshal([]byte(res), &m)
	var r []UrlAndCount
	for key, v := range m {
		var b UrlAndCount
		b.Url = key
		b.Count = v
		r = append(r, b)
	}
	ResponseSuccess(c, r)

}

func (ctl *Controller) CountTotalHandler(c *gin.Context) {
	res, err := service.GetToTalCount()

	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	num, err := strconv.Atoi(res)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, num)
	return

}

//func (ctl *Controller) CountFigureHandler(c *gin.Context) {
//	timeListStr, countListStr, err := service.GetCountBytime()
//	if err != nil {
//		ResponseError(c, CodeServerBusy)
//		return
//	}
//	timeList := strings.SplitN(timeListStr, " ", 1440)
//	countList := strings.SplitN(countListStr, " ", 1440)
//	data := make([]TimeAndCount, len(timeList), 1440)
//	for i := 0; i < len(timeList); i++ {
//		data[i] = TimeAndCount{timeList[i], countList[i]}
//	}
//	ResponseSuccess(c, data)
//	return
//}

func (ctl *Controller) CountFigureHandler(c *gin.Context) {
	res, err := service.GetCountBytime()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	b := []redis.TimeAndCount{}
	for _, v := range res {
		a := redis.TimeAndCount{}
		json.Unmarshal([]byte(v), &a)
		b = append(b, a)
	}
	ResponseSuccess(c, b)
	return
}
