package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// DecryptDU 解密算法
func DecryptDU(f *big.Int, CT1 *Ciphertext, Addr *big.Int) ([]byte, error) {
	// 获得加密文件 E(m)
	E := CT1.C

	// 计算密钥 ck
	ck := CalculateCK(f, CT1, Addr)

	// 使用密钥 ck 解密密文以获得明文文件 m
	m, err := DecryptCipherText(ck, E)
	if err != nil {
		return nil, err
	}

	return m, nil
}

// CalculateCK 计算密钥 ck
func CalculateCK(f *big.Int, CT1 *Ciphertext, Addr *big.Int) *pbc.Element {
	g := CT1.Params.NewG1().SetBytes(CT1.Params.G.Bytes())

	Df := CT1.D.Exp(CT1.D, f.Neg(f)) // D^(-f)

	ck := CT1.C.Mul(CT1.C, g.PowZn(Df)) // ck * e(g, g)^(alpha r2) / e(g, g)^(alpha r2)

	return ck
}

// DecryptCipherText 使用密钥解密密文
func DecryptCipherText(ck *pbc.Element, E *pbc.Element) ([]byte, error) {
	// 使用密钥 ck 解密密文 E
	m := E.Mul(E, ck.Invert(ck)) // E * ck^(-1)

	// 转换为字节数组并返回
	return m.Bytes(), nil
}
