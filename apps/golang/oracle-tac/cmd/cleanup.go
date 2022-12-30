/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Command “update“ to change already existing records within the db.

It can repeat the same operation in loop or a defined numbers of times. It can
be possible to define a delay between one operation and the next one and also
after how many operations commit the changes.
*/
package cmd

import "fmt"
import "time"
import "strings"
import "strconv"

import "github.com/spf13/cobra"
import "github.com/spf13/viper"

import "github.com/alauri/oracle-tac-apps/oracle-tac/database"

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
		tableraw := viper.GetViper().GetString("database.tableraw")
		tablejson := viper.GetViper().GetString("database.tablejson")

		step := 1
                rowid := 1

		tx, err := database.Db.Begin()
		if err != nil {
			panic(err)
		}

		for {
			// Exit condition
			if !loop && step > iters {
				break
			}

		        query := fmt.Sprintf(`SELECT year, track, data
                                              FROM %s
                                              WHERE id=%d`, tableraw, rowid)
                        results, err := tx.Query(query)
                        defer results.Close()
                        if err != nil {
                                panic(err)
                        }
                        cmd.Println(fmt.Sprintf("[%d/%d] - %s", step, iters, query))

			// TODO: Check empty query result

                        // Read results and split data information into
                        // variables
                        var year int
                        var track string
                        var data string
                        for results.Next() {
                                if err := results.Scan(&year, &track, &data); err != nil {
                                        panic(err)
                                }
                        }
			lt, ln, team, driver := func() (string, int, string, string) {
				items := strings.Split(data, "|")
                                i, _ := strconv.Atoi(items[1])
				return strings.Replace(items[0], "0 days ", "", 1), i, items[2], items[3]
			}()

                        // Insert into the table the cleaned data
                        query = fmt.Sprintf(`INSERT INTO %s(year,track,laptime,lapnumber,team,driver) 
                                             VALUES('%d','%s','%s',%d,'%s','%s')`,
                                             tablejson, year, track, lt, ln, team, driver)
			tx.Exec(query)
                        cmd.Println(fmt.Sprintf("[%d/%d] - %s", step, iters, query))

			// Commit changes
			if step%commit_every == 0 {
				tx.Commit()
				cmd.Println(fmt.Sprintf("[%d/%d] - COMMIT", step, iters))

                                // Open a new transaction
                                if step != iters {
                                        tx, err = database.Db.Begin()
                                        if err != nil {
                                                panic(err)
                                        }
                                }
			}

			step += 1
                        rowid += 1
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
	rootCmd.AddCommand(cleanupCmd)

	cleanupCmd.Flags().Bool("loop", false, "repeat the same operation forever")
	cleanupCmd.Flags().Int("iters", 1, "repeat the same operation a given number of times")
	cleanupCmd.Flags().Float64("delay", 0.25, "time to wait before the next iteration")
	cleanupCmd.Flags().Int("commit-every", 1, "after how many operations perform a commit")
}
