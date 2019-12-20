package util

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitInterruptSignal() chan os.Signal {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	return c
}
