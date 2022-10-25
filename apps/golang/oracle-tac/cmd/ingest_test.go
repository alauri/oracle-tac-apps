/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``ingest.go``
*/
package cmd

import (
	"bytes"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Ingest_No_Args(t *testing.T) {
	// Invoke the command ``ingest`` with no options.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "ingest"})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')")
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	// assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func Test_Ingest_With_Args(t *testing.T) {
	// Invoke the command ``ingest`` with the options to iterate the same
	// operation multiple times.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "ingest",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "5"})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')")
	assert.Contains(t, actual.String(), "0 days 00:01:29.103000|2|Car 1|Driver 1')")
	assert.Contains(t, actual.String(), "0 days 00:01:28.827000|3|Car 1|Driver 1')")
	assert.Contains(t, actual.String(), "0 days 00:01:29.026000|4|Car 1|Driver 1')")
	assert.Contains(t, actual.String(), "0 days 00:01:28.718000|5|Car 1|Driver 1')")
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	// assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
