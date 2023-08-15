package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// ProxyDecryptionAlgorithm 实现代理解密算法
func ProxyDecryptionAlgorithm(SKPrime *UserDecryptionKey, CT *Ciphertext) (*Ciphertext, *UserDecryptionKey) {
	// 获取代理解密算法需要的参数
	t := CT.T
	r1 := SKPrime.S1
	r2 := SKPrime.S2
	alpha := SKPrime.S["alpha"]
	beta := SKPrime.S["beta"]

	// 计算 F_root
	FRoot := PairingProduct(SKPrime, CT, t, r1)

	// 若无法计算 F_root，则输出 CT1 = {D: nil, C: nil}
	if FRoot == nil {
		return &Ciphertext{D: nil, C: nil}, SKPrime
	}

	// 计算 D
	D := CalculateD(FRoot, alpha, beta, r1, r2, t)

	// 构建 CT1 并返回
	CT1 := &Ciphertext{D: D, C: CT.C}
	return CT1, SKPrime
}

// PairingProduct 计算 F_root
func PairingProduct(SKPrime *UserDecryptionKey, CT *Ciphertext, t *big.Int, r1 *big.Int) *pbc.Element {
	g := SKPrime.Params.NewG1().SetBytes(SKPrime.G.Bytes())

	FRoot := SKPrime.Params.NewGT().Pair(g, g)
	DeltaY := SKPrime.Params.NewZr().SetInt64(1)

	for _, y := range CT.Y {
		i := y.Attr
		SiPrime := SKPrime.Si[i]

		SiblingSet := SiblingSet(SKPrime, i)

		deltaY := CalculateDeltaY(SiblingSet, y.NodeIndex, SKPrime.Params)

		FRoot.Mul(FRoot, SKPrime.Params.NewGT().Pair(CT.C2, SiPrime).Exp(SKPrime.Params.NewGT().Pair(CT.C1, SKPrime.S1), deltaY))
		DeltaY.Mul(DeltaY, deltaY)
	}

	if !DeltaY.Equals(SKPrime.Params.NewZr().SetInt64(1)) {
		return FRoot.PowZn(DeltaY)
	}
	return nil
}

// SiblingSet 计算 SiblingSet
func SiblingSet(SKPrime *UserDecryptionKey, i int) []*big.Int {
	SiblingSet := []*big.Int{}
	for _, z := range SKPrime.S {
		SiblingSet = append(SiblingSet, z)
	}
	return SiblingSet
}

// CalculateDeltaY 计算 Delta_y
func CalculateDeltaY(SiblingSet []*big.Int, y int, params *pbc.Params) *pbc.Element {
	DeltaY := params.NewZr().SetInt64(1)
	for _, x := range SiblingSet {
		if x != SiblingSet[y] {
			i := SiblingSet[y].Index()
			DeltaY.Mul(DeltaY, x)
		}
	}
	return DeltaY
}

// CalculateD 计算 D
func CalculateD(FRoot *pbc.Element, alpha *big.Int, beta *big.Int, r1 *big.Int, r2 *big.Int, t *big.Int) *pbc.Element {
	g := r1.NewZr().SetBytes(alpha.Bytes())
	h := r2.NewZr().SetBytes(beta.Bytes())

	D := FRoot.Exp(FRoot, h)
	D.Mul(D, g)
	return D
}
