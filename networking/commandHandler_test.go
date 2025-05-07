package networking

import (
	"testing"
)

func TestDefaultCommandHandler(t *testing.T) {
	commandHandler := DefaultCommandHandler()
	if commandHandler == nil {
		t.Fatalf("DefaultCommandHandler returned nil")
	}
}

func TestCommandLifecycle(t *testing.T) {

	commandHandler := DefaultCommandHandler()

	args := []any{1, 2, 3}

	value, err := commandHandler.HandleCommand("ping", args...)

	if err != nil {
		t.Fatalf("HandleCommand returned an error: %v", err)
	}

	if value != "pong" {
		t.Fatalf("Expected 'pong', got %v", value)
	}
}

func TestCommandNotFound(t *testing.T) {
	commandHandler := DefaultCommandHandler()

	_, err := commandHandler.HandleCommand("nonexistent", 1, 2, 3)
	if err == nil {
		t.Fatalf("HandleCommand should have returned an error for nonexistent command")

	}
}
