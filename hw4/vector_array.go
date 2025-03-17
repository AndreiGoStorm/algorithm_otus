package main

type VectorArray struct {
	array  []interface{}
	vector int
	size   int
}

func NewVectorArray(vector int) *VectorArray {
	va := new(VectorArray)
	va.array = make([]interface{}, 0)
	va.vector = vector
	return va
}

func (va *VectorArray) getSize() int {
	return va.size
}

func (va *VectorArray) isEmpty() bool {
	return va.getSize() == 0
}

func (va *VectorArray) put(T interface{}) {
	va.resize(va.getSize())

	va.array[va.size] = T
	va.size++
}

func (va *VectorArray) resize(size int) {
	if size < len(va.array) {
		return
	}

	newArray := make([]interface{}, len(va.array)+va.vector)
	copy(newArray, va.array)
	va.array = newArray
}

func (va *VectorArray) get(index int) interface{} {
	return va.array[index]
}

func (va *VectorArray) add(index int, T interface{}) {
	if index >= va.getSize() {
		va.put(T)
		return
	}

	va.resize(va.getSize())

	for i := va.getSize() - 1; i >= 0; i-- {
		va.array[i+1] = va.array[i]
		if i == index {
			va.array[i] = T
			break
		}
	}
	va.size++
}

func (va *VectorArray) delete(index int) interface{} {
	if index >= va.getSize() {
		panic("delete vector: not founded index")
	}

	T := va.get(index)
	va.size--
	for i := index; i < va.getSize(); i++ {
		va.array[i] = va.array[i+1]
	}
	va.array[va.getSize()] = nil

	return T
}
