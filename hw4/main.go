package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	for i := 10; i < 100_000_000; i *= 10 {
		sa := NewSingleArray()
		va := NewVectorArray(100)
		fa := NewFactorArray()
		ma := NewMatrixArray(100)
		if i <= 10000 {
			testPut(sa, i)
		}
		if i <= 100000 {
			testPut(va, i)
		}
		testPut(fa, i)
		testPut(ma, i)
		if i <= 1000000 {
			pq := NewPriorityQueue()
			testQueue(pq, i)
		}

		fmt.Println()
	}
}

func testPut(sa IArray, total int) {
	start := time.Now()
	for i := 1; i <= total; i++ {
		sa.put(i)
	}
	elapsed := time.Since(start)
	fmt.Printf("size: %d %s\n", sa.getSize(), elapsed)
}

func testQueue(pq *PriorityQueue, total int) {
	start := time.Now()
	for i := 1; i <= total; i++ {
		rnd, _ := rand.Int(rand.Reader, big.NewInt(100))
		pq.Enqueue(i, rnd)
	}
	elapsed := time.Since(start)
	fmt.Printf("queue size: %d %s\n", pq.getSize(), elapsed)
}
