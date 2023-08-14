package model

import (
	"crypto/rand"
	"math/big"
)

// 定义全局参数结构体
type GlobalParams struct {
	G      *big.Int
	W      *big.Int
	GAlpha *big.Int
	GBeta  *big.Int
	EAlpha *big.Int
	H1     func([]byte) *big.Int
	H2     func([]byte) *big.Int
	H3     func([]byte) *big.Int
	UK     map[string]*big.Int
}

// 初始化全局参数
func Setup(lambda int) (*GlobalParams, error) {
	p, err := generatePrime(lambda)
	if err != nil {
		return nil, err
	}

	g, err := generateRandomElement(p)
	if err != nil {
		return nil, err
	}

	w, err := generateRandomElement(p)
	if err != nil {
		return nil, err
	}

	alpha, err := generateRandomElement(p)
	if err != nil {
		return nil, err
	}

	beta, err := generateRandomElement(p)
	if err != nil {
		return nil, err
	}

	gAlpha := new(big.Int).Exp(g, alpha, p)
	gBeta := new(big.Int).Exp(g, beta, p)

	eAlpha := new(big.Int).Exp(g, alpha, p) // 此处可能需要根据实际情况进行调整

	// 定义哈希函数
	h1 := func(data []byte) *big.Int {
		hash := generateHashFunction(p, data)
		return hash
	}

	h2 := func(data []byte) *big.Int {
		hash := generateHashFunction(p, data)
		return hash
	}

	h3 := func(data []byte) *big.Int {
		hash := generateHashFunction(p, data)
		return hash
	}

	// 初始化用户密钥
	UK := make(map[string]*big.Int)

	// TODO: 初始化用户密钥

	params := &GlobalParams{
		G:      g,
		W:      w,
		GAlpha: gAlpha,
		GBeta:  gBeta,
		EAlpha: eAlpha,
		H1:     h1,
		H2:     h2,
		H3:     h3,
		UK:     UK,
	}

	return params, nil
}

// 生成指定位数的素数
func generatePrime(bits int) (*big.Int, error) {
	for {
		num, err := rand.Prime(rand.Reader, bits)
		if err != nil {
			return nil, err
		}

		if num.ProbablyPrime(20) {
			return num, nil
		}
	}
}

// 生成随机元素
func generateRandomElement(p *big.Int) (*big.Int, error) {
	randNum, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, err
	}
	return randNum, nil
}

// 生成哈希函数
func generateHashFunction(p *big.Int, data []byte) *big.Int {
	hash := new(big.Int).SetBytes(data)
	hash.Mod(hash, p)
	return hash
}
