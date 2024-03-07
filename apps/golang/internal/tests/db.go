package tests

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/alauri/oracle-tac-apps/internal/database"
)

// MockDatabase mock the database connection
func MockDatabase() (sqlmock.Sqlmock, func()) {
	// Create a database mock instance
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error in mocking the database: %s", err))
	}

	// Below are common operations to all the CLI commands
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	// Mocking the connection function by returning a mocked database
	// instance
	database.Connect = func(string, string, string) {
		database.Db = db
	}

	// TearDown function to return and use later within the tests
	return mock, func() {
		db.Close()
	}
}
