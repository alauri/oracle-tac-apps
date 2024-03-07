/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “remove.go“
*/
package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/alauri/oracle-tac-apps/internal/tests"
)

func TestRemoveNoArgs(t *testing.T) {
	// Invoke the command ``remove`` with no options.
	mock, tearDownDatabase := tests.MockDatabase()
	defer tearDownDatabase()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{
		"-w", static, "-d", "localhost", "remove",
		"--delay", "0.05",
	})
	RootCmd.Execute()

	assert.Equal(t, 1, strings.Count(actual.String(), "FROM json_tel WHERE LapTime='NaT'"))
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func TestRemoveWithArgs(t *testing.T) {
	/* Invoke the command ``remove`` with the options to iterate the same
	operation multiple times. */
	mock, tearDownDatabase := tests.MockDatabase()
	defer tearDownDatabase()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{
		"-w", static, "-d", "localhost", "remove",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "2",
	})
	RootCmd.Execute()

	assert.Equal(t, 5, strings.Count(actual.String(), "FROM json_tel WHERE LapTime='NaT'"))
	assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
