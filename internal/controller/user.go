package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	State string
	Msg   string
}

func (ctl *Controller) Info(c *gin.Context) {
	rsp := &Response{
		State: "200",
		Msg:   "Pong",
	}
	c.JSON(http.StatusOK, rsp)
}
