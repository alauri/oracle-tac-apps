/* Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>
*
* Unit tests utility for the package `cmd`.
 */
package cmd

import "fmt"
import "testing"

import "github.com/DATA-DOG/go-sqlmock"
import "github.com/alauri/oracle-tac-apps/oracle-tac/database"

func setUpDatabase(t *testing.T) (sqlmock.Sqlmock, func(t *testing.T)) {
	// Create a database mock instance
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error in mocking the database: %s", err))
	}

	// Below are common operations to all the CLI commands
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)
	rows = sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	// Mocking the connection function by returning a mocked database
	// instance
	database.Connect = func(username string, password string, dsn string) {
		database.Db = db
	}

	// TearDown function to return and use later within the tests
	return mock, func(t *testing.T) {
		db.Close()
	}
}
