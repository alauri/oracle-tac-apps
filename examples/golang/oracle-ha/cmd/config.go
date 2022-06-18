/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package cmd

import (
	"fmt"

    "github.com/alauri/oracle-ha-apps/oracle-ha/cfg"
	"github.com/spf13/cobra"
)


var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure the application.",
	Long: `configure the application.`,
	Run: func(cmd *cobra.Command, args []string) {
        // workdirFlag, _ := cmd.Flags().GetString("workdir")
        // config := cfg.ReadTOML(workdirFlag)
        
        infoFlag, _ := cmd.Flags().GetBool("info")
        if infoFlag {
            fmt.Fprintln(cmd.OutOrStdout(), "[+] - Current configuration: ")
        } else {
           cmd.Help()
        }
	},
}


var driverCmd = &cobra.Command{
	Use:   "driver",
	Short: "update section 'driver'",
	Long: `update section 'driver'`,
	Run: func(cmd *cobra.Command, args []string) {
        workdirFlag, _ := cmd.Flags().GetString("workdir")
        username, _ := cmd.Flags().GetString("username")
        password, _ := cmd.Flags().GetString("password")
        
        // Update driver's username if needed
        if username != "" {
            cfg.UpdateTOML(workdirFlag, "driver.username", username)
        }
        
        // Update driver's password if needed
        if password != "" {
            cfg.UpdateTOML(workdirFlag, "driver.password", password)
        }
        
        fmt.Fprintln(cmd.OutOrStdout(), "[+] - Configuration updated")
	},
}


var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "update section 'database'",
	Long: `update section 'database'`,
	Run: func(cmd *cobra.Command, args []string) {
        workdirFlag, _ := cmd.Flags().GetString("workdir")
        table, _ := cmd.Flags().GetString("table")
        
        // Update database's table if needed
        if table != "" {
            cfg.UpdateTOML(workdirFlag, "database.table", table)
        }
	},
}


func init() {
    // Config flags
    configCmd.Flags().Bool("info", false, "print current stored configurations")

    // Driver flags
    driverCmd.Flags().String("username", "", "update the driver's username")
    driverCmd.Flags().String("password", "", "update the driver's password")

    // Database flags
    databaseCmd.Flags().String("table", "", "update the database's table")

    // Register sub-commands
    configCmd.AddCommand(driverCmd)
    configCmd.AddCommand(databaseCmd)
}
