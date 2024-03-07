/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command “ingest“ is used store new records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alauri/oracle-tac-apps/internal/database"
)

// ingestCmd represents the ingest command
var IngestCmd = &cobra.Command{
	Use:   "ingest",
	Short: "Insert new records into the table",
	Long:  `Insert new records into the table`,
	Run: func(cmd *cobra.Command, args []string) {
		iters, _ := cmd.Flags().GetInt("iters")
		delay, _ := cmd.Flags().GetFloat64("delay")
		commit_every, _ := cmd.Flags().GetInt("commit-every")
		workdir, _ := cmd.Flags().GetString("workdir")

		// Define query parameters
		table := viper.GetViper().GetString("database.tableraw")

		// Open the file and check error
		data, err := ioutil.ReadFile(path.Join(workdir, "../../../../raw.txt"))
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")
		step := 1

		tx, err := database.Db.Begin()
		if err != nil {
			panic(err)
		}

		for _, line := range lines[:iters] {
			// Exit condition
			if step > iters {
				break
			}

			// Prepare query with updated conditions
			query := fmt.Sprintf("INSERT INTO %s(year,track,data) VALUES(%s)", table, line)

			// Perform the query
			tx.Exec(query)
			cmd.Println(fmt.Sprintf("[%d/%d] - %s", step, iters, query))

			// Commit changes
			if step%commit_every == 0 {
				tx.Commit()
				cmd.Println(fmt.Sprintf("[%d/%d] - COMMIT", step, iters))
			}

			step += 1
			time.Sleep(time.Duration(delay) * time.Second)
		}

		// Chech the last commit
		if iters%commit_every != 0 {
			tx.Commit()
			cmd.Println(fmt.Sprintf("[%d/%d] - COMMIT", iters, iters))
		}
	},
}

func init() {
	/* Define ingest command's flags.
	 */
	RootCmd.AddCommand(IngestCmd)

	IngestCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	IngestCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	IngestCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
