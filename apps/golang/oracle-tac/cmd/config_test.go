/* Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>
*
* Unit tests utility for the package `cmd`.
 */
package cmd

import "fmt"
import "database/sql"
import "testing"

import "github.com/DATA-DOG/go-sqlmock"

func setUp(t *testing.T) (sqlmock.Sqlmock, func(t *testing.T)) {
	// Create a database mock instance
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Sprintf("Error in mocking the database: %s", err))
	}

	// Mocking the connection function by returning a mocked database
	// instance
	connect = func(username string, password string, dsn string) *sql.DB {
		return db
	}

	// TearDown function to return and use later within the tests
	return mock, func(t *testing.T) {
		db.Close()
	}
}
