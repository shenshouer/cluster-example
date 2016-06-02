package utils

import (
	"os"
	"os/signal"
	"syscall"
)

type closeFunc func()

// HandleSignal handle system signal
func HandleSignal(closeF closeFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for sig := range c {
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			closeF()
			return
		}
	}
}
