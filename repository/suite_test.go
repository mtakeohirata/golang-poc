package repository

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	sqle "github.com/src-d/go-mysql-server"
	"github.com/src-d/go-mysql-server/auth"
	"github.com/src-d/go-mysql-server/memory"
	"github.com/src-d/go-mysql-server/server"
	sqlMock "github.com/src-d/go-mysql-server/sql"
	"github.com/stretchr/testify/suite"
)

type DbTestSuite struct {
	suite.Suite
	db sql.DB
}

func TestDbSuite(t *testing.T) {
	suite.Run(t, new(DbTestSuite))
}

func (suite *DbTestSuite) SetupSuite() {
	go CreateMemDB()
	database, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	err = database.Ping()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(">> Openned MySQL connection")
	suite.db = *database
	// time.Sleep(20 * time.Second)
}

func (suite *DbTestSuite) TearDownSuite() {
	suite.db.Close()
	fmt.Println(">> Exiting MySQL Server")
}

func (suite *DbTestSuite) SetupTest() {
}

func (suite *DbTestSuite) TearDownTest() {
}

func CreateMemDB() {
	driver := sqle.NewDefault()
	driver.AddDatabase(CreateTestTable())

	config := server.Config{
		Protocol: "tcp",
		Address:  "localhost:3306",
		Auth:     auth.NewNativeSingle("user", "pass", auth.AllPermissions),
	}

	s, err := server.NewDefaultServer(config, driver)
	if err != nil {
		panic(err)
	}

	fmt.Println(">> Starting MySQL Server")
	s.Start()
}

func CreateTestTable() *memory.Database {
	const (
		dbName    = "test"
		tableName = "ARTIST"
	)

	db := memory.NewDatabase(dbName)
	table := memory.NewTable(tableName, sqlMock.Schema{
		{Name: "NAME", Type: sqlMock.Text, Nullable: false, Source: tableName},
		{Name: "ID", Type: sqlMock.Int64, Nullable: false, Source: tableName},
	})

	db.AddTable(tableName, table)
	ctx := sqlMock.NewEmptyContext()

	rows := []sqlMock.Row{
		sqlMock.NewRow("John Doe", 1),
		sqlMock.NewRow("Takeo", 2),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}
