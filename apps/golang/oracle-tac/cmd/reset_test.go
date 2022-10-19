/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``reset.go``
*/
package cmd

import (
	"bytes"
	"path"
	"runtime"
	"testing"
        "strings"

	"github.com/stretchr/testify/assert"
)

func Test_Reset_No_Args(t *testing.T) {
	// Invoke the command ``reset`` with no options.

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "reset"})
	rootCmd.Execute()

        // assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
        assert.Equal(t, 2, strings.Count(actual.String(), "TRUNCATE TABLE"))
        assert.Equal(t, 2, strings.Count(actual.String(), "ALTER TABLE"))
        assert.Contains(t, "All tables have been altered.", actual.String())
}
