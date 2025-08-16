package main

import "fmt"

func main() {
	hll := NewHLL(14) // 2^14 = 16384 регистров

	// Добавим 100000 уникальных значений
	for i := 0; i < 100000; i++ {
		hll.Add(fmt.Sprintf("user-%d", i))
	}

	// Добавим 50000 повторов
	for i := 0; i < 50000; i++ {
		hll.Add(fmt.Sprintf("user-%d", i))
	}

	fmt.Println("Примерное количество уникальных:", hll.Count())

	// MinHash
	setA := []string{"apple", "banana", "cherry"}
	setB := []string{"banana", "cherry", "date"}

	mh := NewMinHasher(100)

	sigA := mh.Signature(setA)
	sigB := mh.Signature(setB)

	similarity := EstimateJaccard(sigA, sigB)
	fmt.Printf("Оценочная Jaccard-похожесть: %.2f\n", similarity)
}
