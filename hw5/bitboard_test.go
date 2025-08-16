package hw5

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	tester "algorithm_otus/pkg"
	"github.com/stretchr/testify/require"
)

func TestKingMovesTest(t *testing.T) {
	test := createTester("1.King")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			b := NewBitboard()
			b.SetPoint(strToUInt(task.In))

			// Act
			b.KingMoves()

			// Assert
			count := b.CountBits()
			expectedCount, expectedBitboard := getOut(task.Out)
			require.Equal(t, expectedCount, count)
			require.Equal(t, expectedBitboard, b.bitboard)
		})
	}
}

func TestKnightMovesTest(t *testing.T) {
	test := createTester("2.Knight")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			b := NewBitboard()
			b.SetPoint(strToUInt(task.In))

			// Act
			b.KnightMoves()

			// Assert
			count := b.CountBits()
			expectedCount, expectedBitboard := getOut(task.Out)
			require.Equal(t, expectedCount, count)
			require.Equal(t, expectedBitboard, b.bitboard)
		})
	}
}

func TestRookMovesTest(t *testing.T) {
	test := createTester("3.Rook")
	tasks := test.GetTasks()

	for _, task := range tasks {
		t.Run(fmt.Sprintf("Test for file: %s", task.TestFile), func(t *testing.T) {
			// Arrange
			b := NewBitboard()
			b.SetPoint(strToUInt(task.In))

			// Act
			b.RookMoves()

			// Assert
			count := b.CountBits()
			expectedCount, expectedBitboard := getOut(task.Out)
			require.Equal(t, expectedCount, count)
			require.Equal(t, expectedBitboard, b.bitboard)
		})
	}
}

func createTester(path string) *tester.Tester {
	dir, _ := filepath.Abs(".")
	test := tester.New(dir)
	test.Path = filepath.Join(test.Path, path)
	return test
}

func strToUInt(s string) uint {
	return uint(strToUInt64(s))
}

func strToUInt64(s string) uint64 {
	a, _ := strconv.ParseUint(s, 10, 64)
	return a
}

func getOut(s string) (uint, uint64) {
	parts := strings.Split(s, "\r\n")
	count := strToUInt(parts[0])
	bitboard := strToUInt64(parts[1])

	return count, bitboard
}
