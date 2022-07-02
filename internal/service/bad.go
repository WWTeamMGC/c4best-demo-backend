package service

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
)

// FlushBadIp 初始化Bad名单InitBad
func (s *Service) FlushBadIp() {
	service.BadIp = []string{}
	ipList, err := mysql.GetAllBadIp()
	if err != nil {
		return
	}
	for _, v := range ipList {
		service.BadIp = append(service.BadIp, v.Ip)
	}
}

// FlushBadWords 初始化BadWords名单
func (s *Service) FlushBadWords() {
	service.BadWords = []string{}
	wordsList, err := mysql.GetAllBadWords()
	if err != nil {
		return
	}
	for _, v := range wordsList {
		service.BadWords = append(service.BadWords, v.Word)
	}
}

// SetBadIP 设置BadIP
func (s *Service) SetBadIP(badip *model.BadIp) error {
	err := s.db.Model(&model.BadIp{}).Save(badip).Error
	return err
}

// SetBadWords 设置BadWords
func (s *Service) SetBadWords(badwords *model.BadWords) error {
	err := s.db.Model(&model.BadWords{}).Save(badwords).Error
	return err
}

// DeleteBadIP 删除BadIP
func (s *Service) DeleteBadIP(str string) error {
	err := s.db.Unscoped().Where("ip=?", str).Delete(&model.BadIp{}).Error
	return err
}

// DeleteBadWords 删除BadWords
func (s *Service) DeleteBadWords(str string) error {
	err := s.db.Unscoped().Where("word=?", str).Delete(&model.BadWords{}).Error
	return err
}
