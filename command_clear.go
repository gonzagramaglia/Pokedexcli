package main

import (
	"os"
	"os/exec"
)

func commandClear() error {
	clean := exec.Command("clear")
	clean.Stdout = os.Stdout
	clean.Run()
	return nil
}
