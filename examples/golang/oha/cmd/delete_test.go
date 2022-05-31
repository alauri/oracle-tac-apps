package cmd


import (
    "testing"
    "bytes"

    "github.com/stretchr/testify/assert"
)


func Test_DeleteCommand(t *testing.T) {
    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"delete"})
    rootCmd.Execute()

    expected := "delete called\n"
    assert.Equal(t, expected, actual.String())
}
