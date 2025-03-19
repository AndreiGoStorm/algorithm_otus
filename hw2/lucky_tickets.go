package hw2

import "fmt"

const ROWS = 10

func runLuckyTickets(n int) int64 {
	data := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	x := 1
	for x <= n {
		columns := (ROWS-1)*x + 1
		matrix := buildMatrix(data, columns)
		data = calculateColumnSum(matrix, columns)
		x++
	}

	count := countLuckyTickets(data)
	fmt.Printf("Lucky Tickets for N = %d: %d\r\n", n, count)
	return count
}

func countLuckyTickets(sum []int) int64 {
	var count int64
	for i := 0; i < len(sum); i++ {
		count += int64(sum[i]) * int64(sum[i])
	}
	return count
}

func calculateColumnSum(matrix [][]int, columns int) []int {
	sum := make([]int, columns)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			sum[j] += matrix[i][j]
		}
	}
	return sum
}

func buildMatrix(data []int, columns int) [][]int {
	matrix := make([][]int, ROWS)
	for i := 0; i < ROWS; i++ {
		matrix[i] = make([]int, columns)
		for j := 0; j < len(data); j++ {
			if i+j >= columns {
				break
			}
			matrix[i][i+j] = data[j]
		}
	}
	return matrix
}
