package main

import (
	"hash/fnv"
	"math"
	"math/bits"
)

type HyperLogLog struct {
	precision    uint8
	numRegisters uint32
	registers    []uint8
	alpha        float64
}

func NewHLL(precision uint8) *HyperLogLog {
	numRegisters := uint32(1) << precision
	registers := make([]uint8, numRegisters)
	hll := &HyperLogLog{
		precision:    precision,
		numRegisters: numRegisters,
		registers:    registers,
	}
	hll.alpha = hll.getAlpha(numRegisters)
	return hll
}

func (hll *HyperLogLog) Add(item string) {
	hash := hll.hash64(item)

	// старшие precision битов = индекс регистра
	idx := hash >> (64 - hll.precision)

	// оставшиеся биты — считаем количество ведущих нулей + 1
	remaining := (hash << hll.precision) | (1 << (hll.precision - 1))
	zeros := bits.LeadingZeros64(remaining) + 1

	// обновляем регистр, если новое значение больше
	if uint8(zeros) > hll.registers[idx] {
		hll.registers[idx] = uint8(zeros)
	}
}

func (hll *HyperLogLog) Count() uint64 {
	sum := 0.0
	zeroCount := 0

	for _, reg := range hll.registers {
		sum += 1.0 / math.Pow(2.0, float64(reg))
		if reg == 0 {
			zeroCount++
		}
	}

	estimate := hll.alpha * float64(hll.numRegisters*hll.numRegisters) / sum
	// исправление малых оценок (Linear Counting)
	if estimate <= 2.5*float64(hll.numRegisters) && zeroCount > 0 {
		estimate = float64(hll.numRegisters) * math.Log(float64(hll.numRegisters)/float64(zeroCount))
	}

	return uint64(estimate + 0.5)
}

func (hll *HyperLogLog) hash64(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func (hll *HyperLogLog) getAlpha(m uint32) float64 {
	switch m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1.0 + 1.079/float64(m))
	}
}
