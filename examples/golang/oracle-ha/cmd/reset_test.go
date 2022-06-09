package cmd

import (
    "testing"
    "bytes"

    "github.com/stretchr/testify/assert"
)


func Test_ResetCommand(t *testing.T) {
    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"reset"})
    rootCmd.Execute()

    expected := "reset called\n"
    assert.Equal(t, expected, actual.String())
}
