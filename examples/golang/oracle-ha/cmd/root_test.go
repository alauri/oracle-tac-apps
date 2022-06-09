package cmd

import (

	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)


func Test_CommandRoot(t *testing.T) {

    actual := new(bytes.Buffer)
    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{})
	rootCmd.Execute()

	expected := "root called\n"
	assert.Equal(t, expected, actual.String())
}
