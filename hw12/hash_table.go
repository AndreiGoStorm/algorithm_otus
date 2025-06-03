package hw12

import "container/list"

const NumBuckets = 10

type HashTable struct {
	buckets []*list.List
}

type Item struct {
	key   string
	value interface{}
}

func NewHashTable() *HashTable {
	buckets := make([]*list.List, NumBuckets)
	return &HashTable{buckets: buckets}
}

func (ht *HashTable) Insert(key string, value interface{}) {
	index := ht.getHash(key)
	if ht.buckets[index] == nil {
		ht.buckets[index] = list.New()
		ht.buckets[index].PushFront(&Item{key, value})
		return
	}

	elem := ht.getListElement(ht.buckets[index], key)
	if elem != nil {
		elem.Value.(*Item).value = value
	} else {
		ht.buckets[index].PushFront(&Item{key, value})
	}
}

func (ht *HashTable) Remove(key string) {
	index := ht.getHash(key)
	elem := ht.getListElement(ht.buckets[index], key)
	if elem != nil {
		ht.buckets[index].Remove(elem)
	}
}

func (ht *HashTable) Get(key string) *Item {
	index := ht.getHash(key)
	elem := ht.getListElement(ht.buckets[index], key)
	if elem != nil {
		return elem.Value.(*Item)
	}
	return nil
}

func (ht *HashTable) getListElement(bucket *list.List, key string) *list.Element {
	if bucket != nil {
		for e := bucket.Front(); e != nil; e = e.Next() {
			item := e.Value.(*Item)
			if item.key == key {
				return e
			}
		}
	}
	return nil
}

func (ht *HashTable) getHash(key string) int {
	const base = 31
	hash := 0
	for _, ch := range key {
		hash = (hash*base + int(ch)) % NumBuckets
	}
	return hash
}
