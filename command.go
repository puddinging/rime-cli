package main

import (
	"fmt"
	"os/exec"
)

func exeCommand(command string) {
	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func exeCommandWithReturn(command string) {
	cmd := exec.Command("ls")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
