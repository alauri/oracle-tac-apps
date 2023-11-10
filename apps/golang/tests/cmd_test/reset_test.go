/*
Copyright © 2022 Andrea Lauri <andrea.lauri86@gmail.com>

Tests for the package “reset.go“
*/
package cmd_test

import (
  "bytes"
  "path"
  "runtime"
  "strings"
  "testing"
  
  "github.com/stretchr/testify/assert"
  "github.com/DATA-DOG/go-sqlmock"

  "github.com/alauri/oracle-tac-apps/cmd"
)

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

  cmd.RootCmd.SetOut(actual)
  cmd.RootCmd.SetErr(actual)
  cmd.RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "reset"})
  cmd.RootCmd.Execute()

	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
	assert.Equal(t, 2, strings.Count(actual.String(), "TRUNCATE TABLE"))
	assert.Equal(t, 2, strings.Count(actual.String(), "ALTER TABLE"))
	assert.Contains(t, actual.String(), "All tables have been altered.")
}
