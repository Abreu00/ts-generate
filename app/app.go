package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func appAction(c *cli.Context) error {
	generateConfig()
	generateDirs()

	manager := c.String("manager")

	if manager == "yarn" {
		cmd := exec.Command("yarn", "init", "-y")
		stdout, err := cmd.Output()
		fmt.Printf(string(stdout))
		return err
	} else {
		cmd := exec.Command("npm", "init", "-y")
		stdout, err := cmd.Output()
		fmt.Printf(string(stdout))
		return err
	}
}

func CreateApp() {
	app := &cli.App{
		Name:  "ts-generate",
		Usage: "Create a typescript project with TypeOrm",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "manager",
				Aliases: []string{"m"},
				Usage:   "Sets project package manager npm/yarn",
				Value:   "yarn",
			},
		},
		Action: appAction,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
