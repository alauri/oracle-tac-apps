/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command ``ingest`` is used store new records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

import (
	"fmt"
	"time"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ingestCmd represents the ingest command
var ingestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Insert new records into the table",
	Long:  `Insert new records into the table`,
	Run: func(cmd *cobra.Command, args []string) {
		iters, _ := cmd.Flags().GetInt("iters")
		delay, _ := cmd.Flags().GetFloat64("delay")
		commit_every, _ := cmd.Flags().GetInt("commit-every")

		// Define query parameters
		table := viper.GetViper().GetString("database.tableraw")
		filename := viper.GetViper().GetString("ingest.dumpfile")

		// Open the file and check error
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			// TODO: replace panic command
			panic(err)
		}
		lines := strings.Split(string(data), "\n")

		step := 1
		for _, line := range lines[:iters] {
			// Exit condition
			if step > iters {
				break
			}

			// Prepare query with updated conditions
			query := fmt.Sprintf("INSERT INTO %s(year,track,data) VALUES(%s)", table, line)

			// Perform the query
			// db.DoQuery(conn, query)
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
	ingestCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	ingestCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	ingestCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
