package network

import (
	"web/service"

	"github.com/gin-gonic/gin"
)

type Network struct {
	engine  *gin.Engine
	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	n := &Network{
		engine:  gin.New(),
		service: service,
	}
	newUserRouter(n, service.UserService)
	return n
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}
