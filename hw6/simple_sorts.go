package hw6

func bubbleSort(numbers []int) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := len(numbers) - 1; j > i; j-- {
			if numbers[j] < numbers[j-1] {
				swap(numbers, j, j-1)
			}
		}
	}
}

func insertSort(numbers []int) {
	var key int
	for i := 1; i < len(numbers); i++ {
		key = numbers[i]
		pos := searchBinary(numbers, key, 0, i-1)
		for j := i - 1; j >= pos; j-- {
			numbers[j+1] = numbers[j]
		}
		numbers[pos] = key
	}
}

func searchBinary(numbers []int, key, left, right int) int {
	if left > right {
		return left
	}

	mid := (left + right) / 2
	if key < numbers[mid] {
		return searchBinary(numbers, key, left, mid-1)
	}
	return searchBinary(numbers, key, mid+1, right)
}

func shellSort(numbers []int) {
	for gap := len(numbers) / 2; gap > 0; gap /= 2 {
		for j := gap; j < len(numbers); j++ {
			for i := j; i >= gap; i -= gap {
				if numbers[i-gap] <= numbers[i] {
					break
				}

				swap(numbers, i-gap, i)
			}
		}
	}
}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
