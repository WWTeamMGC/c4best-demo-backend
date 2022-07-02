package controller

import (
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ctl *Controller) CountDetailHandler(c *gin.Context) {

	res, err := service.GetAllRouterAndCount()
	fmt.Println(res)
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)
	return
}
func (ctl *Controller) SingleApiCountHandler(c *gin.Context) {
	api := c.Param("api")

	res, err := service.SingleCount("ip", "/"+api)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)
}
func (ctl *Controller) SingleipCountHandler(c *gin.Context) {
	ip := c.Param("ip")

	res, err := service.SingleCount("url", ip)
	if err != nil {
		fmt.Println(err)
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)

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

	ResponseSuccess(c, res)
	return
}
