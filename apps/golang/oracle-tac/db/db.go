/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

*/
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
	"github.com/spf13/viper"
)

func GetDatabase(d int) (*sql.DB, error) {
	cfg := viper.GetViper()

	// Database credentials
	username := cfg.GetString("driver.username")
	password := cfg.GetString("driver.password")
	dsn := fmt.Sprintf("driver.dsn%d", d)
    dsn = cfg.GetString(dsn)
    fmt.Println(dsn)

	// Connect to the database
	conn, err := sql.Open("godror",
		fmt.Sprintf(`user="%s" password="%s" connectString="%s"`,
			username,
			password,
			dsn))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func DoQuery(c *sql.DB, q string) {
	// Perform the query
	rows, err := c.Query(q)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}
