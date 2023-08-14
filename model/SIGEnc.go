package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// Signature 签名结构体
type Signature struct {
	Sigma1 *pbc.Element
	Sigma2 *pbc.Element
}

// GenerateSignature 生成密文签名
func GenerateSignature(CT *Ciphertext, dataOwner *DataOwner, globalParams *GlobalParams) (*Signature, error) {
	// 计算 H3(CT)
	H3CT := generateHashFunction(globalParams.G, CT.C.Bytes())

	// 计算 sigma1 = g^{H3(CT)}
	sigma1 := globalParams.G.NewElement()
	sigma1.PowZn(globalParams.G, new(big.Int).SetBytes(H3CT))

	// 计算 sigma2 = (w^{H3(CT)})^dataOwner.PrivateKey
	wToH3CT := globalParams.W.NewElement()
	wToH3CT.PowZn(globalParams.W, new(big.Int).SetBytes(H3CT))
	sigma2 := globalParams.G.NewElement()
	sigma2.PowZn(wToH3CT, dataOwner.PrivateKey)

	return &Signature{
		Sigma1: sigma1,
		Sigma2: sigma2,
	}, nil
}
