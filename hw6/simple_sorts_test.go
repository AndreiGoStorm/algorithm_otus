package hw6

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestBubbleSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := getNumberArray(data)

			bubbleSort(numbers)

			actual := buildActual(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestInsertSort(t *testing.T) {
	test := createTester("0.random")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := getNumberArray(data)

			insertSort(numbers)

			actual := buildActual(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func TestShellSort(t *testing.T) {
	test := createTester("1.digits")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			data := test.Separate(task.In)
			numbers := getNumberArray(data)

			shellSort(numbers)

			actual := buildActual(numbers)
			require.Equal(t, task.Out, actual)
		})
	}
}

func getNumberArray(data []string) []int {
	length, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nlength: %d\n", length)

	numbers := make([]int, length)
	for key, value := range strings.Split(data[1], " ") {
		numbers[key], err = strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
	}
	return numbers
}

func buildActual(arr []int) string {
	var builder strings.Builder
	for _, value := range arr {
		builder.WriteString(strconv.Itoa(value))
		builder.WriteString(" ")
	}
	result := builder.String()
	return strings.TrimSpace(result)
}

func createTester(path string) *tester.Tester {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	test.Path = filepath.Join(test.Path, path)
	return test
}
