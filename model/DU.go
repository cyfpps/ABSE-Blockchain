package model

import (
	"math/big"
)

// DataUser 数据使用者结构体
type DataUser struct {
	PrivateKey *big.Int
	PKDU       *big.Int
}

// NewDataUser 创建一个数据使用者实例
func NewDataUser(globalParams *GlobalParams) (*DataUser, error) {
	privateKey, err := generateRandomElement(globalParams.G)
	if err != nil {
		return nil, err
	}
	pkDU := new(big.Int).Exp(globalParams.G, privateKey, nil)
	return &DataUser{
		PrivateKey: privateKey,
		PKDU:       pkDU,
	}, nil
}
