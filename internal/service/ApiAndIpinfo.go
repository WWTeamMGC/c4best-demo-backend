package service

import (
	"encoding/json"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
)

// PhasePackage 解析kafka队列中拿出来的数据
func (s *Service) PhasePackage() {
	for {
		for bytes := range s.KfkChan {
			HttpInfo := &model.HttpInfo{}
			err := json.Unmarshal(bytes, HttpInfo)
			if err != nil {
				// TODO err写入日志
				return
			}
			//存入Api
			api := &model.Api{
				Router: HttpInfo.UrlPath,
				IpInfoList: []model.IpInfo{
					{
						IpAddr: HttpInfo.ClientIP,
						Url:    HttpInfo.UrlPath,
					},
				},
			}
			apiModel := model.NewApiModel(s.db, s.rds)
			//TODO 处理err写入日志
			apiModel.CreateApi(api)
			//存入IpInfo
			ipinfo := &model.IpInfo{
				IpAddr: HttpInfo.ClientIP,
				Url:    HttpInfo.UrlPath,
				ApiList: []model.Api{
					{
						Router: HttpInfo.UrlPath,
					},
				},
			}
			infoModel := model.NewIpInfoModel(s.db, s.rds)
			//TODO 处理错误
			infoModel.CreateIpInfo(ipinfo)
		}
	}
}
