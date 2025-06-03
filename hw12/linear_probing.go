package hw12

type LinearProbing struct {
	size     int
	probing  int
	hashes   []*HashItem
	capacity int
}

type HashItem struct {
	key     int
	deleted bool
}

func NewLinearProbing(capacity int) *LinearProbing {
	return &LinearProbing{hashes: make([]*HashItem, capacity), capacity: capacity}
}

func (lb *LinearProbing) getHash(key, i int) int {
	return (key + i) % lb.capacity
}

func (lb *LinearProbing) Insert(key int) {
	for i, h := range lb.hashes {
		lb.probing++
		if h != nil && h.deleted {
			continue
		}

		hash := lb.getHash(key, i)
		if lb.hashes[hash] == nil {
			lb.hashes[hash] = &HashItem{key, false}
			lb.size++
			if lb.LoadFactor() > 65 {
				lb.rehash()
			}
			return
		}
	}

	lb.rehash()
	lb.Insert(key)
}

func (lb *LinearProbing) Get(key int) *HashItem {
	for _, h := range lb.hashes {
		if h == nil || h.deleted {
			continue
		}
		if key == h.key {
			return h
		}
	}
	return nil
}

func (lb *LinearProbing) Remove(key int) {
	h := lb.Get(key)
	if h != nil {
		lb.size--
		h.deleted = true
	}
}

func (lb *LinearProbing) rehash() {
	hashes := append([]*HashItem(nil), lb.hashes...)
	lb.capacity = (lb.capacity * 2) + 1
	lb.size = 0
	lb.hashes = make([]*HashItem, lb.capacity)
	for _, h := range hashes {
		if h == nil || h.deleted {
			continue
		}
		lb.Insert(h.key)
	}
}

func (lb *LinearProbing) LoadFactor() int {
	return int(float64(100) * float64(lb.size) / float64(lb.capacity))
}
