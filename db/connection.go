package db

import (
	"database/sql"
	"example/data-access/configs"
	"fmt"
	"log"
)

func GetConnection() (*sql.DB, error) {
	confs := configs.GetDBConfig()

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", confs.User, confs.Password, confs.Host, confs.Database)
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	return conn, err
}
