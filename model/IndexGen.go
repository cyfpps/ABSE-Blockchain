package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// Index 定义索引结构体
type Index struct {
	I1 *pbc.Element
	I2 *big.Int
	I3 *big.Int
}

// GenerateIndex 根据关键字组、数据拥有者和数据使用者生成索引
func GenerateIndex(keywords []string, dataOwner *DataOwner, dataUser *DataUser, globalParams *GlobalParams) (*Index, error) {
	// 计算关键字的哈希和
	hashedKeywordSum := big.NewInt(0)
	for _, keyword := range keywords {
		hashedKeyword := generateHashFunction(globalParams.G, []byte(keyword))
		hashedKeywordSum.Add(hashedKeywordSum, new(big.Int).SetBytes(hashedKeyword))
	}

	// 生成随机数 rho
	rho, err := generateRandomElement(globalParams.G)
	if err != nil {
		return nil, err
	}

	// 计算索引的各个部分
	I1Num := globalParams.W.Exp(globalParams.G, new(big.Int).Neg(rho), globalParams.G)
	for _, keyword := range keywords {
		h2 := globalParams.H2([]byte(keyword))
		I1Num.Add(I1Num, h2)
	}
	I1Num.Exp(I1Num, new(big.Int).SetBytes(dataUser.PrivateKey.Bytes()), globalParams.G)
	I1 := globalParams.G.NewElement()
	I1.SetFromHash(I1Num.Bytes())

	I2 := new(big.Int).Exp(globalParams.EAlpha, rho, globalParams.G)
	I3 := new(big.Int).Exp(globalParams.W, dataUser.PrivateKey, globalParams.G)

	return &Index{
		I1: I1,
		I2: I2,
		I3: I3,
	}, nil
}
