package main

import (
	"fmt"
	"log"
	"os"

	"github.com/masterZSH/hfile"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "hash",
				Aliases: []string{"h"},
				Usage:   "get hash of a file",
				Action: func(c *cli.Context) error {
					fileName := c.Args().First()
					if fileName == "" {
						fmt.Printf("invalid file name\n")
						return nil
					}
					hash, err := hfile.HashString(fileName)
					if err != nil {
						return err
					}
					fmt.Printf("hash of %s is: \n", fileName)
					fmt.Println(hash)
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
