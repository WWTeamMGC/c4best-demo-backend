package service

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
)

// FlushBadIp 初始化Bad名单InitBad
func (s *Service) FlushBadIp() {
	//var service.BadIp []&model.BadIp{}
	ipList, err := mysql.GetAllBadIp()
	if err != nil {
		return
	}
	service.BadIp = append(service.BadIp, ipList...)
}

// ReFlushBadWords 初始化Bad名单
func (s *Service) ReFlushBadWords() {
	//var service.BadIp []&model.BadIp{}
	wordsList, err := mysql.GetAllBadWords()
	if err != nil {
		//TODO 处理错误
		return
	}
	service.BadWords = append(service.BadWords, wordsList...)
}
func (s *Service) DeleteBadIP(str string) {

}
