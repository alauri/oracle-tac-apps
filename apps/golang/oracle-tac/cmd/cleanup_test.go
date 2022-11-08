/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “cleanup.go“
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "strings"
import "testing"

import "github.com/stretchr/testify/assert"

func Test_Cleanup_No_Args(t *testing.T) {
	// Invoke the command ``cleanup`` with no options.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup"})
	rootCmd.Execute()

	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
}

func Test_Cleanup_With_Args(t *testing.T) {
	// Invoke the command ``cleaup`` with the options to iterate the same
	// operation multiple times.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "2"})
	rootCmd.Execute()

	assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
}
