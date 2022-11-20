/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “reset.go“
*/
package cmd

import "bytes"
import "path"
import "runtime"
import "strings"
import "testing"

import "github.com/stretchr/testify/assert"
import "github.com/DATA-DOG/go-sqlmock"

func Test_Reset_No_Args(t *testing.T) {
	/* Invoke the command ``reset`` with no options. */
	mock, tearDownDatabase := setUpDatabase(t)
	defer tearDownDatabase(t)

	mock.ExpectQuery("TRUNCATE TABLE raw_tel")
	mock.ExpectQuery("ALTER TABLE raw_tel")
	mock.ExpectQuery("TRUNCATE TABLE json_tel")
	mock.ExpectQuery("ALTER TABLE json_tel")
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../static")

	actual := new(bytes.Buffer)

	rootCmd.SetOut(actual)
	rootCmd.SetErr(actual)
	rootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "reset"})
	rootCmd.Execute()

	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
	assert.Equal(t, 2, strings.Count(actual.String(), "TRUNCATE TABLE"))
	assert.Equal(t, 2, strings.Count(actual.String(), "ALTER TABLE"))
	assert.Contains(t, actual.String(), "All tables have been altered.")
}
