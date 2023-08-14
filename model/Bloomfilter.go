package model

import (
	"hash"
	"hash/fnv"
)

// BloomFilter 布隆过滤器结构体
type BloomFilter struct {
	bits    []bool
	hashFns []hash.Hash64
	size    int
	k       int
}

// NewBloomFilter 创建一个新的布隆过滤器
func NewBloomFilter(size int, k int) *BloomFilter {
	hashFns := make([]hash.Hash64, k)
	for i := 0; i < k; i++ {
		hashFns[i] = fnv.New64a()
	}
	return &BloomFilter{
		bits:    make([]bool, size),
		hashFns: hashFns,
		size:    size,
		k:       k,
	}
}

// Add 添加元素到布隆过滤器
func (bf *BloomFilter) Add(data []byte) {
	for _, hashFn := range bf.hashFns {
		hashFn.Reset()
		hashFn.Write(data)
		hashValue := hashFn.Sum64() % uint64(bf.size)
		bf.bits[hashValue] = true
	}
}

// Contains 判断元素是否可能存在于布隆过滤器中
func (bf *BloomFilter) Contains(data []byte) bool {
	for _, hashFn := range bf.hashFns {
		hashFn.Reset()
		hashFn.Write(data)
		hashValue := hashFn.Sum64() % uint64(bf.size)
		if !bf.bits[hashValue] {
			return false
		}
	}
	return true
}
