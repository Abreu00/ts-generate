package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {

	app := &cli.App{
		Name:  "greet",
		Usage: "fight loneliness",
		Action: func(c *cli.Context) error {
			fmt.Println("Boom I say")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
