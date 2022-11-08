/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>
*/
package cmd

import "fmt"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset database data to factory.",
	Long:  `Reset database data to factory.`,
	Run: func(cmd *cobra.Command, args []string) {

		tableraw := viper.GetViper().GetString("database.tableraw")
		tablejson := viper.GetViper().GetString("database.tablejson")

		// Instrumentation: Set up module action (ACCHK_REPORT)
		// TODO: instrument action
		query := fmt.Sprintf("TRUNCATE TABLE %s", tableraw)
		// TODO: Execute the query
		cmd.Println(fmt.Sprintf("[+] - %s", query))

		// Instrumentation: Set up module action (ACCHK_REPORT)
		// TODO: instrument action
		query = fmt.Sprintf("ALTER TABLE %s", tableraw)
		// TODO: Execute the query
		cmd.Println(fmt.Sprintf("[+] - %s", query))

		// Instrumentation: Set up module action (ACCHK_REPORT)
		// TODO: instrument action
		query = fmt.Sprintf("TRUNCATE TABLE %s", tablejson)
		// TODO: Execute the query
		cmd.Println(fmt.Sprintf("[+] - %s", query))

		// Instrumentation: Set up module action (ACCHK_REPORT)
		// TODO: instrument action
		query = fmt.Sprintf("ALTER TABLE %s", tablejson)
		// TODO: Execute the query
		cmd.Println(fmt.Sprintf("[+] - %s", query))

		cmd.Println("[+] - All tables have been altered.")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
