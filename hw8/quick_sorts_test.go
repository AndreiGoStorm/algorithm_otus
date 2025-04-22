package hw8

import (
	"fmt"
	"path/filepath"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestHoarSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			HoarSort(numbers, 0, len(numbers)-1)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestQuickSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			QuickSort(numbers, 0, len(numbers)-1)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestMergeSortAsc(t *testing.T) {
	test := createTester("1.digits")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			MergeSortAsc(numbers)

			actual := test.BuildActualString(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestMergeSortDesc(t *testing.T) {
	test := createTester("1.digits")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := test.GetNumberArray(data)

			MergeSortDesc(numbers, 0, len(numbers)-1)

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
