/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “cleanup.go“
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "strings"
import "testing"

import "github.com/stretchr/testify/assert"
import "github.com/DATA-DOG/go-sqlmock"

func Test_Cleanup_No_Args(t *testing.T) {
        /* Invoke the command ``cleanup`` with no options. */
        mock, tearDownDatabase := setUpDatabase(t)
        defer tearDownDatabase(t)

        mock.ExpectBegin()
        mock.ExpectQuery("^SELECT year, track, data FROM ").
        	WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
        		AddRow(2021, "Abu Dhabi", "NaT|1|Car 1|Driver 1"))
        mock.ExpectExec("INSERT INTO json_tel")
        mock.ExpectCommit()
        
        mock.ExpectQuery("^SELECT SYS_CONTEXT").
        	WillReturnRows(sqlmock.NewRows([]string{"uname", "host"}).
        		AddRow("server1", "vm1"))

        _, filename, _, _ := runtime.Caller(0)
        static := path.Join(path.Dir(filename), "../static")

        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup"})
        rootCmd.Execute()

        assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
        assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
}

func Test_Cleanup_With_Args(t *testing.T) {
        /* Invoke the command ``cleanup`` by iterating over the input to store
         * read data.
         */
        mock, tearDownDatabase := setUpDatabase(t)
        defer tearDownDatabase(t)

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "NaT|1|Car 1|Driver 1"))
	mock.ExpectExec("INSERT INTO json_tel")
	mock.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
			AddRow(2021, "Abu Dhabi", "0 days 00:01:29.103000|2|Car 1|Driver 1"))
	mock.ExpectExec("INSERT INTO json_tel")
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
       			AddRow(2021, "Abu Dhabi", "0 days 00:01:28.827000|3|Car 1|Driver 1"))
	mock.ExpectExec("INSERT INTO json_tel")
	mock.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
       			AddRow(2021, "Abu Dhabi", "0 days 00:01:29.026000|4|Car 1|Driver 1"))
	mock.ExpectExec("INSERT INTO json_tel")
	mock.ExpectCommit()

	mock.ExpectBegin()
	mock.ExpectQuery("^SELECT year, track, data FROM ").
		WillReturnRows(sqlmock.NewRows([]string{"year", "track", "data"}).
       			AddRow(2021, "Abu Dhabi", "0 days 00:01:28.718000|5|Car 1|Driver 1"))
	mock.ExpectExec("INSERT INTO json_tel")
	mock.ExpectCommit()
	
	mock.ExpectQuery("^SELECT SYS_CONTEXT").
		WillReturnRows(sqlmock.NewRows([]string{"uname", "host"}).
			AddRow("server1", "vm1"))

        _, filename, _, _ := runtime.Caller(0)
        static := path.Join(path.Dir(filename), "../static")

        actual := new(bytes.Buffer)

        rootCmd.SetOut(actual)
        rootCmd.SetErr(actual)
        rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "cleanup",
        	"--iters", "5",
        	"--delay", "0.05",
        	"--commit-every", "2"})
        rootCmd.Execute()

	assert.Contains(t, actual.String(), ",'NaT',1,'Car 1','Driver 1')")
    	assert.Contains(t, actual.String(), ",'00:01:29.103000',2,'Car 1','Driver 1')")
    	assert.Contains(t, actual.String(), ",'00:01:28.827000',3,'Car 1','Driver 1')")
    	assert.Contains(t, actual.String(), ",'00:01:29.026000',4,'Car 1','Driver 1')")
    	assert.Contains(t, actual.String(), ",'00:01:28.718000',5,'Car 1','Driver 1')")
        assert.Equal(t, 3, strings.Count(actual.String(), "COMMIT"))
}
