package httpserver

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/config"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver/router"
	"github.com/gin-gonic/gin"
	"strings"
)

func Run(config *config.Config, controller *controller.Controller) {
	r := gin.Default()
	router.InitRouter(r, controller)
	host := strings.Join([]string{config.Http.Host, config.Http.Port}, ":")
	r.Run(host)
}
