/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “root.go“
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "testing"

import "github.com/stretchr/testify/assert"

func Test_Root_usage(t *testing.T) {
	// Invoke the CLI with no commands, expecting an 'Usage ...' message

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "oracle-tac-go [flags]")
}

func Test_Root_config(t *testing.T) {
	// Invoke the CLI by asking information about the current configuration.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "--config"})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "{")
	assert.NotContains(t, actual.String(), "Usage:")
}

func Test_Root_ping(t *testing.T) {
	// Invoke the CLI by asking to ping the database.
	_, tearDown := setUp(t)
	defer tearDown(t)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "--ping"})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "[+] - Database reachable")
}

func Test_Root_error(t *testing.T) {
	// Invoke the CLI with a wrong dsn value.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "wrong", "ingest"})

	assert.PanicsWithValue(t, "Invalid value for '-d/--dsn'", func() { rootCmd.Execute() })
}
