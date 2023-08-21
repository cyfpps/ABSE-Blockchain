package model

import (
	"time"

	"github.com/Nik-U/pbc"
)

// Ciphertext 密文结构体
type Ciphertext struct {
	C  *pbc.Element
	C0 *pbc.Element
	C1 *pbc.Element
	Cy map[string]*pbc.Element
}

/*
// CTEncrypt 密文生成算法
func CTEncrypt(PK *GlobalParams, P []string, ck string, r1 *pbc.Element, masterKey *MasterKey, r2 *pbc.Element) (*Ciphertext, error) {
	// 将字符串类型的 ck 转换为字节切片
	ckBytes := []byte(ck)

	// 计算 e(g, g)^(alpha * r2)
	egAlphaR2 := PK.Pairing.NewGT().PowZn(PK.EAlpha, r2)

	// 计算 C = ck * e(g, g)^(alpha * r2)
	C := PK.Pairing.NewGT().SetBytes(ckBytes)
	C.Mul(C, egAlphaR2)

	// 计算 C0 = g^(r2)
	C0 := new(pbc.Element).PowZn(PK.G, r2)

	// 计算 C1 = g^(beta * r2 + r1)
	betaR2 := new(pbc.Element).Mul(PK.GBeta, r2)
	betaR2.Add(betaR2, r1) // beta * r2 + r1
	C1 := new(pbc.Element).PowZn(PK.G, betaR2)

	// 初始化 C_y 映射
	Cy := make(map[string]*pbc.Element)
	for _, attr := range P {
		h := masterKey.MK[attr]
		cyElement := new(pbc.Element).PowZn(PK.G, h) // 计算 Cy[attr]
		Cy[attr] = cyElement
	}

	// 在这里根据访问控制树生成 C_y

	// 构造密文结构体
	ciphertext := &Ciphertext{
		C:  C,
		C0: C0,
		C1: C1,
		Cy: Cy,
	}

	return ciphertext, nil
}
*/
func CTEnc(w string) int {
	// 进行搜索操作，暂时省略
	startTime := time.Now()
	for time.Since(startTime) < time.Millisecond {
		// 模拟计算操作
	}
	return 1
}
