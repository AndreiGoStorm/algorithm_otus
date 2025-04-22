package hw8

func HoarSort(numbers []int, l, r int) {
	left := l
	right := r
	pivot := numbers[l]
	left++
	for left <= right {
		for left <= r && numbers[left] < pivot {
			left++
		}
		for right >= l && numbers[right] > pivot {
			right--
		}
		if left < right {
			swap(numbers, left, right)
		}
		if left <= right {
			left++
			right--
		}
	}
	if l != right {
		swap(numbers, l, right)
	}
	right--
	if left < r {
		HoarSort(numbers, left, r)
	}
	if l < right {
		HoarSort(numbers, l, right)
	}
}

func QuickSort(numbers []int, l, r int) {
	if l >= r {
		return
	}

	M := split(numbers, l, r)
	QuickSort(numbers, l, M-1)
	QuickSort(numbers, M+1, r)
}

func split(numbers []int, l, r int) int {
	pivot := numbers[r]
	M := l - 1
	for i := l; i <= r; i++ {
		if pivot >= numbers[i] {
			M++
			if M != i {
				swap(numbers, M, i)
			}
		}
	}

	return M
}

func MergeSortAsc(numbers []int) {
	memory := make([]int, len(numbers))
	for step := 1; step < len(numbers); step *= 2 {
		mergeAsc(numbers, memory, step)

		for i := 0; i < len(numbers); i++ {
			numbers[i] = memory[i]
		}
	}
}

func mergeAsc(numbers, memory []int, step int) {
	index := 0
	L := 0
	M := L + step
	R := L + step*2
	length := len(numbers)

	for {
		if M >= length {
			M = length
		}

		if R >= length {
			R = length
		}

		i1 := L
		i2 := M
		for ; i1 < M && i2 < R; index++ {
			if numbers[i1] <= numbers[i2] {
				memory[index] = numbers[i1]
				i1++
			} else {
				memory[index] = numbers[i2]
				i2++
			}
		}

		for ; i1 < M; index++ {
			memory[index] = numbers[i1]
			i1++
		}

		for ; i2 < R; index++ {
			memory[index] = numbers[i2]
			i2++
		}

		L += step * 2
		M += step * 2
		R += step * 2

		if L >= length {
			break
		}
	}
}

func MergeSortDesc(numbers []int, l, r int) {
	if l >= r {
		return
	}
	m := (l + r) / 2
	MergeSortDesc(numbers, l, m)
	MergeSortDesc(numbers, m+1, r)

	mergeDesc(numbers, l, m, r)
}

func mergeDesc(numbers []int, l, m, r int) {
	index := 0
	memory := make([]int, r-l+1)

	a := l
	b := m + 1
	for ; a <= m && b <= r; index++ {
		if numbers[a] > numbers[b] {
			memory[index] = numbers[b]
			b++
		} else {
			memory[index] = numbers[a]
			a++
		}
	}

	for ; a <= m; index++ {
		memory[index] = numbers[a]
		a++
	}

	for ; b <= r; index++ {
		memory[index] = numbers[b]
		b++
	}

	for i := l; i <= r; i++ {
		numbers[i] = memory[i-l]
	}
}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
