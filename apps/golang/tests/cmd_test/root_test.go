/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “root.go“
*/
package cmd_test

import (
  "bytes"
  "path"
  "runtime"
  "testing"
  "strings"
  
  "github.com/stretchr/testify/assert"
  "github.com/DATA-DOG/go-sqlmock"

  "github.com/alauri/oracle-tac-apps/cmd"
)

func Test_Root_usage(t *testing.T) {
	// Invoke the CLI with no commands, expecting an 'Usage ...' message

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static})
  cmd.RootCmd.Execute()

	assert.Contains(t, actual.String(), "oracle-tac-go [flags]")
}

func Test_Root_config(t *testing.T) {
	// Invoke the CLI by asking information about the current configuration.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "--config"})
  cmd.RootCmd.Execute()

	assert.Contains(t, actual.String(), "{")
	assert.NotContains(t, actual.String(), "Usage:")
	assert.NotContains(t, actual.String(), "('server1', 'vm1')")
}

func Test_Root_ping(t *testing.T) {
	// Invoke the CLI by asking to ping the database.
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "--ping"})
  cmd.RootCmd.Execute()

	assert.Contains(t, actual.String(), "[+] - Database reachable")
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func Test_Root_error(t *testing.T) {
	// Invoke the CLI with a wrong dsn value.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "-d", "wrong", "ingest"})

	assert.PanicsWithValue(t, "Invalid value for '-d/--dsn'", func() { cmd.RootCmd.Execute() })
}
