package models

import (
	"example/data-access/db"
)

var openConnection = db.GetConnection

func Persist(artist Artist) (id int64, err error) {
	conn, err := openConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	dml := "INSERT INTO ARTISTS (NAME) VALUES ($1) RETURNING ID"

	err = conn.QueryRow(dml, artist.Name).Scan(&id)

	return
}
