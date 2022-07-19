/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``root.go``
*/
package cmd

import (
	"bytes"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Root(t *testing.T) {
	// Invoke the main command.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)
	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "5"})
	rootCmd.Execute()

	expected := "Use \"oracle-ha [command] --help"
	assert.Contains(t, actual.String(), expected)
}
