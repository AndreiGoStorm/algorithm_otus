package prime

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestPrimes(t *testing.T) {
	dir, _ := filepath.Abs(".")
	tasks := tester.New(dir).GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			in, err := strconv.ParseInt(task.In, 10, 64)
			require.NoError(t, err)

			out, err := strconv.ParseInt(task.Out, 10, 64)
			require.NoError(t, err)

			actual := Primes(in)
			require.Equal(t, out, actual)
		})
	}
}

func TestPrimesWithMemory(t *testing.T) {
	dir, _ := filepath.Abs(".")
	tasks := tester.New(dir).GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			in, err := strconv.ParseInt(task.In, 10, 64)
			require.NoError(t, err)

			out, err := strconv.ParseInt(task.Out, 10, 64)
			require.NoError(t, err)

			actual := PrimesWithMemory(in)
			require.Equal(t, out, actual)
		})
	}
}

func TestEratosphen(t *testing.T) {
	dir, _ := filepath.Abs(".")
	tasks := tester.New(dir).GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			in, err := strconv.ParseInt(task.In, 10, 64)
			require.NoError(t, err)

			out, err := strconv.ParseInt(task.Out, 10, 64)
			require.NoError(t, err)

			actual := Eratosphen(in)
			require.Equal(t, out, actual)
		})
	}
}
