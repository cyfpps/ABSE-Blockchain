package model

import (
	"math/big"
)

// UserDecryptionKey 定义用户解密密钥结构体
type UserDecryptionKey struct {
	S1 *big.Int
	S2 *big.Int
	S  map[string]*big.Int
}

// GenerateUserDecryptionKey 根据系统属性集和访问控制策略生成用户解密密钥
func GenerateUserDecryptionKey(globalParams *GlobalParams, masterKey *MasterKey, attributes []string) (*UserDecryptionKey, error) {
	// 生成随机数 t
	t, err := generateRandomElement(globalParams.G)
	if err != nil {
		return nil, err
	}

	// 计算 S1 和 S2
	s1 := new(big.Int).Exp(globalParams.G, new(big.Int).Add(globalParams.GAlpha, new(big.Int).Mul(globalParams.GBeta, t)), nil)
	s2 := new(big.Int).Exp(globalParams.G, t, nil)

	// 计算属性密钥
	attributeKeys := make(map[string]*big.Int)
	for _, attr := range attributes {
		h, found := masterKey.MK[attr]
		if !found {
			return nil, err
		}
		attrKey := new(big.Int).Exp(h, t, globalParams.G)
		attributeKeys[attr] = attrKey
	}

	return &UserDecryptionKey{
		S1: s1,
		S2: s2,
		S:  attributeKeys,
	}, nil
}
