package model

import (
	"math/big"
)

// DataOwner 数据拥有者结构体
type DataOwner struct {
	PrivateKey *big.Int
	PKDO       *big.Int
}

// NewDataOwner 创建一个数据拥有者实例
func NewDataOwner(globalParams *GlobalParams) (*DataOwner, error) {
	privateKey, err := generateRandomElement(globalParams.G)
	if err != nil {
		return nil, err
	}
	pkDO := new(big.Int).Exp(globalParams.G, privateKey, nil)
	return &DataOwner{
		PrivateKey: privateKey,
		PKDO:       pkDO,
	}, nil
}
