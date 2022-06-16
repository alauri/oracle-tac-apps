package cmd

import (
    "testing"
    "bytes"
    "runtime"
    "path"
    "io/ioutil"
    "os"

    "github.com/stretchr/testify/assert"
)


func copyConfiguraton(source string, destination string) {
    // Copy a source to a destination

    data, err := ioutil.ReadFile(source)
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(destination, data, 0644)
    if err != nil {
        panic(err)
    }
}

func setupSuite(tb testing.TB, st string) func(tb testing.TB, st string) {
    // Setup for the current module
    // Temporarly duplicate the TOML configuration file.
    src := path.Join(st, "config.toml")
    dst := path.Join(st, "config.bak")
    copyConfiguraton(src, dst)

	return func(tb testing.TB, st string) {
        // Teardown for the current module.
        // Delete temporary copy of the TOML configuration file.
        src := path.Join(st, "config.bak")
        dst := path.Join(st, "config.toml")
        copyConfiguraton(src, dst)
	    os.Remove(src)
    }
}

func Test_Package(t *testing.T) {
    threshold := 4
    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    teardownSuite := setupSuite(t, static)
	defer teardownSuite(t, static)

    parametrize := []struct{
        name string
        arguments []string
        target string
    }{
        {
            "Test_No_Args",
            []string{"-w", static, "config"},
            "Usage:",
        },
        {
            "Test_Info",
            []string{"-w", static, "config", "--info"},
            "Current configuration",
        },
        {
            "Test_Username",
            []string{"-w", static, "config", "driver", "--username", "fake"},
            "Configuration updated",
        },
        {
            "Test_Password",
            []string{"-w", static, "config", "driver", "--password", "fake"},
            "Configuration updated",
        },
        {
            "Test_Table",
            []string{"-w", static, "config", "database", "--table", "fake"},
            "Configuration updated",
        },
    }

    for _, test := range parametrize {
        t.Run(test.name, func(t *testing.T) {
            actual := new(bytes.Buffer)

            rootCmd.SetOut(actual)
            rootCmd.SetErr(actual)
            rootCmd.SetArgs(test.arguments)
            rootCmd.Execute()

            assert.Contains(t, actual.String(), test.target)
        })

        // For sub-commands only
        if len(test.arguments) > threshold {
            // assert.Equal(t, "database", test.arguments[3])
        }
    }
}
