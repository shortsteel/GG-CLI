package main

import (
	"GG-CLI/code"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "The GG CLI"

	app.Commands = []*cli.Command{
		code.Code(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
