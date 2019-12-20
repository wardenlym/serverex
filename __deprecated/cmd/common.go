package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	ps "github.com/keybase/go-ps"
)

func ShowStatus() {
	mongod := findProc("mongod")
	fmt.Printf("%15s%7s %s\n", "NAME ", "PID", "EXE")
	if mongod == nil {
		fmt.Printf("%15s\n", "mongod ")
	} else {
		fmt.Printf("%15s%7d %s\n", "mongod ", mongod.Pid(), mongod.Executable())
	}

	exe := findProc("excalibur-api")
	if exe == nil {
		fmt.Printf("%15s\n", "excalibur-api ")
	} else {
		fmt.Printf("%15s%7d %s\n", "api ", exe.Pid(), exe.Executable())
	}

	demo := findProc("demo-server")
	if demo == nil {
		fmt.Printf("%15s\n", "demo-server")
	} else {
		fmt.Printf("%15s%7d %s\n", "demo-server ", demo.Pid(), demo.Executable())
	}
}

// FindProcessHasPrefix .
func FindProcessHasPrefix(prefix string) (*os.Process, error) {
	procs, err := ps.Processes()
	if err != nil {
		return nil, err
	}
	for _, proc := range procs {
		if strings.HasPrefix(proc.Executable(), prefix) {
			exe, err := os.FindProcess(proc.Pid())
			if err != nil {
				return nil, err
			}
			return exe, nil
		}
	}
	return nil, errors.New("process " + prefix + "* not found")
}

func findProc(prefix string) ps.Process {
	procs, err := ps.Processes()
	if err != nil {
		return nil
	}
	for _, proc := range procs {
		if strings.HasPrefix(proc.Executable(), prefix) {
			return proc
		}
	}
	return nil
}

func RunFreground(exe string, args string) {
	invoke := exec.Command(exe, args)
	invoke.Stdout = os.Stdout
	invoke.Stderr = os.Stdout
	err := invoke.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func Start(exe string, args string) {
	invoke := exec.Command(exe, args)
	invoke.Stdout = os.Stdout
	invoke.Stderr = os.Stdout
	err := invoke.Start()
	if err != nil {
		fmt.Println(err)
	}
}

func Stop(appname string) {
	exe, err := FindProcessHasPrefix(appname)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found, PID:%d\n", exe.Pid)
	err = exe.Signal(syscall.SIGTERM)
	if err != nil {
		fmt.Println("SIGTERM " + err.Error())
		if err.Error() == "not supported by windows" {
			fmt.Println("using Kill on windows")
			err = exe.Kill()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
