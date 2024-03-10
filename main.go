package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		// read keyboard input

		pwd, err := os.Getwd()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Printf("%v > ", pwd)
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {

	// Remove new-line char
	input = strings.TrimSuffix(input, "\n")

	// Prepare command to execute

	args := strings.Split(input, " ")

	// Check for built-in commands

	switch args[0] {

	case "cd":

		// 'cd' to home dir with an empty path not supported
		if len(args) < 2 {
			return errors.New("path required for command 'cd' ")
		}

		// change dir and return any error
		return os.Chdir(args[1])

	case "exit":
		fmt.Println("Exiting shell...")
		os.Exit(0)

	default:
		break
	}

	cmd := exec.Command(args[0], args[1:]...)

	// set the correct output devices
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd2 := exec.Command("cd", "..", "&& cd shell_in_go")
	// execute command
	return cmd2.Run()
}
