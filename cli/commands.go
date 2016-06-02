package cli

import "github.com/codegangsta/cli"

var (
	commands = []cli.Command{
		{
			Name:      "master",
			ShortName: "M",
			Usage:     "Start master",
			Action:    masterAction,
			Flags:     []cli.Flag{endpoints},
		},
		{
			Name:      "slave",
			ShortName: "S",
			Usage:     "Start slave",
			Action:    slaveAction,
			Flags:     []cli.Flag{endpoints, workerName, exampleIP},
		},
	}
)
