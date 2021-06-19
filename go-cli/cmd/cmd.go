package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "@@( .app_name )",
		Action: func(c *cli.Context) error {
			return nil
		},

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Usage: "config path",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "sample",
				Description: "sample",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
