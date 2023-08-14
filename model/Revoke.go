package model

import (
	"math/big"
)

// RevokeAttributes 执行属性撤销操作
func RevokeAttributes(attributes []string, globalParams *GlobalParams, bloomFilter *BloomFilter) (map[string]*big.Int, map[string]*big.Int, map[string]*big.Int, error) {
	overlineUK := make(map[string]*big.Int)
	overlineSI := make(map[string]*big.Int)
	overlineCY := make(map[string]*big.Int)

	for _, attr := range attributes {
		// 生成随机数 Lk_i
		Lk := generateRandomElement(globalParams.G)
		LkInv := new(big.Int).ModInverse(Lk, globalParams.G)

		// 更新 UK_i
		overlineUK[attr] = new(pairing).Mul(globalParams.UK[attr], new(pairing).Pair(globalParams.G, globalParams.GAlpha).Exp(LkInv, globalParams.G))

		// 更新 SI_i
		overlineSI[attr] = new(pairing).Div(globalParams.SI[attr], new(pairing).Pair(globalParams.G, globalParams.GAlpha).Exp(Lk, globalParams.G))

		// 更新 CY
		overlineCY[attr] = new(pairing).Mul(globalParams.CY[attr], new(pairing).Pair(globalParams.G, globalParams.GAlpha).Exp(Lk, globalParams.G))

		// 更新布隆过滤器
		Ai := new(big.Int).SetBytes(hashToBytes(attr, globalParams.H1)).Mod(globalParams.L)
		Bi := new(big.Int).SetBytes(hashToBytes(attr, globalParams.H2)).Mod(globalParams.L)
		Ci := new(big.Int).SetBytes(hashToBytes(attr, globalParams.H3)).Mod(globalParams.L)

		bloomFilter.Decrement(Ai.Int64())
		bloomFilter.Decrement(Bi.Int64())
		bloomFilter.Decrement(Ci.Int64())
	}

	return overlineUK, overlineSI, overlineCY, nil
}

// hashToBytes performs a hash and returns bytes of the result
func hashToBytes(input string, hashFunc func([]byte) *big.Int) []byte {
	hash := hashFunc([]byte(input))
	return hash.Bytes()
}
