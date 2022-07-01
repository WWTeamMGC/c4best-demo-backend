package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

//// BadIPIsExist 查询BadIP是否存在
//func (ctl *Controller) BadIPIsExist(c *gin.Context) {
//	var badIp string
//	err := c.ShouldBind(&badIp)
//	if err != nil {
//		//TODO 处理错误
//		return
//	}
//	if _, ok := ctl.service.BadIp[badIp]; !ok {
//		c.JSON(http.StatusOK, 1)
//		return
//	}
//	c.JSON(http.StatusOK, 0)
//}
//
//// BadWordsIsExist 查询BadWords是否存在
//func (ctl *Controller) BadWordsIsExist(c *gin.Context) {
//	var badWords string
//	err := c.ShouldBind(&badWords)
//	if err != nil {
//		//TODO 处理错误
//		return
//	}
//	if _, ok := ctl.service.BadIp[badWords]; !ok {
//		c.JSON(http.StatusOK, 1)
//		return
//	}
//	c.JSON(http.StatusOK, 0)
//}

// GetBadIPList 返回BadIPList
func (ctl Controller) GetBadIPList(c *gin.Context) {
	BadIPList, err := json.Marshal(ctl.service.BadIp)
	if err != nil {
		//TODO 处理错误
		return
	}
	c.JSON(http.StatusOK, BadIPList)
}

// GetBadWordsList 返回BadWordsList
func (ctl Controller) GetBadWordsList(c *gin.Context) {
	BadWordsList, err := json.Marshal(ctl.service.BadWords)
	if err != nil {
		//TODO 处理错误
		return
	}
	c.JSON(http.StatusOK, BadWordsList)
}

// GetBadIPList 返回BadIPList
func (ctl Controller) DeleteBadIP(c *gin.Context) {
	var BadIP string
	err := c.ShouldBind(&BadIP)
	if err != nil {
		//TODO 处理错误
		return
	}

}
