/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “remove.go“
*/
package cmd_test

import (
  "bytes"
  "path"
  "runtime"
  "strings"
  "testing"
  
  "github.com/stretchr/testify/assert"
  "github.com/DATA-DOG/go-sqlmock"

  "github.com/alauri/oracle-tac-apps/cmd"
)

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

  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "remove",
		"--delay", "0.05"})
	cmd.RootCmd.Execute()

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

  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "remove",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "2"})
	cmd.RootCmd.Execute()

	assert.Equal(t, 5, strings.Count(actual.String(), "FROM json_tel WHERE LapTime='NaT'"))
	assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
