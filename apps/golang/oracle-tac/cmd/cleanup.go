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

	// "github.com/alauri/oracle-tac-apps/oracle-tac/db"
)

// cleanupCmd represents the update command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Update records into the table.",
	Long:  `Update records into the table.`,
	Run: func(cmd *cobra.Command, args []string) {
		// dsn, _ := cmd.Flags().GetInt("dsn")
		loop, _ := cmd.Flags().GetBool("loop")
		iters, _ := cmd.Flags().GetInt("iters")
		delay, _ := cmd.Flags().GetFloat64("delay")
		commit_every, _ := cmd.Flags().GetInt("commit-every")

                // Get the first ID to read from the table raw
                tail := viper.GetViper().GetInt("cleanup.tail")
                tableraw := viper.GetViper().GetString("database.tableraw")
                // tablejson := viper.GetViper().GetString("database.tablejson")

		// Retrieve a fresh Database connection
		// conn, err := db.GetDatabase(dsn)
		// if err != nil {
		// 	panic(err)
		// }

		if loop {
			iters = 0
		}
		step := 1
		for {
			// Exit condition
			if !loop && step > iters {
				break
			}

                        query := fmt.Sprintf(`SELECT * 
                                              FROM %s 
                                              WHERE id=%d`, tableraw, tail)
                        // TODO: db.DoQuery(conn, query)
			fmt.Fprintln(cmd.OutOrStdout(),
				fmt.Sprintf("[%d/%d] - %s", step, iters, query))
                                
                        // TODO: Check empty query result

                        // TODO: Unpack result into multiple variables
                        // TODO: Perform query with unpacked variables

			// Commit changes
			if step%commit_every == 0 {
				// TODO: Commit operation
				fmt.Fprintln(cmd.OutOrStdout(),
					fmt.Sprintf("[%d/%d] - COMMIT", step, iters))
			}

                        step += 1
                        tail += 1
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
	cleanupCmd.Flags().Bool("loop", false, "repeat the same operation forever")
	cleanupCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	cleanupCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	cleanupCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
