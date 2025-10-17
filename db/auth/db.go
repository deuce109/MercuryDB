package auth

import (
	"context"
	"database/sql"
	"path/filepath"

	_ "modernc.org/sqlite"
)

var INIT_SCRIPT string = `CREATE TABLE IF NOT EXISTS users (username TEXT PRIMARY KEY, password TEXT, salt TEXT);`
var GET_USER_SCRIPT string = `SELECT username, password, salt FROM users WHERE username = ?;`
var INSERT_USER_SCRIPT string = `INSERT INTO users VALUES (?, ?, ?);`

type authDatabase struct {
	db *sql.DB
}

func (a *authDatabase) executeTransaction(script string, args ...any) (sql.Result, error) {
	ctx := context.Background()
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	result, err := tx.ExecContext(ctx, script, args...)
	if err != nil {
		return nil, err
	}

	return result, tx.Commit()
}

func (a *authDatabase) executeScript(script string, args ...any) (*sql.Rows, error) {
	ctx := context.Background()

	return a.db.QueryContext(ctx, script, args...)
}

func (a *authDatabase) init() error {
	_, err := a.executeTransaction(INIT_SCRIPT)
	return err
}

func (a *authDatabase) getUser(username string) (*user, error) {
	result, err := a.executeScript(GET_USER_SCRIPT, username)
	if err != nil {
		return nil, err
	}

	values, err := getResultsFromRows(result)
	if len(values) < 1 {
		return nil, nil
	}
	return values[0], err
}

func (a *authDatabase) insertUser(u *user) error {
	_, err := a.executeTransaction(INSERT_USER_SCRIPT, u.username, u.password, u.salt)
	return err
}

func (a *authDatabase) close() error {
	if a.db != nil {
		return a.db.Close()
	}
	return nil
}

func newAuthDatabase(path string) (*authDatabase, error) {
	var absPath string
	var err error
	if path != "" && path != ":memory:" {
		absPath, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	} else {
		absPath = ":memory:"
	}
	sqliteDb, err := sql.Open("sqlite", absPath)
	if err != nil {
		return nil, err
	}

	db := &authDatabase{
		db: sqliteDb,
	}

	err = db.init()

	return db, err
}

func getResultsFromRows(rows *sql.Rows) ([]*user, error) {
	results := make([]*user, 0)
	var u *user
	var err error
	u, err = getNextResultFromRows(rows)
	for u != nil && err == nil {
		results = append(results, u)
		u, err = getNextResultFromRows(rows)
	}

	return results, nil
}

func getNextResultFromRows(rows *sql.Rows) (*user, error) {
	var username string
	var password string
	var salt string
	if !rows.Next() {
		return nil, rows.Err()
	}
	err := rows.Scan(&username, &password, &salt)
	if err != nil {
		return nil, err
	}

	u := &user{
		username: username,
		password: password,
		salt:     salt,
	}
	return u, nil
}
