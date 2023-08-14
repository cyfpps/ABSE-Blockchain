package model

import (
	"crypto/sha256"
	"math/big"
)

// Trapdoor 定义陷门结构体
type Trapdoor struct {
	TW1 *big.Int
	TW2 *big.Int
	TW3 *big.Int
}

// GenerateTrapdoor 生成陷门
func GenerateTrapdoor(pkDO *big.Int, skDU *big.Int, keywords []string, omega *big.Int, g *big.Int, h *big.Int, eAlpha *big.Int) (*Trapdoor, error) {
	// 计算 Sec
	sec := new(big.Int)
	for _, kw := range keywords {
		h2Value := generateHashFunction2([]byte(kw))
		sec.Sub(sec, h2Value)
	}

	// 计算 TW1
	tw1 := new(big.Int).Mul(skDU, omega)

	// 计算 TW2
	expValue := new(big.Int).Exp(g, sec, nil) // 不需要 mod p
	pkDOgSec := new(big.Int).Mul(pkDO, expValue)
	tw2 := new(big.Int).Exp(pkDOgSec, omega, nil) // 不需要 mod p

	// 计算 TW3
	tw3 := new(big.Int).Exp(eAlpha, omega, nil) // 不需要 mod p

	return &Trapdoor{
		TW1: tw1,
		TW2: tw2,
		TW3: tw3,
	}, nil
}

func generateHashFunction2(data []byte) *big.Int {
	hash := sha256.Sum256(data)
	return new(big.Int).SetBytes(hash[:])
}
