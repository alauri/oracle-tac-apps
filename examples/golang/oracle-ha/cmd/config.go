/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the application.",
	Long: `Configure the application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
