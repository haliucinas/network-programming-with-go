package main

var commands = map[string]Command{}

type Command interface {
	Execute([]string)
}

func Register(name string, command Command) {
	commands[name] = command
}