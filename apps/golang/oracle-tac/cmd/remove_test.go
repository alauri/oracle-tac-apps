/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “remove.go“
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "strings"
import "testing"

import "github.com/stretchr/testify/assert"
import "github.com/DATA-DOG/go-sqlmock"

func Test_Remove_No_Args(t *testing.T) {
	// Invoke the command ``remove`` with no options.
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "remove",
		"--delay", "0.05"})
	rootCmd.Execute()

	assert.Equal(t, 1, strings.Count(actual.String(), "FROM json_tel WHERE LapTime='NaT'"))
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func Test_Remove_With_Args(t *testing.T) {
	/* Invoke the command ``remove`` with the options to iterate the same
	operation multiple times. */
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	mock.ExpectBegin()
        mock.ExpectExec("DELETE FROM")
        mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)


	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "remove",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "2"})
	rootCmd.Execute()

	assert.Equal(t, 5, strings.Count(actual.String(), "FROM json_tel WHERE LapTime='NaT'"))
	assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
