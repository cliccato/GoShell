package main

import (
	"GoShell/commands"
	"GoShell/superpath"
	"GoShell/errors"
	"GoShell/colors"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printPS1()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		processCommand(input)
	}
}

func printPS1() {
	ps1 := colors.Green + superpath.GetCurrentDirectory() + "$ " + colors.Reset
	fmt.Print(ps1)
}

func processCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]
	params := parts[1:]
	options := make([]string, 0)

	for i := 0; i < len(params); i++ {
		if params[i][0] == '-' {
			options = append(options, params[i][1:])
			params = append(params[:i], params[i+1:]...)
			i--
		}
	}

	executeCommand(cmd, params, options)
}

func executeCommand(cmd string, params []string, options []string) {
	switch cmd {
	case "exit":
		fmt.Print("exit")
		os.Exit(0)
	case "cd":
		err := superpath.ChangeDirectory(params[0])
		if err != nil {
			errors.PrintError(cmd, err)
		}
	case "echo":
		commands.Echo(params)
	case "pwd":
		commands.Pwd()
	case "clear":
		commands.Clear()
	case "ls":
		commands.Ls(params, options)
	case "touch":
		commands.Touch(params)
	case "mkdir":
		commands.Mkdir(params)
	case "rm":
		commands.Rm(params, options)
	default:
		errors.PrintError(cmd, fmt.Errorf("Comand not found"))
	}
}
