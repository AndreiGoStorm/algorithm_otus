package hw21

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestGCD(t *testing.T) {
	test := createTester("1.gcp")
	tasks := test.GetTasks()

	regIn := regexp.MustCompile(`^(\d+)/(\d+)\+(\d+)/(\d+)$`)
	regOut := regexp.MustCompile(`^(\d+)/(\d+)`)

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			in := regIn.FindStringSubmatch(task.In)
			a := strToInt64(in[1])
			b := strToInt64(in[2])
			c := strToInt64(in[3])
			d := strToInt64(in[4])
			m := a*d + b*c
			n := b * d

			// Act
			res := gcd(m, n)

			// Assert
			out := regOut.FindStringSubmatch(task.Out)
			require.Equal(t, strToInt64(out[1]), m/res)
			require.Equal(t, strToInt64(out[2]), n/res)
		})
	}
}

func TestGarland(t *testing.T) {
	test := createTester("2.garland")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			tree := getMatrix(task.In)

			// Act
			sum := garland(tree)

			// Assert
			require.Equal(t, strToInt64(task.Out), sum)
		})
	}
}

func TestFiveAndEighth(t *testing.T) {
	test := createTester("3.5and8")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			N := strToInt64(task.In)

			// Act
			res := fiveAndEighth(N, 1, 0, 1, 0)

			// Assert
			require.Equal(t, strToInt64(task.Out), res)
		})
	}
}

func TestBigIsland(t *testing.T) {
	test := createTester("4.big_island")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			matrix := getMatrix(task.In)

			// Act
			count := bigIsland(matrix)

			// Assert
			require.Equal(t, strToInt64(task.Out), count)
		})
	}
}

func createTester(path string) *tester.Tester {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	test.Path = filepath.Join(test.Path, path)
	return test
}

func strToInt64(s string) int64 {
	a, _ := strconv.ParseInt(s, 10, 64)
	return a
}

func getMatrix(s string) [][]int64 {
	parts := strings.Split(s, "\r\n")
	length := strToInt64(parts[0])
	tree := make([][]int64, length)
	for i := int64(0); i < length; i++ {
		tree[i] = make([]int64, 0, length)
		p := strings.Split(strings.Trim(parts[i+1], " "), " ")
		for j := 0; j < len(p); j++ {
			tree[i] = append(tree[i], strToInt64(p[j]))
		}
	}
	return tree
}
