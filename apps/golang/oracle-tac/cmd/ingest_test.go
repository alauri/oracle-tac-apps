/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “ingest.go“.
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "strings"
import "testing"

import "github.com/stretchr/testify/assert"
import "github.com/DATA-DOG/go-sqlmock"

func Test_Ingest_No_Args(t *testing.T) {
	/* Invoke the command ``ingest`` with no options. */
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raw_tel")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "ingest"})
	rootCmd.Execute()

	assert.Contains(t, actual.String(), "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')")
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func Test_Ingest_Args(t *testing.T) {
	/* Invoke the command ``ingest`` by iterating over the input to store
	 * read data.
	 */
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raw_tel")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "ingest",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "5"})
	rootCmd.Execute()

	assert.Equal(t, 1, strings.Count(actual.String(), "NaT"))
	assert.Equal(t, 4, strings.Count(actual.String(), "0 days 00:0"))
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
