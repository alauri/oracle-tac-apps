/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)


var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the application.",
	Long: `configure the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "config called")
	},
}


var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "update section 'driver'",
	Long: `update section 'driver'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "[+] - Configuration updated")
	},
}


var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "update section 'database'",
	Long: `update section 'database'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "[+] - Configuration updated")
	},
}


func init() {
    configCmd.Flags().Bool("info", false, "print current stored configurations")

    // Register sub-commands
    configCmd.AddCommand(driverCmd)
    configCmd.AddCommand(databaseCmd)
}
