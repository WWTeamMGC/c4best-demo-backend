package httpserver

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver/router"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"strings"
)

func Run(config *config.Config, controller *controller.Controller) {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	if config.Debug {
		pprof.Register(r)
	}
	router.InitRouter(r, controller)
	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host)
}
