package hw7

import (
	"fmt"
	"path/filepath"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestPyramidSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			PyramidSort(numbers)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestShakeSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			ShakeSort(numbers)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestSelectSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			SelectSort(numbers)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func createTester(path string) *tester.Tester {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	test.Path = filepath.Join(test.Path, path)
	return test
}
