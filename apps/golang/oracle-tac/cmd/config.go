/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command ``config`` helps managing the TOML configuration file by listing or
updating the current configuration.

It supports sub-commands to manage all the TOML sections with a Click's option
for each value of the section.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the application.",
	Long:  `configure the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		infoFlag, _ := cmd.Flags().GetBool("info")
		cfg := viper.GetViper()

		if infoFlag {
			fmt.Fprintf(cmd.OutOrStdout(), "[+] - Current configuration: %s\n", cfg.AllSettings())
		} else {
			cmd.Help()
		}
	},
}

var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "update section 'driver'",
	Long:  `update section 'driver'`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		cfg := viper.GetViper()

		// Update driver's username if needed
		if username != "" {
			cfg.Set("driver.username", username)
		}

		// Update driver's password if needed
		if password != "" {
			cfg.Set("driver.password", password)
		}

		// Update TOML configuration file
		cfg.WriteConfig()
		fmt.Fprintln(cmd.OutOrStdout(), "[+] - Configuration updated")
	},
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "update section 'database'",
	Long:  `update section 'database'`,
	Run: func(cmd *cobra.Command, args []string) {
		table, _ := cmd.Flags().GetString("table")
		cfg := viper.GetViper()

		// Update driver's password if needed
		if table != "" {
			cfg.Set("database.table", table)
		}

		// Update TOML configuration file
		cfg.WriteConfig()
		fmt.Fprintln(cmd.OutOrStdout(), "[+] - Configuration updated")
	},
}

func init() {
	// Config flags
	configCmd.Flags().Bool("info", false, "print current stored configurations")

	// Driver flags
	driverCmd.Flags().String("username", "", "update the driver's username")
	driverCmd.Flags().String("password", "", "update the driver's password")

	// Database flags
	databaseCmd.Flags().String("table", "", "update the database's table")

	// Register sub-commands
	configCmd.AddCommand(driverCmd)
	configCmd.AddCommand(databaseCmd)
}
