package main

import (
	config "go-proton/cmd"
	"os"

	"go-proton/cmd/utils"

	"gopkg.in/urfave/cli.v1"
)

var (
	app = utils.NewApp("the go-proton command line interface")
	cfg config.Config
)

func init() {
	localConf, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}
	cfg = localConf
	app.Commands = []cli.Command{
		accountCommand,
		swapCommand,
	}
}

func main() {
	app.Run(os.Args)
}
