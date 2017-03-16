package main

import (
	"log"
	"os"

	"fmt"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "command_line"
	app.Commands = append(app.Commands, getCommands()...)
	app.Run(os.Args)
}

//can be extracted to different file
func getCommands() []cli.Command {
	return []cli.Command{
		{
			Name:    "great-just-great-command",
			Aliases: []string{"gjg"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "not-so-great-but-still-good-flag, nsgbsg",
					Usage: "Changes behaviour of great-just-great-command",
				},
				cli.StringFlag{
					Name:  "limit, l",
					Value: "100",
					Usage: "Limits stuff",
				},
			},
			Description: "Performs great-just-great-command",
			Action:      greatMethod,
		},
		{
			Name:    "best-command",
			Aliases: []string{"b"},
			Subcommands: []cli.Command{
				{
					Name:    "bestest-subcommand",
					Aliases: []string{"bs"},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "finger, f",
							Usage: "Finger to show",
						},
					},
					Description: "Beautiful description",
					Usage:       "Just do it!",
					Action:      bestMethod,
				},
			},
		},
	}
}

func greatMethod(c *cli.Context) error {
	log.Println(fmt.Sprintf("Flags: %s and %s", c.String("gjg"), c.String("limit")))
	log.Println("Doing great stuff..")
	log.Println("Done.")
	return nil
}

func bestMethod(c *cli.Context) error {
	log.Println(fmt.Sprintf("Flag: %s ", c.String("f")))
	log.Println("Doing best stuff..")
	log.Println("Done.")
	return nil
}
