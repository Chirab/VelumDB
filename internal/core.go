package core

import (
	"bufio"
	"fmt"
	"os"
)

type C struct {}

type CI interface {
	Core()
}

func (c *C) Core() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
