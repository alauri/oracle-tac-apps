/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command ``insert`` is used store new records within the db.

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
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert new records into the table",
	Long: `Insert new records into the table`,
	Run: func(cmd *cobra.Command, args []string) {
        loop, _ := cmd.Flags().GetBool("loop")
        iters, _ := cmd.Flags().GetInt("iters")
        delay, _ := cmd.Flags().GetFloat64("delay")
        commit_every, _ := cmd.Flags().GetInt("commit-every")

        // Define query parameters
        table := viper.GetViper().GetString("database.table")
        conditions := Conds{0}

        if loop { iters = 0 }
        step := 1
        for {
            // Exit condition
            if !loop && step > iters {
                break
            }

            // Prepare query with updated conditions
            pairs := Conds{conditions.id + step}
            query := fmt.Sprintf("INSERT INTO %s(id) VALUES(%d)", table, pairs.id)

            // TODO: replace with call to cx_Oracle driver
            fmt.Fprintln(cmd.OutOrStdout(),
                         fmt.Sprintf("[%d/%d] - %s", step, iters, query))

            // Commit changes
            if step % commit_every == 0 {
                // TODO: insert a database commit operation here
                fmt.Fprintln(cmd.OutOrStdout(),
                             fmt.Sprintf("[%d/%d] - COMMIT", step, iters))
            }

            step += 1
            time.Sleep(time.Duration(delay) * time.Millisecond)
        }

        // Chech the last commit
        if iters % commit_every != 0 {
            // TODO: insert a database commit operation here
            fmt.Fprintln(cmd.OutOrStdout(),
                         fmt.Sprintf("[%d/%d] - COMMIT", iters, iters))
        }
	},
}

func init() {
    insertCmd.Flags().Bool("loop", false, "repeat the same operation forever")
    insertCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
    insertCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
    insertCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
