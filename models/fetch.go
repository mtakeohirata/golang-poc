package models

import "example/data-access/db"

func FetchAll() (artist []Artist, err error) {
	conn, err := db.GetConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	dml := "SELECT * FROM ARTIST"

	rows, err := conn.Query(dml)
	if err != nil {
		return
	}

	for rows.Next() {
		var artist Artist
		err = rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			continue
		}
	}

	return
}

func FetchById(id int64) (artist Artist, err error) {
	conn, err := db.GetConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	dml := "SELECT * FROM ARTIST WHERE ID=$1"
	row := conn.QueryRow(dml, id)
	err = row.Scan(&artist.ID, artist.Name)

	return
}
