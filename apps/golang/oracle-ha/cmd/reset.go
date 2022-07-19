/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset database data to factory.",
	Long:  `Reset database data to factory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "[+] - Database has been reset")
	},
}
