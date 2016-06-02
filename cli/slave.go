package cli

import (
	"github.com/codegangsta/cli"

	"cluster-example/cluster"
	"cluster-example/utils"
)

func slaveAction(c *cli.Context) {
	cluster.NewWorker(c.String("name"), c.String("IP"), c.StringSlice("endpoints"))
	utils.HandleSignal(func() {})
}
