package networking

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/google/uuid"
)

type networkListener struct {
	port   uint16
	host   string
	socket *socket
	uuid   string
}

func RegisterNetworkListener(host string, port uint16) {
	l := &networkListener{
		host: host,
		port: port,
		uuid: uuid.New().String(),
	}

	listeners = append(listeners, l)
}

func (nl *networkListener) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", nl.host, nl.port))

	if err != nil {
		slog.Error("Error starting TCP listener", "error", err)
		return
	}

	nl.socket = &socket{
		listener: listener,
	}
	err = nl.socket.Listen(func() { slog.Info("Listening...", "address", nl.host, "port", nl.port, "uuid", nl.uuid) })
	if err != nil {
		slog.Error("Error during Listen", "error", err)
	}
}

func (nl *networkListener) Close() {
	slog.Info("Closing network listener...", "uuid", nl.uuid)
	if nl.socket != nil {
		err := nl.socket.Close()
		if err != nil {
			slog.Error("Error closing network socket", "uuid", nl.uuid, "error", err)
		}
	}
}
