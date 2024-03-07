package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/alauri/oracle-tac-apps/internal/tests"
)

func TestIngestNoArgs(t *testing.T) {
	/* Invoke the command ``ingest`` with no options. */
	mock, tearDownDatabase := tests.MockDatabase()
	defer tearDownDatabase()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raw_tel")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "ingest"})
	RootCmd.Execute()

	assert.Contains(t, actual.String(), "(2021,'Abu Dhabi','NaT|1|Car 1|Driver 1')")
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}

func TestIngestArgs(t *testing.T) {
	/* Invoke the command ``ingest`` by iterating over the input to store
	 * read data.
	 */
	mock, tearDownDatabase := tests.MockDatabase()
	defer tearDownDatabase()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO raw_tel")
	mock.ExpectCommit()
	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	mock.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{
		"-w", static, "-d", "localhost", "ingest",
		"--iters", "5",
		"--delay", "0.05",
		"--commit-every", "5",
	})
	RootCmd.Execute()

	assert.Equal(t, 1, strings.Count(actual.String(), "NaT"))
	assert.Equal(t, 4, strings.Count(actual.String(), "0 days 00:0"))
	assert.Equal(t, 1, strings.Count(actual.String(), "COMMIT"))
	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
}
