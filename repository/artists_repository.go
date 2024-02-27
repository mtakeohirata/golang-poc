package repository

import (
	"database/sql"
	"fmt"

	"example/data-access/models"
)

type ArtistsRepository struct {
	db *sql.DB
}

func NewArtistsRepository(db *sql.DB) *ArtistsRepository {
	return &ArtistsRepository{
		db: db,
	}
}

func FetchAll(r *ArtistsRepository) (artists []models.Artist, err error) {
	dml := "SELECT * FROM ARTIST"
	rows, err := r.db.Query(dml)
	if err != nil {
		return
	}

	for rows.Next() {
		var artist models.Artist
		err = rows.Scan(&artist.ID, &artist.Name)
		artists = append(artists, artist)
		if err != nil {
			continue
		}
	}

	return
}

func FetchById(r *ArtistsRepository, id int) (artist models.Artist, err error) {
	dml := "SELECT * FROM ARTIST where id = ?"
	rows, err := r.db.Query(dml)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		fmt.Println(artist)
		err = rows.Scan(&artist.ID, &artist.Name)
		if err != nil {
			continue
		}
	}

	return
	// dml := "SELECT * FROM ARTIST WHERE ID = 1"
	// fmt.Println(id)
	// row := r.db.QueryRow(dml, id)
	// fmt.Println(row)
	// err = row.Scan(&artist.ID, artist.Name)

	// return
}
