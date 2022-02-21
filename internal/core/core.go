package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type C struct {
	Command string
}

type CI interface {
	Core()
	TextInit()
}

/*
	Print Text when launching the core
*/
func (c *C) TextInit() {
	fmt.Println("VelumDB v.0.1")
	fmt.Println("Enter .help for usage hints.")
	fmt.Println("By Glio.")
	fmt.Println("\n")
}

func (c *C) Help() {
	fmt.Println("helpkfozkopfzeofze")
}

func (c *C) Core() {
	c.readCommand()
}

func (c *C) readCommand() {
	for {
		fmt.Print(">> ")
		scanner := bufio.NewReader(os.Stdin)
		input, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}
		// remove the delimeter from the string
		input = strings.TrimSuffix(input, "\n")
		c.Command = strings.ToLower(input)
		//check if inputt is null so the default case on the switch isnt reached out
		if input != "" {
			c.checkCommand()
		}
	}
}

func (c *C) checkCommand() {
	switch c.Command {
	case "exit":
		os.Exit(84)
	case "help":
		c.Help()
	case "create":
		fmt.Println("Creating Database")
	default:
		fmt.Println("Unrecognized command")
	}
}
