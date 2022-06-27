package router

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/controller"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/httpserver/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine, ctl *controller.Controller) {
	{
		UserapiRouter := r.Group("/user")
		UserapiRouter.GET("/info", middleware.CORSMiddleware(), ctl.Info)
		UserapiRouter.POST("/SignUp", middleware.CORSMiddleware(), ctl.SignUpHandler)
		UserapiRouter.POST("/SignIn", middleware.CORSMiddleware(), ctl.SignInHandler)
	}

	{
		CountapiRouter := r.Group("/Count")
		CountapiRouter.GET("/detail", middleware.CORSMiddleware(), middleware.JWTAuthMiddleware(), ctl.CountDetailHandler)
		CountapiRouter.GET("/Total", middleware.CORSMiddleware(), middleware.JWTAuthMiddleware(), ctl.CountTotalandler)
		CountapiRouter.GET("/Figure", middleware.CORSMiddleware(), middleware.JWTAuthMiddleware(), ctl.CountFigureHandler)
	}
}
