package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/alauri/oracle-tac-apps/internal/tests"
)

func TestCleanUpNoArgs(t *testing.T) {
	/* Invoke the command ``cleanup`` with no options. */
	conn, shutdown := tests.MockDatabase()
	defer shutdown()

	conn.ExpectBegin()
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "NaT|1|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectCommit()

	conn.ExpectQuery("^SELECT SYS_CONTEXT").
		WillReturnRows(sqlmock.NewRows([]string{"uname", "host"}).
			AddRow("server1", "vm1"))

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup"})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
}

func TestCleanUpWithArgs(t *testing.T) {
	/* Invoke the command ``cleanup`` by iterating over the input to store
	 * read data.
	 */
	conn, shutdown := tests.MockDatabase()
	defer shutdown()

	conn.ExpectBegin()
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "NaT|1|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "0 days 00:01:29.103000|2|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectCommit()

	conn.ExpectBegin()
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "0 days 00:01:28.827000|3|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "0 days 00:01:29.026000|4|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectCommit()

	conn.ExpectBegin()
	conn.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "0 days 00:01:28.718000|5|Car 1|Driver 1"))
	conn.ExpectExec("INSERT INTO json_tel")
	conn.ExpectCommit()

	conn.ExpectQuery("^SELECT SYS_CONTEXT").
		WillReturnRows(sqlmock.NewRows([]string{"uname", "host"}).
			AddRow("server1", "vm1"))

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{
		"-w", static, "-d", "localhost", "cleanup",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "2",
	})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
	assert.Contains(t, actual.String(), ",'00:01:29.103000',2,'Car 1','Driver 1')")
	assert.Contains(t, actual.String(), ",'00:01:28.827000',3,'Car 1','Driver 1')")
	assert.Contains(t, actual.String(), ",'00:01:29.026000',4,'Car 1','Driver 1')")
	assert.Contains(t, actual.String(), ",'00:01:28.718000',5,'Car 1','Driver 1')")
	assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
}
