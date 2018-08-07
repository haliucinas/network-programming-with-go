package main

import "strings"

var commands = map[string]Command{}

type Command interface {
	Execute([]string)
}

func Register(name string, command Command) {
	commands[strings.ToLower(name)] = command
}