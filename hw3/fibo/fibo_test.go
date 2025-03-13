package fibo

import (
	"fmt"
	"path/filepath"
	"strconv"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestFiboGoldenRatio(t *testing.T) {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			k, err := strconv.ParseInt(task.In, 10, 64)
			require.NoError(t, err)

			actual := fiboGoldenRatio(k)
			require.Equal(t, task.Out, actual.String())
		})
	}
}

func TestFiboMatrix(t *testing.T) {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			k, err := strconv.ParseInt(task.In, 10, 64)
			require.NoError(t, err)

			actual := fiboMatrix(k)
			require.Equal(t, task.Out, actual.String())
		})
	}
}
