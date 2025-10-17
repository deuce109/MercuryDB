package auth

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/deuce109/mercurydb/v2/db/config"
)

func setenv(t *testing.T) {
	t.Setenv("MERCURY_PORT", "0")
	t.Setenv("MERCURY_INTERFACE_IP", "0.0.0.0")
	t.Setenv("MERCURY_SOCKET_PATH", "./")
}

func authHandlerLifecycle(t *testing.T) {
	setenv(t)
	config.Load()

	t.Log(config.GetUserDbPath())

	handler, err := NewAuthHandler()

	if err != nil {
		t.Fatalf("Expected err to be nil got %v", err)
	}

	if handler == nil {
		t.Fatal("auth handler is null")
	}

	defer handler.Close()

	if created := handler.MakeUser("test", "test"); !created {
		t.Fatal("failed creating user")
	}

	if verified := handler.VerifyUser("test", "test"); !verified {
		t.Fatal("failed verifying user")
	}

}

func TestAuthHandlerLifecycleWithInMemoryDB(t *testing.T) {
	t.Setenv("MERCURY_USER_DB_PATH", "")
	authHandlerLifecycle(t)

}

func TestAuthHandlerLifecycleWithFileDB(t *testing.T) {
	t.Setenv("MERCURY_USER_DB_PATH", "./test.db")
	authHandlerLifecycle(t)

	dbPath, _ := filepath.Abs("./test.db")

	err := os.Remove(dbPath)

	if err != nil {
		t.Logf("Error removing test.db: %v", err)
	}
}
