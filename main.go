package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Abreu00/ts_tool/json_output"
	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func generateConfig() {
	var ts = json_output.NewTsconfig()
	ts.Write()
	var o = json_output.NewOrmConfig()
	o.Write()
}

func generateDirs() {
	paths := []string{"dist", "src", "src/database", "src/database/migrations", "src/entity"}
	perm := os.FileMode(0755)

	for _, path := range paths {
		os.Mkdir(path, perm)
	}
}

func main() {

	app := &cli.App{
		Name:  "greet",
		Usage: "fight loneliness",
		Action: func(c *cli.Context) error {
			cmd := exec.Command("yarn", "init", "-y")
			stdout, error := cmd.Output()
			fmt.Printf(string(stdout))
			generateConfig()
			generateDirs()

			return error
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
