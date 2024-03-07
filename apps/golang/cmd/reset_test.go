package cmd

import (
	"bytes"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/alauri/oracle-tac-apps/internal/tests"
)

// Test command reset
func TestResetNoArgs(t *testing.T) {
	conn, shutdown := tests.MockDatabase()
	defer shutdown()

	conn.ExpectQuery("TRUNCATE TABLE raw_tel")
	conn.ExpectQuery("ALTER TABLE raw_tel")
	conn.ExpectQuery("TRUNCATE TABLE json_tel")
	conn.ExpectQuery("ALTER TABLE json_tel")

	rows := sqlmock.NewRows([]string{"uname", "host"}).AddRow("server1", "vm1")
	conn.ExpectQuery("^SELECT SYS_CONTEXT").WillReturnRows(rows)

	_, filename, _, _ := runtime.Caller(0)
	static := path.Join(path.Dir(filename), "../internal/tests/static")

	actual := new(bytes.Buffer)

	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)
	RootCmd.SetArgs([]string{"-w", static, "-d", "localhost", "reset"})
	RootCmd.Execute()

	assert.Equal(t, 2, strings.Count(actual.String(), "('server1', 'vm1')"))
	assert.Equal(t, 2, strings.Count(actual.String(), "TRUNCATE TABLE"))
	assert.Equal(t, 2, strings.Count(actual.String(), "ALTER TABLE"))
	assert.Contains(t, actual.String(), "All tables have been altered.")
}
