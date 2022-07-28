/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``insert.go``
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

func Test_Insert_No_Args(t *testing.T) {
	// Invoke the command ``insert`` with no options.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "insert"})
	rootCmd.Execute()

	expected := []string{"[1/1] - INSERT INTO test(id) VALUES(1)",
		"[1/1] - COMMIT"}

	splitFn := func(c rune) bool { return c == '\n' }
	assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}

func Test_Insert_With_Args(t *testing.T) {
	// Invoke the command ``insert`` with the options to iterate the same
	// operation multiple times.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "insert", "--iters", "5", "--delay", "0.05", "--commit-every", "2"})
	rootCmd.Execute()

	expected := []string{"[1/5] - INSERT INTO test(id) VALUES(1)",
		"[2/5] - INSERT INTO test(id) VALUES(2)",
		"[2/5] - COMMIT",
		"[3/5] - INSERT INTO test(id) VALUES(3)",
		"[4/5] - INSERT INTO test(id) VALUES(4)",
		"[4/5] - COMMIT",
		"[5/5] - INSERT INTO test(id) VALUES(5)",
		"[5/5] - COMMIT"}
	splitFn := func(c rune) bool { return c == '\n' }
	assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}
