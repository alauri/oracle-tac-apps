/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command “remove“ is used to delete one or more records from the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

import "fmt"
import "time"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"

import "github.com/alauri/oracle-tac-apps/internal/database"

// removeCmd represents the delete command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Delete records from the table.",
	Long:  `Delete records from the table.`,
	Run: func(cmd *cobra.Command, args []string) {
		// dsn, _ := cmd.Flags().GetInt("dsn")
		loop, _ := cmd.Flags().GetBool("loop")
		iters, _ := cmd.Flags().GetInt("iters")
		delay, _ := cmd.Flags().GetFloat64("delay")
		commit_every, _ := cmd.Flags().GetInt("commit-every")

		// Define query parameters
		table := viper.GetViper().GetString("database.tablejson")

		step := 1

		tx, err := database.Db.Begin()
		if err != nil {
			panic(err)
		}

		for {
			// Exit condition
			if !loop && step > iters {
				break
			}

			// Prepare query with updated conditions
			query := fmt.Sprintf("DELETE FROM %s WHERE LapTime='NaT'", table)
			tx.Exec(query)
			fmt.Fprintln(cmd.OutOrStdout(),
				fmt.Sprintf("[%d/%d] - %s", step, iters, query))

			// Commit changes
			if step%commit_every == 0 {
				tx.Commit()
				fmt.Fprintln(cmd.OutOrStdout(),
					fmt.Sprintf("[%d/%d] - COMMIT", step, iters))
			}

			step += 1
			time.Sleep(time.Duration(delay) * time.Second)
		}

		// Chech the last commit
		if iters%commit_every != 0 {
			tx.Commit()
			fmt.Fprintln(cmd.OutOrStdout(),
				fmt.Sprintf("[%d/%d] - COMMIT", iters, iters))
		}
	},
}

func init() {
	RootCmd.AddCommand(RemoveCmd)

	RemoveCmd.Flags().Bool("loop", false, "repeat the same operation forever")
	RemoveCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	RemoveCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	RemoveCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
