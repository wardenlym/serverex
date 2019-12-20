package util

import (
	"fmt"
	"os"
)

var hostname string

func init() {
	hostname, _ = os.Hostname()
}

func Pid() string {
	return hostname + "." + fmt.Sprint(os.Getpid())

}
