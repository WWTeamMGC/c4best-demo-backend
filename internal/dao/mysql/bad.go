package mysql

import "github.com/WWTeamMGC/c4best-demo-backend/internal/model"

// GetAllBadIp 查询所有BadIp
func GetAllBadIp() ([]model.BadIp, error) {
	var str []model.BadIp
	if err := db.Raw("SELECT * FROM bad_ips").Scan(&str).Error; err != nil {
		return nil, err
	}
	return str, nil
}

// GetAllBadWords 查询所有BadWords
func GetAllBadWords() ([]model.BadWords, error) {
	var str []model.BadWords
	if err := db.Raw("SELECT * FROM bad_words").Scan(&str).Error; err != nil {
		return nil, err
	}
	return str, nil
}
