package main

type FactorArray struct {
	array []interface{}
	size  int
}

func NewFactorArray() *FactorArray {
	sa := new(FactorArray)
	sa.array = make([]interface{}, 0)
	return sa
}

func (fa *FactorArray) getSize() int {
	return fa.size
}

func (fa *FactorArray) isEmpty() bool {
	return fa.getSize() == 0
}

func (fa *FactorArray) put(T interface{}) {
	fa.resize(fa.getSize())

	fa.array[fa.size] = T
	fa.size++
}

func (fa *FactorArray) resize(size int) {
	if size < len(fa.array) {
		return
	}

	newArray := make([]interface{}, len(fa.array)*2+1)
	copy(newArray, fa.array)
	fa.array = newArray
}

func (fa *FactorArray) get(index int) interface{} {
	return fa.array[index]
}

func (fa *FactorArray) add(index int, T interface{}) {
	if index >= fa.getSize() {
		fa.put(T)
		return
	}

	fa.resize(fa.getSize() + 1)
	for i := fa.getSize() - 1; i >= 0; i-- {
		fa.array[i+1] = fa.array[i]
		if i == index {
			fa.array[i] = T
			break
		}
	}
	fa.size++
}

func (fa *FactorArray) delete(index int) interface{} {
	if index >= fa.getSize() {
		panic("delete factor: not founded index")
	}

	T := fa.get(index)
	fa.size--
	for i := index; i < fa.getSize(); i++ {
		fa.array[i] = fa.array[i+1]
	}
	fa.array[fa.getSize()] = nil

	return T
}
