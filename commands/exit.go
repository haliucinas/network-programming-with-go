package commands

import "os"

type ExitCommand struct{}
func (p *ExitCommand) Execute(args []string) {
	os.Exit(0)
}