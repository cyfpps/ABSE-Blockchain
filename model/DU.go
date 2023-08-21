package model

import (
	"math/big"

	"github.com/Nik-U/pbc"
)

// DataUser 数据拥有者结构体
type DataUser struct {
	PrivateKey *big.Int
	PKDU       *pbc.Element
}

/*
// NewDataUser 创建一个数据拥有者实例
func NewDataUser(p *big.Int) (*DataUser, error) {
	// 生成一个随机的素数字符串，位数为 1024
	pStr, err := GenerateRandomPrime(1024)
	if err != nil {
		log.Fatal("Failed to generate random prime:", err)
	}

	// 将字符串转换为大整数
	privateKey, ok := new(big.Int).SetString(pStr, 10)
	if !ok {
		log.Fatal("Failed to convert p to big integer")
	}

	// 使用生成的素数和一个随机元素生成数据拥有者的公钥
	params := pbc.GenerateA1(p)
	pairing := pbc.NewPairing(params)
	expZr := pairing.NewZr().Rand()
	pkDU := pairing.NewG1().PowZn(p, expZr)

	return &DataUser{
		PrivateKey: privateKey,
		PKDU:       pkDU,
	}, nil
}

// GenerateRandomExponentAndPower 生成随机指数和指数运算结果
func GenerateRandomExponentAndPower(base *pbc.Element) (*pbc.Element, *pbc.Element) {
	pairing := base.Pairing

	// 生成随机的 Zr 类型的指数
	randomExponent := pairing.NewZr().Rand()

	// 进行指数运算
	powerResult := pairing.NewG1().PowZn(base, randomExponent)

	return randomExponent, powerResult
}
*/
