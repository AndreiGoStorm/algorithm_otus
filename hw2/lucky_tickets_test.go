package hw2

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	"algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestLuckyTickets(t *testing.T) {
	dir, _ := filepath.Abs(".")
	tasks := tester.New(dir).GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			in, err := strconv.Atoi(task.In)
			require.NoError(t, err)

			out, err := strconv.ParseInt(task.Out, 10, 64)
			require.NoError(t, err)

			actual := runLuckyTickets(in)
			require.Equal(t, out, actual)
		})
	}
}
