package tests

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

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

	suite.db = *database
	err = suite.db.Ping()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(">> Openned MySQL connection")
}

func (suite *DbTestSuite) TearDownSuite() {
	suite.db.Close()
	fmt.Println(">> Exiting MySQL Server")
}

func (suite *DbTestSuite) SetupTest() {
}

func (suite *DbTestSuite) TearDownTest() {
}

func (suite *DbTestSuite) TestCalc() {
	suite.Equal(1, 1)
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
		tableName = "mytable"
	)

	db := memory.NewDatabase(dbName)
	table := memory.NewTable(tableName, sqlMock.Schema{
		{Name: "name", Type: sqlMock.Text, Nullable: false, Source: tableName},
		{Name: "email", Type: sqlMock.Text, Nullable: false, Source: tableName},
		{Name: "phone_numbers", Type: sqlMock.JSON, Nullable: false, Source: tableName},
		{Name: "created_at", Type: sqlMock.Timestamp, Nullable: false, Source: tableName},
	})

	db.AddTable(tableName, table)
	ctx := sqlMock.NewEmptyContext()

	rows := []sqlMock.Row{
		sqlMock.NewRow("John Doe", "john@doe.com", []string{"555-555-555"}, time.Now()),
		sqlMock.NewRow("John Doe", "johnalt@doe.com", []string{}, time.Now()),
		sqlMock.NewRow("Jane Doe", "jane@doe.com", []string{}, time.Now()),
		sqlMock.NewRow("Evil Bob", "evilbob@gmail.com", []string{"555-666-555", "666-666-666"}, time.Now()),
	}

	for _, row := range rows {
		table.Insert(ctx, row)
	}

	return db
}
