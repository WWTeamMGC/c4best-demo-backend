package router

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, ctl *controller.Controller) {
	apiRouter := r.Group("/user")
	apiRouter.GET("/info", middleware.CORSMiddleware(), ctl.Info)
}
