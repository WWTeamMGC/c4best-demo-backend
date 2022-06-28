package router

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, ctl *controller.Controller) {

	r.Use(middleware.CORSMiddleware())

	UserapiRouter := r.Group("/user")
	{
		UserapiRouter.GET("/info", ctl.Info)
		UserapiRouter.POST("/SignUp", ctl.SignUpHandler)
		UserapiRouter.POST("/SignIn", ctl.SignInHandler)
	}

	CountapiRouter := r.Group("/Count")
	CountapiRouter.Use(middleware.JWTAuthMiddleware())
	{
		CountapiRouter.GET("/detail", ctl.CountDetailHandler)
		CountapiRouter.GET("/Total", ctl.CountTotalandler)
		CountapiRouter.GET("/Figure", ctl.CountFigureHandler)
	}
}
