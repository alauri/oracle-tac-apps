/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``cleanup.go``
*/
package cmd

import (
	"bytes"
	"path"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cleanup_No_Args(t *testing.T) {
	// Invoke the command ``cleanup`` with no options.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup"})
	rootCmd.Execute()

        assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
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

        assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
        assert.Contains(t, actual.String(), ",'00:01:29.103000',2,'Car 1','Driver 1')")
        assert.Contains(t, actual.String(), ",'00:01:28.827000',3,'Car 1','Driver 1')")
        assert.Contains(t, actual.String(), ",'00:01:29.026000',4,'Car 1','Driver 1')")
        assert.Contains(t, actual.String(), ",'00:01:28.718000',5,'Car 1','Driver 1')")
}
