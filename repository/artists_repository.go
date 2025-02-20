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
	dml := "SELECT * FROM ARTIST where id = $1"
	row := r.db.QueryRow(dml, id)
	if err != nil {
		fmt.Println(err)
	}

	err = row.Scan(&artist.ID, &artist.Name)
	if err != nil {
		fmt.Println(err)
	}

	return
}
