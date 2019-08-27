package commands

import (
	"fmt"
)

type ListRoutineCommand struct{}

func (p *ListRoutineCommand) Execute(args []string) {
	switch len(routines) {
	case 0:
		fmt.Println("No running routines")
	case 1:
		fmt.Println("Running routine:")
	default:
		fmt.Println("Running routines:")
	}
	for routine := range routines {
		fmt.Println(routine)
	}
}

type StopRoutineCommand struct{}

func (p *StopRoutineCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: sr <routine name>")
		return
	}
	name := args[0]
	if routine := routines[name]; routine == nil {
		fmt.Printf("%s is not running.\n", name)
	} else {
		close(routine)
		delete(routines, name)
	}
}
