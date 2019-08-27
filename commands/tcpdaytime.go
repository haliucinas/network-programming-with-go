package commands

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

type TCPDaytimeServerCommand struct{}

func (p *TCPDaytimeServerCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: TCPDaytimeServer <host:port>")
		return
	}
	service := args[0]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println(err)
		return
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	stopch := RegisterRoutine("TCPDaytimeServer")
	go func() {
		for {
			select {
			case <-stopch:
				return
			default:
				listener.SetDeadline(time.Now().Add(1 * time.Second))
				conn, err := listener.Accept()
				if err != nil {
					continue
				}

				daytime := time.Now().String()
				conn.Write([]byte(daytime))
				conn.Close()
			}
		}
	}()
}

type TCPDaytimeClientCommand struct{}

func (p *TCPDaytimeClientCommand) Execute(args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: TCPDaytimeClient <host:port>")
		return
	}
	service := args[0]
	conn, err := net.Dial("tcp", service)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))
}
