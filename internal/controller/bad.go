package controller

import (
	"encoding/json"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// SetBadIP 设置BadIP
func (ctl *Controller) SetBadIP(c *gin.Context) {
	ip := c.PostForm("badip")
	badip := &model.BadIp{
		Ip:    ip,
		Count: 0,
	}
	err := ctl.service.SetBadIP(badip)
	if err != nil {
		//TODO 处理错误
		return
	}
	ctl.service.FlushBadIp()
}

// SetBadWords 设置BadWords
func (ctl *Controller) SetBadWords(c *gin.Context) {
	words := c.PostForm("badwords")
	badwords := &model.BadWords{
		Word:  words,
		Count: 0,
	}
	err := ctl.service.SetBadWords(badwords)
	if err != nil {
		//TODO 处理错误
		return
	}
	ctl.service.FlushBadWords()
}

type BadIPListRsp struct {
	Ip      string `json:"ip"`
	PcMp    string `json:"pc_mp"`
	Address string `json:" address"`
}
type BadWordsListRsp struct {
	Badwords string `json:"badwords"`
	Delete   string `json:"delete"`
}

// GetBadIPList 返回BadIPList
func (ctl *Controller) GetBadIPList(c *gin.Context) {
	var badiplist []BadIPListRsp
	for i := range ctl.service.BadIp {
		badip := BadIPListRsp{
			Ip:      ctl.service.BadIp[i],
			PcMp:    "aaa",
			Address: "lll",
		}
		badiplist = append(badiplist, badip)
	}

	c.JSON(http.StatusOK, gin.H{"badiplist": badiplist})
}

// GetBadWordsList 返回BadWordsList
func (ctl *Controller) GetBadWordsList(c *gin.Context) {
	var badwordslist []BadWordsListRsp
	for i := range ctl.service.BadWords {
		badword := BadWordsListRsp{
			Badwords: ctl.service.BadWords[i],
			Delete:   "",
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
	c.JSON(http.StatusOK, gin.H{"Msg": "删除成功"})
}

// DeleteBadWords 删除BadWords
func (ctl *Controller) DeleteBadWords(c *gin.Context) {
	words := c.PostForm("badwords")
	err := ctl.service.DeleteBadWords(words)
	if err != nil {
		//TODO 处理错误
		c.JSON(http.StatusOK, gin.H{"Msg": err})
		return
	}
	ctl.service.FlushBadWords()
	c.JSON(http.StatusOK, gin.H{"Msg": "删除成功"})
}
