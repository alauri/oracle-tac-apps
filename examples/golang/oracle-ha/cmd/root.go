/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"os"
    "fmt"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oracle-ha",
	Short: "Oracle High Availability CLI in Golang",
	Long: `Oracle High Availability CLI in Golang`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Fprintln(cmd.OutOrStdout(), "root called")
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


