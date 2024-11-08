package repo

import (
	"context"
	"errors"
)

func (db *Database) GetUserToken(login, password string) (string, error) {
	var userID int64

	err := db.Pool.QueryRow(context.Background(),
		"SELECT id FROM users WHERE login=$1 AND password_hash=md5($2)", login, password).Scan(&userID)
	if err != nil {
		return "", errors.New("invalid login/password")
	}

	var token string
	err = db.Pool.QueryRow(context.Background(),
		"INSERT INTO sessions (uid) VALUES ($1) RETURNING id", userID).Scan(&token)
	if err != nil {
		return "", errors.New("failed to create session")
	}

	return token, nil
}
