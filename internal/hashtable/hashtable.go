package hashtable

import h "samples/internal/hash"

type HashTable struct {
	size    int
	buckets []*Bucket
}

type Bucket struct {
	key   string
	value string
}

// NewHashTable returns a new HashTable.
func NewHashTable(size int) *HashTable {
	return &HashTable{
		size:    size,
		buckets: make([]*Bucket, size),
	}
}

// add adds a new key-value pair to the HashTable.
func (ht *HashTable) Add(key string, value string) {
	index := h.Hash(key) % uint32(ht.size)
	ht.buckets[index] = &Bucket{key, value}
}

// get returns the value of the given key.
func (ht *HashTable) Get(key string) string {
	index := h.Hash(key) % uint32(ht.size)
	return ht.buckets[index].value
}

// print prints the HashTable.
func (ht *HashTable) Print() {
	for i := 0; i < ht.size; i++ {
		if ht.buckets[i] != nil {
			println(ht.buckets[i].key, ht.buckets[i].value)
		}
	}
}

// size returns the size of the HashTable.
func (ht *HashTable) Size() int {
	return ht.size
}
