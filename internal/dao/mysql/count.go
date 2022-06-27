package mysql

import "github.com/WWTeamMGC/c4best-demo-backend/internal/model"

func GetAllRouterAndCount() (res []model.Api, err error) {

	err = db.Order("count desc").Find(&res).Error
	if err != nil {
		return nil, ErrorBusy
	} else {
		return res, nil
	}
}
