package main

import (
	"fmt"
	"log"
	"os"

	"lantool/cmd"

	"github.com/urfave/cli"
)

var (
	version string
	gitHash string
	logger  = log.Default()
)

func showBanner() {
	figlet := `

	░██           ██     ████     ██       ██████████   ███████     ███████   ██      
	░██          ████   ░██░██   ░██      ░░░░░██░░░   ██░░░░░██   ██░░░░░██ ░██      
	░██         ██░░██  ░██░░██  ░██          ░██     ██     ░░██ ██     ░░██░██      
	░██        ██  ░░██ ░██ ░░██ ░██ █████    ░██    ░██      ░██░██      ░██░██      
	░██       ██████████░██  ░░██░██░░░░░     ░██    ░██      ░██░██      ░██░██      
	░██      ░██░░░░░░██░██   ░░████          ░██    ░░██     ██ ░░██     ██ ░██      
	░████████░██     ░██░██    ░░███          ░██     ░░███████   ░░███████  ░████████
	░░░░░░░░ ░░      ░░ ░░      ░░░           ░░       ░░░░░░░     ░░░░░░░   ░░░░░░░░ 

Version: %s/%s
`
	fmt.Printf(figlet, version, gitHash)
}

func NewApp() *cli.App {
	showBanner()
	app := &cli.App{}
	app.Usage = "lan tool"
	app.HideVersion = true
	app.HideHelp = false
	app.Name = "lan-tool"

	app.Commands = []cli.Command{
		{
			Name:   "ddns",
			Usage:  "aliyun ddns",
			Flags:  cmd.DDNSFlags,
			Action: cmd.DDNSAction,
		},
	}

	app.Action = func(ctx *cli.Context) error {
		return cli.ShowAppHelp(ctx)
	}

	return app
}

func main() {
	app := NewApp()
	if err := app.Run(os.Args); err != nil {
		logger.Fatalf(err.Error())
	}
}
