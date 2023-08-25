package MODEL

import (
	"github.com/Nik-U/pbc"
)

type UserKeys struct {
	SK    map[string]*pbc.Element
	PkDO  *pbc.Element
	PkDU  *pbc.Element
	SkDO  *pbc.Element // 新增
	SkDU  *pbc.Element // 新增
	GBFA  *BloomFilter
	SCirc []int
}

func KeyGen(PK PublicKey, MK MasterKey, S []string, P []string) UserKeys {
	pairing := PK.G.Pairing()

	// 1) 为数据所有者和用户生成密钥
	xi := pairing.NewZr().Rand()
	pkDO := pairing.NewG1().PowZn(PK.G, xi)

	psi := pairing.NewZr().Rand()
	pkDU := pairing.NewG1().PowZn(PK.G, psi)

	// 2) 为属性生成 GBFA
	bf := NewBloomFilter()
	for _, attr := range S {
		bf.Add(attr)
	}

	// 3) 生成用户解密密钥
	t := pairing.NewZr().Rand()
	S1 := pairing.NewG1().PowZn(PK.Ga, pairing.NewZr().Add(MK.Alpha, pairing.NewZr().Mul(MK.Beta, t)))
	S2 := pairing.NewG1().PowZn(PK.G, t)

	SK := make(map[string]*pbc.Element)
	SK["S1"] = S1
	SK["S2"] = S2
	for _, attr := range S {
		hi := MK.Hi[attr]
		SK[attr] = pairing.NewG1().PowZn(PK.G, pairing.NewZr().Div(t, hi))
	}

	// 生成 SCirc
	SCirc := make([]int, 3*len(S))
	for i, attr := range S {
		hashes := bf.hash(attr)
		SCirc[3*i], SCirc[3*i+1], SCirc[3*i+2] = hashes[0], hashes[1], hashes[2]
	}

	return UserKeys{
		SK:    SK,
		PkDO:  pkDO,
		PkDU:  pkDU,
		SkDO:  psi, // 设置skDO的值
		SkDU:  xi,  // 设置skDU的值
		GBFA:  bf,
		SCirc: SCirc,
	}
}
