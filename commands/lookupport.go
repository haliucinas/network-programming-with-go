package commands

import (
	"fmt"
	"net"
)

type LookupPortCommand struct{}

func (p *LookupPortCommand) Execute(args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: PortLookup <network type> <service>")
		return
	}
	networkType := args[0]
	service := args[1]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Println("Service port:", port)
}
