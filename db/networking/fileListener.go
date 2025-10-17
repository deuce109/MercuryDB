package networking

import (
	"log/slog"
	"net"

	"github.com/google/uuid"
)

type unixSocketListener struct {
	path   string
	socket *socket
	uuid   string
}

func RegisterUnixSocketListener(path string) {
	l := &unixSocketListener{
		path: path,
		uuid: uuid.New().String(),
	}

	listeners = append(listeners, l)
}

func (fl *unixSocketListener) Listen() {
	listener, err := net.Listen("unix", fl.path)

	if err != nil {
		slog.Error("Error starting file socket listener", "error", err)
		return
	}

	fl.socket = &socket{
		listener: listener,
	}
	err = fl.socket.Listen(func() { slog.Info("Listening on file socket...", "path", fl.path, "uuid", fl.uuid) })
	if err != nil {
		slog.Error("Error during Listen", "error", err)
	}
}

func (fl *unixSocketListener) Close() {
	slog.Info("Closing file socket listener...", "uuid", fl.uuid)

	if fl.socket != nil {
		err := fl.socket.Close()
		if err != nil {
			slog.Error("Error closing file socket", "uuid", fl.uuid, "error", err)
		}
	}
}
