package model

import (
	"hash"
	"hash/fnv"
)

// BloomFilter 布隆过滤器结构体
type BloomFilter struct {
	name    string        // 布隆过滤器的标识或名称
	bits    []int         // 一个用于存储布隆过滤器位的切片，每个元素代表一个位，用于表示元素是否存在于集合中。
	hashFns []hash.Hash64 // 哈希函数切片，用于对元素进行多次哈希以获取不同的哈希值，用于设置和查询位。
	size    int           // 布隆过滤器的位数，即位数组的长度。
	k       int           // 哈希函数的个数，即对元素进行多次哈希的次数。
}

// NewBloomFilter 创建一个新的布隆过滤器
func NewBloomFilter(name string, size int, k int) *BloomFilter {
	hashFns := make([]hash.Hash64, k)
	for i := 0; i < k; i++ {
		hashFns[i] = fnv.New64a() // 使用 fnv 哈希函数创建哈希对象
	}
	return &BloomFilter{
		name:    name,
		bits:    make([]int, size), // 初始化 bits 切片，每个元素初始值为 0
		hashFns: hashFns,
		size:    size,
		k:       k,
	}
}

// BloomAdd 添加元素到布隆过滤器
func (bf *BloomFilter) BloomAdd(data string) {
	// 将字符串转换为字节切片
	dataBytes := []byte(data)

	for _, hashFn := range bf.hashFns {
		// 重置哈希函数的状态，以便对新的数据重新计算哈希值
		hashFn.Reset()
		// 将要添加的元素数据写入哈希函数
		hashFn.Write(dataBytes)
		// 计算哈希值，并取模得到对应的位在布隆过滤器中的位置
		hashValue := hashFn.Sum64() % uint64(bf.size)
		// 将对应位置的元素加 1，表示该位被设置为存在
		bf.bits[hashValue]++
	}
}

// BloomContains 判断元素是否可能存在于布隆过滤器中
func (bf *BloomFilter) BloomContains(data string) bool {
	// 将字符串转换为字节切片
	dataBytes := []byte(data)

	for _, hashFn := range bf.hashFns {
		// 重置哈希函数的状态，以便对新的数据重新计算哈希值
		hashFn.Reset()
		// 将要判断的元素数据写入哈希函数
		hashFn.Write(dataBytes)
		// 计算哈希值，并取模得到对应的位在布隆过滤器中的位置
		hashValue := hashFn.Sum64() % uint64(bf.size)
		// 如果对应位置的元素为 0，表示元素肯定不存在于集合中
		if bf.bits[hashValue] == 0 {
			return false
		}
	}
	// 如果所有的哈希函数对应位置的元素都不为 0，表示元素可能存在于集合中
	return true
}

// CollectHashValues 收集布隆过滤器中所有设置为存在的位的哈希值，并返回一个集合
func (bf *BloomFilter) CollectHashValues() []uint64 {
	hashValues := make([]uint64, 0)

	for hashValue, bit := range bf.bits {
		if bit > 0 {
			hashValues = append(hashValues, uint64(hashValue))
		}
	}

	return hashValues
}
