package main

import (
	"flag"
	"web/init/cmd"
)

var configPath = flag.String("config", "./config.toml", "config file path")

func main() {
	cmd.NewCmd(*configPath)
}
