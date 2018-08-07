package main

import (
	"os"
	"fmt"
	"net"
)

var commands = map[string]Command{}

type Command interface {
	Execute([]string)
}

func Register(name string, command Command) {
	commands[name] = command
}

type ExitCommand struct{}
func (p *ExitCommand) Execute(args []string) {
	os.Exit(0)
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

type IPMaskCommand struct{}
func (p *IPMaskCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: mask <address>")
		return
	}
	dotAddr := args[0]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		return
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is", addr.String(),
		"\nMask length is", bits,
		"\nLeading ones count is", ones,
		"\nMask is (hex)", mask.String(),
		"\nNetwork is", network.String(),
	)
}

type HostLookupCommand struct{}
func (p *HostLookupCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: lookup <hostname>")
		return
	}
	name := args[0]
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
}