package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
