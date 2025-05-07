package networking

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestNetworkListenerLifecycle(t *testing.T) {
	listener := NewNetworkListener("", 20000)

	defer listener.Close()

	go listener.Listen(func() { fmt.Printf("Listening") })

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:20000")

	if err != nil {
		t.Fatalf("Failed to connect to listener: %v", err)
	}

	defer conn.Close()

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

func TestNetworkListenerStartFailure(t *testing.T) {
	list, err := net.Listen("tcp", ":20000")

	if err != nil {
		t.Fatalf("Failed to start listener: %v", err)
	}

	defer list.Close()

	listener := NewNetworkListener("", 20000)
	err = listener.Listen(func() { fmt.Printf("Listening") })
	if err == nil {
		t.Fatal("Expected error when starting listener on already used port, got nil")
	}

}
