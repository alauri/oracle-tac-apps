package cmd

import (
    "testing"
    "bytes"
    "runtime"
    "path"
    "io/ioutil"
    "os"

    "github.com/stretchr/testify/assert"
	"github.com/spf13/viper"
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

type driver struct {
    Username string `toml:"username"`
    Password string `toml:"password"`
}

type database struct {
    Table string `toml:"table"`
}

type config struct {
    Driver driver `toml:"driver"`
    Database database `toml:"database"`
}

func readConfiguration(path string) config {
    // Read configuration TOML file

    var vp = viper.New()
    var ctoml config

    vp.AddConfigPath(path)
    vp.SetConfigName("config")
    vp.SetConfigType("toml")

    // Read the TOML file
    err := vp.ReadInConfig()
    if err != nil {
        panic(err)
    }

    // Convert the TOML file into struct
    err = vp.Unmarshal(&ctoml)
    if err != nil {
        panic(err)
    }
    return ctoml
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

        // TODO: find a way more elegant to revert conf file's values
        viper.Set("driver.username", "test")
        viper.Set("driver.password", "test")
        viper.Set("database.table", "test")
    }
}

func Test_Package(t *testing.T) {
    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    teardownSuite := setupSuite(t, static)
	defer teardownSuite(t, static)

    t.Run("Test_No_Args", func(t *testing.T) {
        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "config"})
        rootCmd.Execute()

        assert.Contains(t, actual.String(), "Usage:")
    })

    t.Run("Test_Info", func(t *testing.T) {
        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "config", "--info"})
        rootCmd.Execute()

        assert.Contains(t, actual.String(), "Current configuration")
        assert.NotContains(t, actual.String(), "Usage:")
    })

    t.Run("Test_Username", func(t *testing.T) {
        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "config", "driver", "--username", "fake"})
        rootCmd.Execute()

        infile := readConfiguration(static)
        assert.Equal(t, "fake", infile.Driver.Username)
    })

    t.Run("Test_Password", func(t *testing.T) {
        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "config", "driver", "--password", "fake"})
        rootCmd.Execute()

        infile := readConfiguration(static)
        assert.Equal(t, "fake", infile.Driver.Password)
    })

    t.Run("Test_Table", func(t *testing.T) {
        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "config", "database", "--table", "fake"})
        rootCmd.Execute()

        infile := readConfiguration(static)
        assert.Equal(t, "fake", infile.Database.Table)
    })
}
