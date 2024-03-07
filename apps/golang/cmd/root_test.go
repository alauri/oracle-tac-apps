/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “root.go“
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

func TestRootUsage(t *testing.T) {
	// Invoke the CLI with no commands, expecting an 'Usage ...' message

	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), "oracle-tac-go [flags]")
}

func TestRootConfig(t *testing.T) {
	// Invoke the CLI by asking information about the current configuration.

	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "--config"})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), "{")
	assert.NotContains(t, actual.String(), "Usage:")
	assert.NotContains(t, actual.String(), "('server1', 'vm1')")
}

func TestRootPing(t *testing.T) {
	// Invoke the CLI by asking to ping the database.
	mock, tearDownDatabase := tests.MockDatabase()
	defer tearDownDatabase()

	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "--ping"})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), "[+] - Database reachable")
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func TestRootError(t *testing.T) {
	// Invoke the CLI with a wrong dsn value.

	actual := new(bytes.Buffer)
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "-d", "wrong", "ingest"})

	assert.PanicsWithValue(t, "Invalid value for '-d/--dsn'", func() { RootCmd.Execute() })
}
