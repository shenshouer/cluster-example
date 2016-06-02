package cli

import (
	"cluster-example/cluster"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

func masterAction(c *cli.Context) {
	endpoints := c.StringSlice("endpoints")
	log.Infoln(endpoints)
	master := cluster.NewMaster(endpoints)

	for {
		time.Sleep(time.Second)
		members := master.Members()
		for k, v := range members {
			log.Info("key:", k, "value:", v)
		}
	}
}
