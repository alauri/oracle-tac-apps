package cmd


import (
    "testing"
    "bytes"
    "runtime"
    "path"
    "strings"

    "github.com/stretchr/testify/assert"
)


func Test_Delete_No_Args(t *testing.T) {
    // Invoke the command ``delete`` with no options.

    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"-w", static, "delete"})
    rootCmd.Execute()

    expected := []string{"[1/1] - DELETE FROM test WHERE id=1",
                         "[1/1] - COMMIT"}

    splitFn := func(c rune) bool { return c == '\n' }
    assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}


func Test_Delete_Iters(t *testing.T) {
    // Invoke the command ``delete`` with the options to iterate the same
    // operation multiple times.

    _, filename, _, _ := runtime.Caller(0)
    static := path.Join(path.Dir(filename), "../static")

    actual := new(bytes.Buffer)

    rootCmd.SetOut(actual)
    rootCmd.SetErr(actual)
    rootCmd.SetArgs([]string{"-w", static, "delete", "--iters", "5", "--delay", "0.05", "--commit-every", "2"})
    rootCmd.Execute()

    expected := []string{"[1/5] - DELETE FROM test WHERE id=1",
                         "[2/5] - DELETE FROM test WHERE id=2",
                         "[2/5] - COMMIT",
                         "[3/5] - DELETE FROM test WHERE id=3",
                         "[4/5] - DELETE FROM test WHERE id=4",
                         "[4/5] - COMMIT",
                         "[5/5] - DELETE FROM test WHERE id=5",
                         "[5/5] - COMMIT"}
    splitFn := func(c rune) bool { return c == '\n' }
    assert.Equal(t, expected, strings.FieldsFunc(actual.String(), splitFn))
}
