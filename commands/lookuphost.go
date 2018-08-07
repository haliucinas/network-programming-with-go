package commands

import (
	"fmt"
	"net"
)

type LookupHostCommand struct{}
func (p *LookupHostCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: HostLookup <hostname>")
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