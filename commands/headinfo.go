package commands

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

type HeadInfoCommand struct{}

func (p *HeadInfoCommand) Execute(args []string) {
	if len(os.Args) != 1 {
		fmt.Println("Usage: HeadInfo <host:port>")
		return
	}
	service := args[0]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(strings.Trim(string(result), "\r\n"))
}
