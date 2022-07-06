package controller

import (
	"encoding/json"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
)

// SetBadIP 设置BadIP
func (ctl *Controller) SetBadIP(c *gin.Context) {
	ip := c.PostForm("ip")
	badip := &model.BadIp{
		Ip:      ip,
		Address: RandAddress(),
		PcMp:    PcMp(),
	}
	err := ctl.service.SetBadIP(badip)
	if err != nil {
		//TODO 处理错误
		return
	}
	err = ctl.service.FlushBadIp()
	if err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// SetBadWords 设置BadWords
func (ctl *Controller) SetBadWords(c *gin.Context) {
	words := c.PostForm("word")
	badwords := &model.BadWords{
		Word: words,
	}
	err := ctl.service.SetBadWords(badwords)
	if err != nil {
		//TODO 处理错误
		ResponseError(c, CodeServerBusy)
	}
	ctl.service.FlushBadWords()
	ResponseSuccess(c, nil)
}

type BadIPListRsp struct {
	Ip      string `json:"ip"`
	PcMp    string `json:"pc_mp"`
	Address string `json:"address"`
}
type BadWordsListRsp struct {
	Badwords string `json:"word"`
}

func PcMp() string {
	PcMplist := make([]string, 2)
	PcMplist[0] = "PC"
	PcMplist[1] = "Phone"
	s := rand.Intn(2)
	return PcMplist[s]
}
func RandAddress() string {
	PcMplist := make([]string, 6)
	PcMplist[0] = "成都"
	PcMplist[1] = "武汉"
	PcMplist[2] = "沈阳"
	PcMplist[3] = "广州"
	PcMplist[4] = "上海"
	PcMplist[5] = "南宁"
	s := rand.Intn(6)
	return PcMplist[s]
}

// GetBadIPList 返回BadIPList
func (ctl *Controller) GetBadIPList(c *gin.Context) {
	var badiplist []BadIPListRsp
	for _, v := range ctl.service.BadIp {
		badip := BadIPListRsp{
			Ip:      v.Ip,
			PcMp:    v.PcMp,
			Address: v.Address,
		}
		badiplist = append(badiplist, badip)
	}

	c.JSON(http.StatusOK, gin.H{"badiplist": ctl.service.BadIp})
}

// GetBadWordsList 返回BadWordsList
func (ctl *Controller) GetBadWordsList(c *gin.Context) {
	var badwordslist []BadWordsListRsp
	for i := range ctl.service.BadWords {
		badword := BadWordsListRsp{
			Badwords: ctl.service.BadWords[i],
		}
		badwordslist = append(badwordslist, badword)
	}
	c.JSON(http.StatusOK, gin.H{"badwordslist": badwordslist})
}

// DeleteBadIP 删除BadIP
func (ctl *Controller) DeleteBadIP(c *gin.Context) {
	body := c.Request.Body
	all, err := io.ReadAll(body)
	badip := &BadIPListRsp{}
	json.Unmarshal(all, &badip)
	err = ctl.service.DeleteBadIP(badip.Ip)
	if err != nil {
		//TODO 处理错误
		c.JSON(http.StatusOK, gin.H{"Msg": err})
		return
	}
	ctl.service.FlushBadIp()
	ResponseSuccess(c, nil)
}

// DeleteBadWords 删除BadWords
func (ctl *Controller) DeleteBadWords(c *gin.Context) {
	body := c.Request.Body
	all, err := io.ReadAll(body)
	badip := &BadWordsListRsp{}
	json.Unmarshal(all, &badip)
	fmt.Println(badip)
	err = ctl.service.DeleteBadWords(badip.Badwords)
	if err != nil {
		//TODO 处理错误
		ResponseError(c, CodeServerBusy)
		return
	}
	ctl.service.FlushBadWords()
	ResponseSuccess(c, nil)
}
