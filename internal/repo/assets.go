package repo

import (
	"context"
	"errors"
)

func (db *Database) StoreAsset(token, name string, data []byte) error {
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO assets (uid, name, data) VALUES ((SELECT id FROM sessions WHERE id=$1), $2, $3)", token, name, data)
	return err
}

func (db *Database) GetAsset(token, name string) ([]byte, error) {
	var data []byte
	err := db.Pool.QueryRow(context.Background(),
		"SELECT data FROM assets WHERE uid=(SELECT id FROM sessions WHERE id=$1) AND name=$2", token, name).Scan(&data)
	if err != nil {
		return nil, errors.New("asset not found")
	}
	return data, nil
}
