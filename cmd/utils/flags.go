package utils

import (
	"os"
	"path/filepath"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	version = "0.0.1"
)

func NewApp(usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = version
	app.Usage = usage
	return app
}
