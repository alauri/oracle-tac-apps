package cmd

import (
    "testing"
    "bytes"

    "github.com/stretchr/testify/assert"
)


func Test_InsertCommand(t *testing.T) {
    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"insert"})
    rootCmd.Execute()

    expected := "insert called\n"
    assert.Equal(t, expected, actual.String())
}
