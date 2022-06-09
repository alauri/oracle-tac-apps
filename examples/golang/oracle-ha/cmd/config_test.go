package cmd

import (
    "testing"
    "bytes"

    "github.com/stretchr/testify/assert"
)

func Test_CommandConfig(t *testing.T) {
    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"config"})
    rootCmd.Execute()

    expected := "config called\n"
    assert.Equal(t, expected, actual.String())
}
