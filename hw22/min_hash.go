package main

import (
	"math/rand"
	"time"
)

const maxUint32 = ^uint32(0)

// Хэш-функция вида: h(x) = (a*x + b) mod prime
type hashFunc struct {
	a, b  uint32
	prime uint64
}

func (h hashFunc) hash(x uint32) uint32 {
	return uint32((uint64(h.a)*uint64(x) + uint64(h.b)) % h.prime)
}

type MinHasher struct {
	hashes []hashFunc
}

func NewMinHasher(k int) *MinHasher {
	const prime = 4294967311 // большое простое число (больше maxUint32)
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	hashes := make([]hashFunc, k)
	for i := 0; i < k; i++ {
		hashes[i] = hashFunc{
			a:     uint32(rnd.Intn(prime-1) + 1),
			b:     uint32(rnd.Intn(prime)),
			prime: prime,
		}
	}
	return &MinHasher{hashes: hashes}
}

// Вычисление MinHash-подписи множества (элементы переводим в uint32)
func (mh *MinHasher) Signature(set []string) []uint32 {
	sig := make([]uint32, len(mh.hashes))
	for i := range sig {
		sig[i] = maxUint32
	}

	for _, s := range set {
		x := hashStringToUint32(s)
		for i, hf := range mh.hashes {
			hv := hf.hash(x)
			if hv < sig[i] {
				sig[i] = hv
			}
		}
	}
	return sig
}

// Сравнение двух подписей
func EstimateJaccard(sig1, sig2 []uint32) float64 {
	if len(sig1) != len(sig2) {
		panic("signature lengths don't match")
	}
	matches := 0
	for i := range sig1 {
		if sig1[i] == sig2[i] {
			matches++
		}
	}
	return float64(matches) / float64(len(sig1))
}

// Простой способ перевода строки в uint32
func hashStringToUint32(s string) uint32 {
	var h uint32 = 2166136261
	for i := 0; i < len(s); i++ {
		h = (h * 16777619) ^ uint32(s[i])
	}
	return h
}
