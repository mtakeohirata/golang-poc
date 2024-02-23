package main

import (
	"database/sql"
	"example/data-access/configs"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Takeo", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("albums", Persist)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func Persist(c *gin.Context) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", "postgres", "password", "localhost", "golang")
	conn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	log.Print(err)

	if err != nil {
		return
	}

	defer conn.Close()

	dml := "INSERT INTO public.ARTISTS (NAME) VALUES ($1)"

	row, err := conn.Query(dml, "takeo")
	log.Print(err)

	c.IndentedJSON(http.StatusOK, row)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

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
