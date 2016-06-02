package cli

import "github.com/codegangsta/cli"

var (
	endpoints = cli.StringSliceFlag{
		Name:  "endpoints",
		Value: &cli.StringSlice{"http://127.0.0.1:2379"},
		Usage: "etcd endpoints",
	}

	workerName = cli.StringFlag{
		Name:  "name",
		Value: "test",
		Usage: "The name for worker",
	}

	exampleIP = cli.StringFlag{
		Name:  "IP",
		Value: "1.1.1.1",
		Usage: "The IP for worker",
	}
)
