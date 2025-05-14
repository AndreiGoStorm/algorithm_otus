package tester

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

type BinaryTester struct {
	Path string
}

func NewBinary(path string) *BinaryTester {
	return &BinaryTester{Path: filepath.Join(path, "testdata")}
}

func (bt *BinaryTester) CountTestFiles() int {
	entries, err := os.ReadDir(bt.Path)
	if err != nil {
		return 0
	}
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			count++
		}
	}
	return count
}

func (bt *BinaryTester) WriteNumbers(max int) {
	index := 0
	for i := 100; i <= 1_000_000_000; i *= 10 {
		name := fmt.Sprintf("test.%d.bin", index)
		file, err := os.Create(filepath.Join(bt.Path, name))
		if err != nil {
			break
		}
		defer file.Close()

		var num uint16
		for j := 0; j < i; j++ {
			num = uint16(rand.Intn(max)) // 65536 [0, 65535], 1000 [0,999]
			err = binary.Write(file, binary.LittleEndian, num)
			if err != nil {
				break
			}
		}
		index++
	}
}

func (bt *BinaryTester) ReadNumbers(num int) []int {
	name := fmt.Sprintf("test.%d.bin", num)
	file, err := os.Open(filepath.Join(bt.Path, name))
	if err != nil {
		return nil
	}
	defer file.Close()

	pr := 100
	for i := 0; i < num; i++ {
		pr *= 10
	}

	numbers := make([]int, 0, pr)
	for {
		var num uint16
		err = binary.Read(file, binary.LittleEndian, &num)
		if err != nil {
			break
		}
		numbers = append(numbers, int(num))
	}
	return numbers
}
