package commands

import (
	"fmt"
	"net"
)

type ParseIPCommand struct{}
func (p *ParseIPCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: ParseIP <address>")
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