package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"log"
)

func main() {
	nomalUsage()
}

func nomalUsage() {
	app := cli.NewApp()
	app.Name = "nomalUsage Name"
	app.Usage = "nothing"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
