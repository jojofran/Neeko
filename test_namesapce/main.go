//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	createUTS()
}

//UTS Test
func createUTS() {
	cmd := exec.Command("sh")

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf(os.Hostname())

	// ProcAttr holds the attributes that will be applied to a new process started by StartProcess.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWIPC | syscall.
			CLONE_NEWNET | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
	}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/sh command - %s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf(os.Hostname())
	}

}
