package gcd

import (
	"fmt"
	"math/big"
	"path/filepath"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestGCD(t *testing.T) {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			numbers := test.Separate(task.In)
			M, _ := new(big.Int).SetString(numbers[0], 10)

			N, _ := new(big.Int).SetString(numbers[1], 10)

			actual := gcd(M, N)
			require.Equal(t, task.Out, actual.String())
		})
	}
}
