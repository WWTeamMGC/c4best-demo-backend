package main

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
)

func main() {
	cfg, err := config.Phase()
	if err != nil {

	}
	db := mysql.New(cfg)
	rdb := redis.New(cfg)
	service := service.New(cfg, db, rdb)
	ctl := controller.New(service)
	httpserver.Run(cfg, ctl)
}
