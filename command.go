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

func exeCommandll(command string) {
	cmd := exec.Command("ll")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func exeCommandEcho(command string) {
	cmd := exec.Command("echo 'aaa'")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func exeCommandpr(command string) {
	cmd := exec.Command("echo 'print'")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func exeCommandp1r(command string) {
	cmd := exec.Command("echo 'print'")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
