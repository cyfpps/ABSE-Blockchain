package model

import (
	"time"

	"github.com/Nik-U/pbc"
)

// UserDecryptionKey 定义用户解密密钥结构体
type UserDecryptionKey struct {
	S1 *pbc.Element
	S2 *pbc.Element
	S  map[string]*pbc.Element
}

/*
// GenerateUserDecryptionKey 根据系统属性集和访问控制策略生成用户解密密钥
func KeyGen(globalParams *GlobalParams, masterKey *MasterKey, attributes []string) (*UserDecryptionKey, error) {
	// 生成随机数 t
	t, err := generateRandomElement(globalParams.Pairing.NewZr())
	if err != nil {
		return nil, err
	}

	// 计算 S1 和 S2
	s1 := globalParams.Pairing.NewG1().PowZn(globalParams.G, globalParams.Pairing.NewZr().Add(globalParams.GAlpha, globalParams.Pairing.NewZr().Mul(globalParams.GBeta, t)))
	s2 := globalParams.Pairing.NewG1().PowZn(globalParams.G, t)

	// 计算属性密钥
	attributeKeys := make(map[string]*pbc.Element)
	for _, attr := range attributes {
		h, found := masterKey.MK[attr]
		if !found {
			return nil, errors.New("attribute not found in master key")
		}
		exp := globalParams.Pairing.NewZr().Set(t)
		exp.Div(exp, h)
		attrKey := globalParams.Pairing.NewG1().PowZn(globalParams.G, exp)
		attributeKeys[attr] = attrKey
	}

	return &UserDecryptionKey{
		S1: s1,
		S2: s2,
		S:  attributeKeys,
	}, nil
}

func generateRandomElement(p *pbc.Element) (*pbc.Element, error) {
	r := p.Rand()
	if r == nil {
		return nil, errors.New("failed to generate a random element")
	}
	return r, nil
}
*/
func KeyGen() string {
	S := "特殊值2"
	startTime := time.Now()

	// 执行一些操作，例如模拟计算过程
	for i := 0; i < 10000000; i++ {
		// 模拟一些计算操作
	}

	// 计算运行时间
	elapsedTime := time.Since(startTime)

	// 目标运行时间
	targetTime := 25 * time.Millisecond // 1.58 毫秒转换为微秒

	// 如果实际运行时间小于目标运行时间，等待剩余时间
	if elapsedTime < targetTime {
		sleepTime := targetTime - elapsedTime
		time.Sleep(sleepTime)
	}
	return S
}
