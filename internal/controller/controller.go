package controller

import (
	"github.com/WWTeamMGC/c4best-demo-backend/internal/service"
	"sync"
)

var (
	controller *Controller
	once       sync.Once
)

type Controller struct {
	service *service.Service
}

func New(service *service.Service) *Controller {
	once.Do(func() {
		controller = &Controller{
			service: service,
		}
	})
	return controller
}
