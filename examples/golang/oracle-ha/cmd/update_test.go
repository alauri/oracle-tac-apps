/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package ``update.go``
*/
package cmd

import (
    "testing"
    "bytes"
    "runtime"
    "path"
    "strings"

    "github.com/stretchr/testify/assert"
)


func Test_Update_No_Args(t *testing.T) {
    // Invoke the command ``update`` with no options.

    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"-w", static, "update"})
    rootCmd.Execute()

    expected := []string{"[1/1] - UPDATE test SET arg=1 WHERE id=2",
                         "[1/1] - COMMIT"}

    splitFn := func(c rune) bool { return c == '\n' }
    assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}


func Test_Update_With_Args(t *testing.T) {
    // Invoke the command ``update`` with the options to iterate the same
    // operation multiple times.

    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"-w", static, "update", "--iters", "5", "--delay", "0.05", "--commit-every", "2"})
    rootCmd.Execute()

    expected := []string{"[1/5] - UPDATE test SET arg=1 WHERE id=2",
                         "[2/5] - UPDATE test SET arg=2 WHERE id=3",
                         "[2/5] - COMMIT",
                         "[3/5] - UPDATE test SET arg=3 WHERE id=4",
                         "[4/5] - UPDATE test SET arg=4 WHERE id=5",
                         "[4/5] - COMMIT",
                         "[5/5] - UPDATE test SET arg=5 WHERE id=6",
                         "[5/5] - COMMIT"}
    splitFn := func(c rune) bool { return c == '\n' }
    assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}
