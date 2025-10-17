package networking

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log/slog"
	"net"
	"strings"

	"github.com/deuce109/mercurydb/v2/db/commands"
)

type socket struct {
	listener       net.Listener
	commandHandler *commands.CommandHandler
}

func (s *socket) Listen(callback func()) error {
	callback()

	conn, err := s.listener.Accept()
	for {
		if err != nil {
			return err
		}
		go s.handle(conn)
		conn, err = s.listener.Accept()
	}
}

func (s *socket) readSize(conn net.Conn) (int64, error) {
	var size int64
	err := binary.Read(conn, binary.LittleEndian, &size)
	return size, err
}

func (s *socket) readFrom(conn net.Conn) (*bytes.Buffer, error) {

	size, err := s.readSize(conn)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	_, err = io.CopyN(&buf, conn, size)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func (s *socket) handle(conn net.Conn) {

	defer conn.Close()

	data, err := s.readFrom(conn)

	if err != nil {
		slog.Error("Error reading connection", "error", err)
		return
	}

	commandString := strings.Trim(data.String(), string(rune(0x00)))

	s.commandHandler, err = commands.DefaultCommandHandler()
	if err != nil {
		slog.Error("Error creating command handler", "error", err)
		return
	}

	dataBuffer, err := s.commandHandler.HandleMessage(commandString)
	if err != nil {
		fmt.Printf("Error handling command: %v\n", err)
	}

	_, err = io.Copy(conn, dataBuffer)
	if err != nil {
		fmt.Printf("Error writing connection response: %v\n", err)
	}

}

func (s *socket) Close() error {
	if s.listener != nil {
		return s.listener.Close()
	}
	return nil
}
