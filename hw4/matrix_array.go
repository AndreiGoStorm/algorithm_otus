package main

type MatrixArray struct {
	array  IArray
	vector int
	size   int
}

func NewMatrixArray(vector int) *MatrixArray {
	ma := new(MatrixArray)
	ma.vector = vector
	ma.array = NewSingleArray()
	return ma
}

func (ma *MatrixArray) getSize() int {
	return ma.size
}

func (ma *MatrixArray) isEmpty() bool {
	return ma.getSize() == 0
}

func (ma *MatrixArray) put(T interface{}) {
	if ma.size == ma.array.getSize()*ma.vector {
		ma.array.put(NewVectorArray(ma.vector))
	}

	value := ma.array.get(ma.size / ma.vector)
	vectorArray := value.(*VectorArray)
	vectorArray.put(T)
	ma.size++
}

func (ma *MatrixArray) get(index int) interface{} {
	vectorArray := ma.getVectorArray(index)
	return vectorArray.get(index % ma.vector)
}

func (ma *MatrixArray) add(index int, T interface{}) {
	vectorArray := ma.getVectorArray(index)

	position := index % ma.vector
	if vectorArray.getSize() < ma.vector {
		vectorArray.add(position, T)
		return
	}

	last := vectorArray.delete(ma.vector - 1)
	vectorArray.add(position, T)
	ma.add((index-position)+ma.vector, last)
}

func (ma *MatrixArray) delete(index int) interface{} {
	vectorArray := ma.getVectorArray(index)

	position := index % ma.vector
	T := vectorArray.delete(position)

	ma.removeFirst(vectorArray, (index-position)+ma.vector)

	return T
}

func (ma *MatrixArray) removeFirst(vectorArray *VectorArray, indexNext int) {
	if indexNext < ma.size {
		vectorArrayNext := ma.getVectorArray(indexNext)

		first := vectorArrayNext.delete(0)
		vectorArray.put(first)
		ma.removeFirst(vectorArrayNext, indexNext+ma.vector)
	}
}

func (ma *MatrixArray) getVectorArray(index int) *VectorArray {
	value := ma.array.get(index / ma.vector)
	return value.(*VectorArray)
}
