package hw7

func PyramidSort(numbers []int) {
	heap(numbers)
	for i := len(numbers) - 1; i >= 0; i-- {
		swap(numbers, 0, i)
		sortHeap(numbers, 0, i)
	}
}

func heap(numbers []int) {
	for i := len(numbers)/2 - 1; i >= 0; i-- {
		sortHeap(numbers, i, len(numbers))
	}
}

func sortHeap(numbers []int, root, size int) {
	p := root
	l := 2*p + 1
	r := 2*p + 2
	if l < size && numbers[l] > numbers[p] {
		p = l
	}
	if r < size && numbers[r] > numbers[p] {
		p = r
	}
	if p != root {
		swap(numbers, p, root)
		sortHeap(numbers, p, size)
	}
}

func ShakeSort(numbers []int) {
	l := 0
	r := len(numbers) - 1
	var flag bool
	for l < r {
		flag = false
		for i := r; i > l; i-- {
			if numbers[i] < numbers[i-1] {
				flag = true
				swap(numbers, i, i-1)
			}
		}
		l++

		for i := l; i < r; i++ {
			if numbers[i] > numbers[i+1] {
				flag = true
				swap(numbers, i, i+1)
			}
		}
		r--

		if !flag {
			break
		}
	}
}

func SelectSort(numbers []int) {
	var minimum int
	for i := 0; i < len(numbers)-1; i++ {
		minimum = i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[minimum] > numbers[j] {
				minimum = j
			}
		}
		swap(numbers, i, minimum)
	}
}

func swap(numbers []int, i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
