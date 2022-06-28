package main

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/kafka"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/mysql"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/dao/redis"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"go.uber.org/fx"
)

func main() {

	//实现fx模块注入，简化、复用模块应用
	app := fx.New(
		fx.Provide(
			config.Phase,
			mysql.New,
			redis.New,
			service.New,
			controller.New,
			kafka.NewConsumer,
		),
		fx.Invoke(
			httpserver.Run,
		),
	)

	app.Run()

	//cfg, err := config.Phase()
	//if err != nil {
	//
	//}
	//db := mysql.New(cfg)
	//rdb := redis.New(cfg)
	//s := service.New(cfg, db, rdb)
	//ctl := controller.New(s)
	//
	//httpserver.Run(cfg, ctl)

}
