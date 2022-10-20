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
var dsn string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:    "oracle-tac-go",
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
		"The absolute path of the configuration folder")
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "localhost",
		"The connection string to use")
	rootCmd.Flags().Bool("config", false, "Show the current configuration")
	rootCmd.Flags().Bool("ping", false, "Check database connection")

	// Register sub-commands
	rootCmd.AddCommand(cleanupCmd)
	rootCmd.AddCommand(ingestCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(resetCmd)
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
