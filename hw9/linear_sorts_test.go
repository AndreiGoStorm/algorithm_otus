package hw9

import (
	"fmt"
	"path/filepath"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestBucketSort(t *testing.T) {
	bin := createBinaryTester("0_999")
	count := bin.CountTestFiles()

	for i := 0; i < count; i++ {
		numbers := bin.ReadNumbers(i)
		t.Run(fmt.Sprintf("Test for bucket sort len: %d", len(numbers)), func(t *testing.T) {
			BucketSort(numbers, 999)

			actual := isSortedArray(numbers)
			require.Equal(t, true, actual)
		})
	}
}

func TestCountSort(t *testing.T) {
	bin := createBinaryTester("0_65535")
	count := bin.CountTestFiles()

	for i := 0; i < count; i++ {
		numbers := bin.ReadNumbers(i)
		t.Run(fmt.Sprintf("Test for count sort len: %d", len(numbers)), func(t *testing.T) {
			CountSort(numbers, 65535)

			actual := isSortedArray(numbers)
			require.Equal(t, true, actual)
		})
	}
}

func TestRadixSort(t *testing.T) {
	bin := createBinaryTester("0_999")
	count := bin.CountTestFiles()

	for i := 0; i < count; i++ {
		numbers := bin.ReadNumbers(i)
		t.Run(fmt.Sprintf("Test for radix sort len: %d", len(numbers)), func(t *testing.T) {
			RadixSort(numbers, 3)

			actual := isSortedArray(numbers)
			require.Equal(t, true, actual)
		})
	}
}

func createBinaryTester(path string) *tester.BinaryTester {
	dir, _ := filepath.Abs(".")
	test := tester.NewBinary(dir)
	test.Path = filepath.Join(test.Path, path)
	return test
}

func isSortedArray(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return false
		}
	}
	return true
}
