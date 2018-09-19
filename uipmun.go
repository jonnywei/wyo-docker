package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	cmd := exec.Command("/bin/bash")
	cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER| syscall.CLONE_NEWNET,}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
//        cmd.Stderr = os.Stderr

	if error := cmd.Run(); error != nil {
		log.Fatal(error)
	}
}
