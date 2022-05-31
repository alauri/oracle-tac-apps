/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete records from the table.",
	Long: `Delete records from the table.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
