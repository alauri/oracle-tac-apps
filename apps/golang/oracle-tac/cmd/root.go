/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Main command implementation.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var workdir string
var dsn string

var rootCmd = &cobra.Command{
	Use:   "oracle-tac-go",
	Short: "Oracle High Availability CLI in Golang",
	Long:  `Oracle High Availability CLI in Golang`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Check the given DSN
		dsn, _ := cmd.Flags().GetString("dsn")
		dsn = strings.Join([]string{"dsn", dsn}, ".")

		// Execution can't procede with an invalid dsn value
		if !viper.IsSet(dsn) {
			panic("Invalid value for '-d/--dsn'")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ping, _ := cmd.Flags().GetBool("ping")
		config, _ := cmd.Flags().GetBool("config")

		// Ping the database and exit
		if ping {
			cmd.Println("[+] - Database reachable")
		// Puts out the configuration and exit
		} else if config {
			config, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
			cmd.Println(string(config))
                // Default behavior to puts out the help message
		} else { cmd.Help() }
	},
}

func HandleError() {
        /* Handle `panic` error. */
	if err := recover(); err != nil {
		rootCmd.Println(fmt.Sprintf("Error occurred: %s", err))
	}
}

func Execute() {
	/* Main command entry point.

        Set up a defer function to handle any possible error occurred during
        command execution.
        */
	defer HandleError()

        // Start the main command
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setUp() {
        /* Initialization function, called before any command execution. */

	// Read configuration file
	viper.AddConfigPath(workdir)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		rootCmd.Println(fmt.Sprintf("Error reading confugration file: %s", err))
		os.Exit(1)
	}
}

func tearDown() {
	/* Finalizing function, called after any command execution. */
	rootCmd.Println("Finalizing...")
}

func init() {
        /* Define main command's hooks and flags. */

	// Define CLI hooks
	cobra.OnInitialize(setUp)
	cobra.OnFinalize(tearDown)

	// Default value for the working directory
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")

	// Persistent flags, these can be forwarded to subcommands
	rootCmd.PersistentFlags().StringVarP(&workdir, "workdir", "w", dir,
		"The absolute path of the configuration folder")
	rootCmd.PersistentFlags().StringVarP(&dsn, "dsn", "d", "localhost",
		"The connection string to use")

	// Local flags, these flags can not be used by subcommands
	rootCmd.Flags().Bool("config", false, "Show the current configuration")
	rootCmd.Flags().Bool("ping", false, "Check database connection")
}
