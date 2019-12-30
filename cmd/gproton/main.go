package main

import (
	"os"

	"go-proton/cmd/utils"

	"gopkg.in/urfave/cli.v1"
)

var (
	app = utils.NewApp("the go-proton command line interface")
)

func init() {
	app.Commands = []cli.Command{
		accountCommand,
	}
}

func main() {
	app.Run(os.Args)
}
