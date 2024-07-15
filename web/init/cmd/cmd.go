package cmd

import (
	"web/config"
	"web/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	err := c.network.ServerStart(c.config.Server.Port)
	if err != nil {
		panic(err)
	}

	return c
}
