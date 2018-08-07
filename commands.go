package main

import (
	"os"
	"net"
	"fmt"
)

var commands = map[string]Command{}

type Command interface {
	Execute([]string)
}

func Register(name string, command Command) {
	commands[name] = command
}

type ParseIPCommand struct{}
func (p *ParseIPCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: parse <address>")
		return
	}
	name := args[0]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is", addr.String())
	}
}

type ExitCommand struct{}
func (p *ExitCommand) Execute(args []string) {
	os.Exit(0)
}
