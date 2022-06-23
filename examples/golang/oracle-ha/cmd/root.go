/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"os"
	"path"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var workdir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:    "oracle-ha",
	Short:  "Oracle High Availability CLI in Golang",
	Long:   `Oracle High Availability CLI in Golang`,
	PreRun: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")

	// Persistent flags
	rootCmd.PersistentFlags().StringVarP(&workdir, "workdir", "w", dir,
		"the absolute path of the configuration folder")
	rootCmd.PersistentFlags().IntP("dsn", "d", 1,
		"the connection string to use")

	// Register sub-commands
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(insertCmd)
	rootCmd.AddCommand(resetCmd)
	rootCmd.AddCommand(updateCmd)
}

func initConfig() {
	viper.AddConfigPath(workdir)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
