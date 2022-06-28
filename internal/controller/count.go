package controller

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func (ctl *Controller) CountDetailHandler(c *gin.Context) {

	res, err := service.GetAllRouterAndCount()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, res)
	return
}
func (ctl *Controller) CountTotalandler(c *gin.Context) {
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
func (ctl *Controller) CountFigureHandler(c *gin.Context) {
	timeListStr, countListStr, err := service.GetCountBytime()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	timeList := strings.SplitN(timeListStr, " ", 1440)
	countList := strings.SplitN(countListStr, " ", 1440)
	data := make([]TimeAndCount, len(timeList), 1440)
	for i := 0; i < len(timeList); i++ {
		data[i] = TimeAndCount{timeList[i], countList[i]}
	}
	ResponseSuccess(c, data)
	return
}
