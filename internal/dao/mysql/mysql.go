package mysql

import (
	"errors"
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	once sync.Once
	db   *gorm.DB
)

func New(cfg *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.Database)
		var err error
		mysqlconfig := mysql.Config{
			DSN: dsn,
		}
		for i := 0; i < 3; i++ {
			db, err = gorm.Open(mysql.New(mysqlconfig))
			if err == nil {
				break
			}
			time.Sleep(time.Second * 3)
		}
		if err != nil {
			panic(errors.New("Cannot  connect to mysql"))
		}
		if err = db.AutoMigrate(&model.User{}); err != nil {
			panic(err)
		}
	})
	return db
}
