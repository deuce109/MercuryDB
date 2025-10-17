package networking

import (
	"encoding/binary"
	"log/slog"
	"net"
	"testing"
)

func intToBytes(n int) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, uint64(n))
	return bytes
}

func TestMain(m *testing.M) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	RegisterNetworkListener("", 20000)

	defer CloseListeners()

	StartListeners()

	m.Run()
}

func TestNetworkListenerLifecycle(t *testing.T) {

	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		t.Fatalf("Failed to connect to listener: %v", err)
	}

	_, err = conn.Write(intToBytes(4))
	if err != nil {
		t.Fatalf("Failed to write size to connection: %v", err)
	}

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		t.Fatalf("Failed to write to connection: %v", err)
	}

	buf := make([]byte, 4)

	_, err = conn.Read(buf)

	if err != nil {
		t.Fatalf("Failed to read from connection: %v", err)
	}

	if string(buf) != "pong" {
		t.Fatalf("Expected 'pong', got %s", string(buf))
	}

}
