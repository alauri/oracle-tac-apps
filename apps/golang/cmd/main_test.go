package cmd

import (
	"os"
	"path"
	"runtime"
	"testing"
)

var static string

// Main test function
func TestMain(m *testing.M) {
	// Setup
	_, filename, _, _ := runtime.Caller(0)
	static = path.Join(path.Dir(filename), "../internal/tests/static")

	execVal := m.Run()
	os.Exit(execVal)
}
