package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Register("exit", &ExitCommand{})
	Register("parse", &ParseIPCommand{})
	fmt.Println(commands)
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
		if command := commands[name]; command == nil {
			fmt.Println("No such command found")
		} else {
			command.Execute(args)
		}
	}
}