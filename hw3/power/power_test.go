package power

import (
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

const delta = 0.000001

func TestPowerN(t *testing.T) {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			numbers := test.Separate(task.In)
			A, err := strconv.ParseFloat(numbers[0], 64)
			require.NoError(t, err)
			N, err := strconv.ParseInt(numbers[1], 10, 64)
			require.NoError(t, err)

			out, err := strconv.ParseFloat(task.Out, 64)
			require.NoError(t, err)

			actual := powerN(A, N)
			require.InDelta(t, out, actual, delta)
		})
	}
}

func TestPowerNBig(t *testing.T) {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			numbers := test.Separate(task.In)
			A, _ := new(big.Float).SetString(numbers[0])

			N, err := strconv.ParseInt(numbers[1], 10, 64)
			require.NoError(t, err)

			out, _ := new(big.Float).SetString(task.Out)

			actual := NBigFloat(A, N)
			require.True(t, floatsEqualWithDelta(out, actual))
		})
	}
}

func floatsEqualWithDelta(a, b *big.Float) bool {
	de := new(big.Float).SetFloat64(delta)
	diff := new(big.Float).Sub(a, b)

	if diff.Sign() < 0 {
		diff.Neg(diff)
	}

	return diff.Cmp(de) <= 0
}
