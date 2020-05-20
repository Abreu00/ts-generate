package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

var appName = "ts-generate"
var appAuthor = &cli.Author{Name: "Fábio de Abreu", Email: "fabiowabreu@gmail.com"}
var appUsage = "Generates a Typescript project template with TypeOrm"

func appAction(c *cli.Context) error {
	generateConfig()
	generateStructure()

	manager := c.String("manager")

	var cmd *exec.Cmd

	if manager == "yarn" {
		cmd = exec.Command("yarn", "init", "-y")
	} else {
		cmd = exec.Command("npm", "init", "-y")
	}
	stdout, err := cmd.Output()
	fmt.Printf(string(stdout))
	return err
}

func CreateApp() {
	app := &cli.App{
		Name:  appName,
		Usage: appUsage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "manager",
				Aliases: []string{"m"},
				Usage:   "Sets project package manager npm/yarn",
				Value:   "yarn",
			},
		},
		Action:  appAction,
		Authors: []*cli.Author{appAuthor},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
