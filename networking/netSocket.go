package networking

import (
	"fmt"
	"net"
	"strings"

	"github.com/deuce109/mercurydb/v2/parsers"
)

type NetworkListener struct {
	Port     int
	Host     string
	listener net.Listener
}

func NewNetworkListener(host string, port int) *NetworkListener {
	return &NetworkListener{
		Host: host,
		Port: port,
	}
}

func (nl *NetworkListener) Listen(callback func()) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", nl.Host, nl.Port))

	if err != nil {
		return err
	}

	nl.listener = listener

	callback()

	conn, err := nl.listener.Accept()
	for {
		if err != nil {
			return err
		}
		go nl.handle(conn)
		conn, err = nl.listener.Accept()
	}
}

func (nl *NetworkListener) handle(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 1024)

	_, err := conn.Read(buf)

	if err != nil {
		fmt.Printf("Error reading from connection: %v\n", err)
	}

	commandString := strings.Trim(string(buf), string(rune(0x00)))

	command, args, err := parsers.ParseCommand(commandString)

	if err != nil {
		fmt.Printf("Error parsing command: %v\n", err)
	}

	commandHandler := DefaultCommandHandler()

	value, err := commandHandler.HandleCommand(command, args...)
	if err != nil {
		fmt.Printf("Error handling command: %v\n", err)
	}

	_, err = conn.Write([]byte(fmt.Sprintf("%v", value)))
	if err != nil {
		fmt.Printf("Error writing connection response: %v\n", err)
	}

}

func (nl *NetworkListener) Close() error {
	if nl.listener != nil {
		err := nl.listener.Close()
		return err
	}
	return nil
}
