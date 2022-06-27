package controller

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
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
	res, err := service.GetCountBytime()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, res)
	return
}
