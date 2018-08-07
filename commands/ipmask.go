package commands

import (
	"fmt"
	"net"
)

type IPMaskCommand struct{}
func (p *IPMaskCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: IPMask <address>")
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