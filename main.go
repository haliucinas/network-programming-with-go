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

	cmds.Register("Exit", &cmds.ExitCommand{})
	cmds.Register("lr", &cmds.ListRoutineCommand{})
	cmds.Register("sr", &cmds.StopRoutineCommand{})
	cmds.Register("ParseIP", &cmds.ParseIPCommand{})
	cmds.Register("IPMask", &cmds.IPMaskCommand{})
	cmds.Register("LookupHost", &cmds.LookupHostCommand{})
	cmds.Register("LookupPort", &cmds.LookupPortCommand{})
	cmds.Register("HeadInfo", &cmds.HeadInfoCommand{})
	cmds.Register("TCPDaytimeServer", &cmds.TCPDaytimeServerCommand{})
	cmds.Register("TCPDaytimeClient", &cmds.TCPDaytimeClientCommand{})
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
		if command := cmds.GetCommand(name); command == nil {
			fmt.Println("No such command found")
		} else {
			command.Execute(args)
		}
	}
}
