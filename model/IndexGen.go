package model

import (
	"github.com/Nik-U/pbc"
)

// Index 定义索引结构体
type Index struct {
	I1 *pbc.Element
	I2 *pbc.Element
	I3 *pbc.Element
}

/*
// GenerateIndex 根据关键字组、数据拥有者和数据使用者生成索引
func GenerateIndex(keywords []string, dataOwner *DataOwner, dataUser *DataUser, globalParams *GlobalParams) (*Index, error) {
	// 计算关键字的哈希和
	hashedKeywordSum := globalParams.Pairing.NewZr().SetInt64(0)
	for _, keyword := range keywords {
		hashedKeyword := globalParams.H2([]byte(keyword)) // 使用 H2 哈希函数
		hashedKeywordSum.Add(hashedKeywordSum, globalParams.Pairing.NewZr().SetBytes(hashedKeyword))
	}

	// 生成随机数 rho
	rho, err := generateRandomElement(globalParams.p)
	if err != nil {
		return nil, err
	}

	// 计算索引的各个部分
	I1Num := globalParams.Pairing.NewZr().Set(rho).Neg(rho)
	for _, keyword := range keywords {
		h2 := globalParams.H2([]byte(keyword))
		I1Num.Add(I1Num, h2)
	}
	I1 := globalParams.Pairing.NewG1().SetFromHash(I1Num.Bytes())

	I2 := globalParams.Pairing.NewGT().PowZn(globalParams.Pairing.NewGT().Pair(globalParams.G, globalParams.G), rho)

	I3 := globalParams.W.PowZn(dataUser.PrivateKey)

	return &Index{
		I1: I1,
		I2: I2,
		I3: I3,
	}, nil
}
*/
