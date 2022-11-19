/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Main command implementation.
*/
package cmd

import "encoding/json"
import "fmt"
import "os"
import "path"
import "runtime"
import "strings"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"

import "github.com/alauri/oracle-tac-apps/oracle-tac/database"

// Package's variables
var workdir string
var dsn string

var rootCmd = &cobra.Command{
	Use:   "oracle-tac-go",
	Short: "Oracle High Availability CLI in Golang",
	Long:  `Oracle High Availability CLI in Golang`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Retrieve database information
		username := viper.GetString("database.username")
		password := viper.GetString("database.password")
		dsn, _ := cmd.Flags().GetString("dsn")

		// Execution can't procede with an invalid dsn value
		dsn_cfg := strings.Join([]string{"dsn", dsn}, ".")
		if !viper.IsSet(dsn_cfg) {
			panic("Invalid value for '-d/--dsn'")
		}

		// Initialize the database instance within the variable Db
		database.Connect(username, password, dsn)

		// Get DB information before any other action with the database
		ping, _ := cmd.Flags().GetBool("ping")
		root_cmd := cmd.CalledAs() == "oracle-tac-go"
		if (root_cmd && ping) || !root_cmd {
			result, err := database.Get_Db_Info()
			if err != nil {
				panic(err)
			}
			cmd.Println("[+] - ", result)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		ping, _ := cmd.Flags().GetBool("ping")
		config, _ := cmd.Flags().GetBool("config")

		// Ping the database and exit
		if ping {
			// Try to ping the database
			err := database.Db.Ping()
			if err != nil {
				panic("[-] - Database not reachable")
			}
			cmd.Println("[+] - Database reachable")
			// Puts out the configuration and exit
		} else if config {
			config, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
			cmd.Println(string(config))
			// Default behavior to puts out the help message
		} else {
			cmd.Usage()
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Get DB information before any other action with the database
		ping, _ := cmd.Flags().GetBool("ping")
		root_cmd := cmd.CalledAs() == "oracle-tac-go"
		if (root_cmd && ping) || !root_cmd {
			result, err := database.Get_Db_Info()
			if err != nil {
				panic(err)
			}
			cmd.Println("[+] - ", result)
		}
		database.Db.Close()
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

func initConfig() {
	/* Initialization function, called before any command execution. */

	// Read configuration file
	viper.AddConfigPath(workdir)
	viper.SetConfigType("toml")
	viper.SetConfigName("config")

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Error reading confugration file: %s", err))
	}
}

func init() {
	/* Define main command's hooks and flags. */

	// CLI hooks
	cobra.OnInitialize(initConfig)

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
