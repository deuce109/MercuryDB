package auth

import (
	"log/slog"

	"github.com/deuce109/mercurydb/v2/db/config"
)

type authenticationHandler struct {
	db *authDatabase
}

func (a *authenticationHandler) VerifyUser(username string, password string) bool {
	user, err := a.db.getUser(username)
	if err != nil {
		slog.Error("Error finding user", "username", username, "error", err)
		return false
	}

	return user != nil && user.verifyPassword(password)
}

func (a *authenticationHandler) MakeUser(username string, password string) bool {
	u := &user{
		username: username,
	}

	if err := u.generatePasswordHash(password); err != nil {
		slog.Error("Error generating new user", "error", err)
		return false
	}

	err := a.db.insertUser(u)
	if err != nil {
		slog.Error("Error insterting user into database", "error", err)
		return false
	}
	return true
}

func (a *authenticationHandler) Close() error {
	if a.db != nil {
		return a.db.close()
	}
	return nil
}

func NewAuthHandler() (*authenticationHandler, error) {
	db, err := newAuthDatabase(config.GetUserDbPath())

	if err != nil {
		return nil, err
	}

	return &authenticationHandler{
		db: db,
	}, nil
}
