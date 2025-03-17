package main

type IArray interface {
	getSize() int
	isEmpty() bool
	put(T interface{})
	get(index int) interface{}
	add(index int, T interface{})
	delete(index int) interface{}
}
