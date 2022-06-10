/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"os"
    "path"
    "runtime"

	"github.com/spf13/cobra"
)


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "oracle-ha",
	Short: "Oracle High Availability CLI in Golang",
	Long: `Oracle High Availability CLI in Golang`,
    PreRun: func(cmd *cobra.Command, args []string) {
        // Print help with no sub-commands invokation
        if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
    },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
    _, filename, _, _ := runtime.Caller(0)
    workdir := path.Join(path.Dir(filename), "../../..")
    rootCmd.PersistentFlags().StringP("workdir", "w", workdir,
                                      "the absolute path of the configuration folder")

    // Register sub-commands
    rootCmd.AddCommand(configCmd)
    rootCmd.AddCommand(deleteCmd)
    rootCmd.AddCommand(insertCmd)
    rootCmd.AddCommand(resetCmd)
    rootCmd.AddCommand(updateCmd)
}
