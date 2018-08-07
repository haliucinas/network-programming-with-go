package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	cmds "bitbucket.org/Haliucinas/network-programming-with-go/commands"
)

var (
	reader *bufio.Reader
)

func init() {
	reader = bufio.NewReader(os.Stdin)

	Register("Exit", &cmds.ExitCommand{})
	Register("ParseIP", &cmds.ParseIPCommand{})
	Register("IPMask", &cmds.IPMaskCommand{})
	Register("LookupHost", &cmds.LookupHostCommand{})
	Register("LookupPort", &cmds.LookupPortCommand{})
	Register("HeadInfo", &cmds.HeadInfoCommand{})
}

func main() {
	for {
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		text = strings.Trim(text, "\n\r")
		splits := strings.Split(text, " ")
		name := splits[0]
		args := splits[1:]
		if command := commands[strings.ToLower(name)]; command == nil {
			fmt.Println("No such command found")
		} else {
			command.Execute(args)
		}
	}
}
