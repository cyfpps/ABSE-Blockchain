package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// Trapdoor 陷门结构体
type Trapdoor struct {
	TW1 *big.Int
	TW2 *pbc.Element
	TW3 *pbc.Element
}

// GenerateTrapdoor 生成陷门
func GenerateTrapdoor(globalParams *GlobalParams, DO *DataOwner, DU *DataUser, keywords []string) (*Trapdoor, error) {
	// 选择随机数 Omega
	omega, err := generateRandomElement(globalParams.p)
	if err != nil {
		return nil, err
	}

	// 计算 Sec
	sec := new(big.Int)
	for _, keyword := range keywords {
		sec.Add(sec, generateHashFunction(globalParams.p, []byte(keyword)))
	}
	sec.Neg(sec)

	// 计算 TW1
	tw1 := new(big.Int).Mul(DU.PrivateKey, omega)

	// 计算 TW2
	tw2 := globalParams.G.NewElement()
	tw2.PowZn(DO.PKDO, new(big.Int).Mul(DU.PrivateKey, omega)).MulZn(tw2, globalParams.G.NewElement().PowZn(sec))

	// 计算 TW3
	tw3 := globalParams.Pairing.NewG2().Pair(globalParams.W, globalParams.G).PowZn(new(big.Int).Mul(DU.PrivateKey, omega))

	return &Trapdoor{
		TW1: tw1,
		TW2: tw2,
		TW3: tw3,
	}, nil
}
