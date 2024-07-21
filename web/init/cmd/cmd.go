package cmd

import (
	"web/config"
	"web/network"
	"web/repository"
	"web/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}
	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	err := c.network.ServerStart(c.config.Server.Port)
	if err != nil {
		panic(err)
	}

	return c
}
