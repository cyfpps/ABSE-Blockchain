package MODEL

import (
	"crypto/sha256"
)

const L = 256 // 布隆过滤器算法数组大小

type BloomFilter struct {
	bits []int
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		bits: make([]int, L),
	}
}

func (bf *BloomFilter) Add(data string) {
	hashes := bf.hash(data)
	for _, h := range hashes {
		bf.bits[h]++
	}
}

func (bf *BloomFilter) Check(index int) bool {
	if index < 0 || index >= L {
		return false
	}
	return bf.bits[index] > 0
}
func (bf *BloomFilter) Contains(data string) bool {
	hashes := bf.hash(data)
	for _, h := range hashes {
		if bf.bits[h] == 0 {
			return false
		}
	}
	return true
}

func (bf *BloomFilter) hash(data string) [3]int {
	hashed := sha256.Sum256([]byte(data))
	return [3]int{
		int(hashed[0]) % L,
		int(hashed[1]) % L,
		int(hashed[2]) % L,
	}
}

func (bf *BloomFilter) Remove(attr string) {
	hashes := bf.hash(attr)
	for _, h := range hashes {
		if bf.bits[h] > 0 {
			bf.bits[h]--
		}
	}
}
