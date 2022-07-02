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
		CountapiRouter.GET("/Total", ctl.CountTotalHandler)
		CountapiRouter.GET("/Api/:api", ctl.SingleApiCountHandler)
		CountapiRouter.GET("/ip/:ip", ctl.SingleipCountHandler)
		CountapiRouter.GET("/Figure", ctl.CountFigureHandler)
	}

	//IP/Words的Web查询接口
	BadApiRouter := r.Group("/BadApi")
	//BadApiRouter.Use()
	{
		//查询BadIP和BadWords
		BadApiRouter.POST("/Ip", ctl.GetBadIPList)
		BadApiRouter.POST("/Words", ctl.GetBadWordsList)
		BadApiRouter.POST("/setIp", ctl.SetBadIP)
		BadApiRouter.POST("/setWords", ctl.SetBadWords)
		BadApiRouter.POST("/delIp", ctl.DeleteBadIP)
		BadApiRouter.POST("/delWords", ctl.DeleteBadWords)
	}

	//	//IP/Words过滤查询接口，非此web端接口
	//	BadApiIQRouter := r.Group("/BadApiIQ")
	//	//BadApiIQRouter.Use()
	//	{
	//		//返回0即表示IP/Words被封,返回1表示未被封
	//		BadApiIQRouter.POST("/Ip", ctl.BadIPIsExist)
	//		BadApiIQRouter.POST("/Words", ctl.BadWordsIsExist)
	//	}
}
