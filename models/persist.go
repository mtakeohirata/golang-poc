package models

import (
	"example/data-access/db"
)

func Persist() (id int64, err error) {
	conn, err := db.GetConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	dml := "INSERT INTO ARTISTS (NAME) VALUES ($1) RETURNING ID"

	err = conn.QueryRow(dml, "takeo").Scan(&id)

	return
}
