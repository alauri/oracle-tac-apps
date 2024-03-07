package database

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

var Db *sql.DB

var Connect = func(username string, password string, dsn string) {
	/* Database connection function using the GodRor driver. */
	var err error

	Db, err = sql.Open("godror", fmt.Sprintf("%s/%s@%s", username, password, dsn))
	if err != nil {
		panic(err.Error())
	}
}

func Get_Db_Info() (string, error) {
	/* Retrieve database information about the system unique name and the
	system host.
	*/
	query := `SELECT
		SYS_CONTEXT('USERENV', 'DB_UNIQUE_NAME') AS DB_UNIQUE_NAME,
		SYS_CONTEXT('USERENV', 'SERVER_HOST') AS HOST
		FROM DUAL`
	results, err := Db.Query(query)
	defer results.Close()
	if err != nil {
		panic(err)
	}

	// Read all the results
	var uname string
	var host string
	for results.Next() {
		if err := results.Scan(&uname, &host); err != nil {
			panic(err)
		}
	}

	// Prepare the response
	response := fmt.Sprintf("('%s', '%s')", uname, host)
	return response, nil
}
