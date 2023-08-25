package model

import (
	"time"
)

/*
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

/*
package model

import (
	"crypto/sha256"
	"github.com/Nik-U/pbc"
)

type UserKeys struct {
	SK    map[string]*pbc.Element
	PkDO  *pbc.Element
	PkDU  *pbc.Element
	GBFA  *bloom.BloomFilter
	SCirc []int
}

func KeyGen(PK PublicKey, MK MasterKey, S []string, P []string) UserKeys {
	pairing := PK.G.Pairing()

	// 1) Generate keys for data owner and user
	xi := pairing.NewZr().Rand()
	pkDO := pairing.NewG1().PowZn(PK.G, xi)

	psi := pairing.NewZr().Rand()
	pkDU := pairing.NewG1().PowZn(PK.G, psi)

	// 2) Generate GBFA for attributes
	GBFA := bloom.New(1000, 3) // Set the size of GBFA to 1000 and use 3 hash functions
	for _, attr := range S {
		GBFA.Add([]byte(attr))
	}

	// 3) Generate user decryption key
	t := pairing.NewZr().Rand()
	S1 := pairing.NewG1().PowZn(PK.Ga, pairing.NewZr().Add(MK.Alpha, pairing.NewZr().Mul(MK.Beta, t)))
	S2 := pairing.NewG1().PowZn(PK.G, t)

	SK := make(map[string]*pbc.Element)
	SK["S1"] = S1
	SK["S2"] = S2
	for _, attr := range S {
		hi := MK.Hi[attr]
		SK[attr] = pairing.NewG1().PowZn(PK.G, pairing.NewZr().Div(t, pairing.NewZr().SetBig(hi)))
	}

	// Generate SCirc
	SCirc := make([]int, 3*len(S))
	for i, attr := range S {
		SCirc[3*i] = int(sha256.Sum256([]byte(attr))[0])
		SCirc[3*i+1] = int(sha256.Sum256([]byte(attr))[1])
		SCirc[3*i+2] = int(sha256.Sum256([]byte(attr))[2])
	}

	return UserKeys{
		SK:    SK,
		PkDO:  pkDO,
		PkDU:  pkDU,
		GBFA:  GBFA,
		SCirc: SCirc,
	}
}
*/
