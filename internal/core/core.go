package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

type C struct {
	Command string
}

type CI interface {
	Core()
	TextInit()
}

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
	c.checkSize()
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

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)


func (c *C) checkSize() {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return
	}
	t := fs.Bfree * uint64(fs.Bsize)
	test := float64(t)/float64(GB)
	fmt.Println(test)
}
