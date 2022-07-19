/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command ``update`` to change already existing records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alauri/oracle-ha-apps/oracle-ha/db"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update records into the table.",
	Long:  `Update records into the table.`,
	Run: func(cmd *cobra.Command, args []string) {
		dsn, _ := cmd.Flags().GetInt("dsn")
		loop, _ := cmd.Flags().GetBool("loop")
		iters, _ := cmd.Flags().GetInt("iters")
		delay, _ := cmd.Flags().GetFloat64("delay")
		commit_every, _ := cmd.Flags().GetInt("commit-every")

		// Define query parameters
		table := viper.GetViper().GetString("database.table")
		conditions := Conds{1}
		arguments := Arguments{0}

		// Retrieve a fresh Database connection
		conn, err := db.GetDatabase(dsn)
		if err != nil {
			panic(err)
		}

		if loop {
			iters = 0
		}
		step := 1
		for {
			// Exit condition
			if !loop && step > iters {
				break
			}

			// Prepare query with updated conditions
			sets := Arguments{arguments.arg + step}
			wheres := Conds{conditions.id + step}
			query := fmt.Sprintf("UPDATE %s SET arg=%d WHERE id=%d",
				table, sets.arg, wheres.id)

			// Perform the query
			db.DoQuery(conn, query)
			fmt.Fprintln(cmd.OutOrStdout(),
				fmt.Sprintf("[%d/%d] - %s", step, iters, query))

			// Commit changes
			if step%commit_every == 0 {
				// TODO: insert a database commit operation here
				fmt.Fprintln(cmd.OutOrStdout(),
					fmt.Sprintf("[%d/%d] - COMMIT", step, iters))
			}

			step += 1
			time.Sleep(time.Duration(delay) * time.Second)
		}

		// Chech the last commit
		if iters%commit_every != 0 {
			// TODO: insert a database commit operation here
			fmt.Fprintln(cmd.OutOrStdout(),
				fmt.Sprintf("[%d/%d] - COMMIT", iters, iters))
		}
	},
}

func init() {
	updateCmd.Flags().Bool("loop", false, "repeat the same operation forever")
	updateCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	updateCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	updateCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
