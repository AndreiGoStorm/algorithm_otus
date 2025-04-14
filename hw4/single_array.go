package main

type SingleArray struct {
	array []interface{}
}

func NewSingleArray() *SingleArray {
	sa := new(SingleArray)
	sa.array = make([]interface{}, 0)
	return sa
}

func (sa *SingleArray) getSize() int {
	return len(sa.array)
}

func (sa *SingleArray) isEmpty() bool {
	return sa.getSize() == 0
}

func (sa *SingleArray) put(T interface{}) {
	sa.resize()
	sa.array[sa.getSize()-1] = T
}

func (sa *SingleArray) resize() {
	newArray := make([]interface{}, len(sa.array)+1)
	for i := 0; i < sa.getSize(); i++ {
		newArray[i] = sa.array[i]
	}
	sa.array = newArray
}

func (sa *SingleArray) get(index int) interface{} {
	return sa.array[index]
}

func (sa *SingleArray) add(index int, T interface{}) {
	if index >= sa.getSize() {
		sa.put(T)
		return
	}

	newArray := make([]interface{}, len(sa.array)+1)
	for i := 0; i <= sa.getSize(); i++ {
		if i < index {
			newArray[i] = sa.array[i]
			continue
		}

		if i > index {
			newArray[i] = sa.array[i-1]
		} else {
			newArray[i] = T
		}
	}
	sa.array = newArray
}

func (sa *SingleArray) delete(index int) interface{} {
	if index >= sa.getSize() {
		panic("delete single: not founded index")
	}

	newArray := make([]interface{}, len(sa.array)-1)
	for i := 0; i < sa.getSize()-1; i++ {
		if i < index {
			newArray[i] = sa.array[i]
		} else {
			newArray[i] = sa.array[i+1]
		}
	}
	T := sa.get(index)
	sa.array = newArray
	return T
}
