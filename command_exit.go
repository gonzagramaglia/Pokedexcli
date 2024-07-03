package main

import "os"

func commandExit() error {
	// Exit the program with a status code of 0 (success)
	os.Exit(0)
	return nil
}
