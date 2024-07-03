package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The main function that starts the REPL (Read-Eval-Print Loop)
func startRepl() {

	// Create a scanner for standard input (stdin)
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		// Read user input
		reader.Scan()

		// Clean and process user input
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		// Get the command name from the entered words
		commandName := words[0]

		// Retrieve the corresponding command from the commands map
		command, exists := getCommands()[commandName]
		if exists {
			// Execute the command's callback function
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue // Continue to the next iteration of the REPL
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}

}

// Function to clean user input and split it into words
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	// Split input into words (tokens)
	words := strings.Fields(output)
	return words
}

// Structure that defines a CLI command
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Function that returns a map of available commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"clear": {
			name:        "clear",
			description: "Clean the Pokedex screen",
			callback:    commandClear,
		},
	}
}
