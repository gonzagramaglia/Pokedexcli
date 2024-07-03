package main

import (
	"os"
	"os/exec"
)

// Function to clear the terminal screen
func commandClear() error {
	// Create a new command to execute the "clear" command (clears terminal screen)
	clean := exec.Command("clear")

	// Set the command's standard output to the current process's standard output
	// so the output of the clear command is displayed in the terminal.
	clean.Stdout = os.Stdout

	// Run the command to clear the terminal screen
	clean.Run()

	// Return nil error indicating successful execution
	return nil
}
