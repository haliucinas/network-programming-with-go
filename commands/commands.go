package commands

import (
	"strconv"
	"strings"
)

var commands = map[string]Command{}
var routines = map[string]chan struct{}{}

type Command interface {
	Execute([]string)
}

func Register(name string, command Command) {
	commands[strings.ToLower(name)] = command
}

func RegisterRoutine(name string) chan struct{} {
	name += "-" + strconv.Itoa(len(routines))
	routines[name] = make(chan struct{})
	return routines[name]
}

func GetCommand(name string) Command {
	return commands[strings.ToLower(name)]
}
