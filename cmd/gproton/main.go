package main

import (
	"go-proton/cmd/gproton/commands"
	"go-proton/cmd/utils"
	"os"

	"gopkg.in/urfave/cli.v1"
)

var (
	app = utils.NewApp("the go-proton command line interface")
)

func init() {
	localConf, err := utils.Load("config.json")
	if err != nil {
		panic(err)
	}
	commands.SetConfig(localConf)
	app.Commands = []cli.Command{
		commands.AccountCommand,
		commands.SwapCommand,
	}
}

func main() {
	app.Run(os.Args)
}
