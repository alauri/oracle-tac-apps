package cmd

import (
    "testing"
    "bytes"

    "github.com/stretchr/testify/assert"
)


func Test_UpdateCommand(t *testing.T) {
    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"update"})
    rootCmd.Execute()

    expected := "update called\n"
    assert.Equal(t, expected, actual.String())
}
