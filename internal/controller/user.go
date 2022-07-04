package controller

import (
	"fmt"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/model"
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	jwt "github.com/WWTeamMGC/c4best-demo-backend/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	State string
	Msg   string
}

const CtxUserName = "ctxUserName"

func (ctl *Controller) Info(c *gin.Context) {
	rsp := &Response{
		State: "200",
		Msg:   "Pong",
	}
	c.JSON(http.StatusOK, rsp)
}
func (ctl *Controller) SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	p := new(model.User)
	p.Username = c.PostForm("username")
	p.Password = c.PostForm("password")
	if len(p.Username) < 8 {
		ResponseError(c, CodeUserNameShort)
		return
	}
	if len(p.Password) < 8 {
		ResponseError(c, CodeUserExsist)
	}
	// 2. 业务处理
	if err := service.SignUp(p); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeUserExsist)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func (ctl *Controller) SignInHandler(c *gin.Context) {
	p := new(model.User)
	p.Username = c.PostForm("username")
	p.Password = c.PostForm("password")
	if err := service.SignIn(p); err != nil {
		fmt.Println(err)
		ResponseError(c, CodeUserOrPasswordNotExsist)
		return
	} else {
		var token string
		token, err = jwt.GenToken(p.Username)
		if err != nil {
			fmt.Println(err)
		}
		ResponseSuccess(c, token)
		return
	}

}
