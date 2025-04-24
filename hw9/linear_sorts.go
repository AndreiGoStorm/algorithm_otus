package hw9

import (
	"container/list"
)

const NumBuckets = 1000

func BucketSort(numbers []int, maxValue int) {
	buckets := make([]*list.List, NumBuckets)
	for i := 0; i < len(numbers); i++ {
		m := numbers[i] * NumBuckets / (maxValue + 1)
		if buckets[m] == nil {
			buckets[m] = list.New()
			buckets[m].PushBack(numbers[i])
			continue
		}
		addToBucket(buckets[m], numbers[i])
	}
	assembleBuckets(buckets, numbers)
}

func addToBucket(bucket *list.List, num int) {
	for e := bucket.Front(); e != nil; e = e.Next() {
		if num < e.Value.(int) {
			bucket.InsertBefore(num, e)
			return
		}
	}
	bucket.PushBack(num)
}

func assembleBuckets(buckets []*list.List, numbers []int) {
	index := 0
	for i := 0; i < len(buckets); i++ {
		if buckets[i] != nil {
			for e := buckets[i].Front(); e != nil; e = e.Next() {
				numbers[index] = e.Value.(int)
				index++
			}
		}
	}
}

func CountSort(numbers []int, maxValue int) {
	T := make([]int, maxValue+1)
	for i := 0; i < len(numbers); i++ {
		T[numbers[i]]++
	}

	sum := 0
	for i := 0; i <= maxValue; i++ {
		sum += T[i]
		T[i] = sum
	}

	Z := make([]int, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		T[numbers[i]]--
		Z[T[numbers[i]]] = numbers[i]
	}

	for i := 0; i < len(numbers); i++ {
		numbers[i] = Z[i]
	}
}

func RadixSort(numbers []int, digits int) {
	base := 10
	T := make([]int, len(numbers))
	r := 1
	for d := 0; d < digits; d++ {
		radix := make([]int, base)
		for i := 0; i < len(numbers); i++ {
			dig := numbers[i] / r % base
			radix[dig]++
		}

		z := 0
		for i := 0; i < base; i++ {
			z += radix[i]
			radix[i] = z
		}

		for i := len(numbers) - 1; i >= 0; i-- {
			dig := numbers[i] / r % base
			radix[dig]--
			T[radix[dig]] = numbers[i]
		}

		copy(numbers, T)
		r *= base
	}
}
