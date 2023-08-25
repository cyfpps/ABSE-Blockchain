package MODEL

import (
	"fmt"

	"github.com/Nik-U/pbc"
)

type TW struct {
	TW1 Element
	TW2 Element
	TW3 Element
}

func Trapdoor(KW []string, PK PublicKey, sk_DU *pbc.Element, pk_DO *pbc.Element) TW {
	// 直接从PublicKey结构体中获取配对
	pairing := PK.Pairing

	// 计算Sec = -Σ H2(kwi')
	Sec := pairing.NewZr().Set0()
	for _, kw := range KW {
		h2 := hashToBigInt(PK.H2, kw)
		h2Element := pairing.NewZr().SetBig(h2)
		Sec = Sec.Sub(Sec, h2Element)
	}
	// 选择随机数Ω
	Omega := pairing.NewZr().Rand()
	// 计算TW1, TW2, TW3
	TW1 := sk_DU.Mul(sk_DU, Omega)
	temp := pairing.NewG1().Mul(pk_DO, pairing.NewG1().PowZn(PK.G, Sec))
	TW2 := temp.PowZn(temp, TW1)

	TW3 := pairing.NewGT().PowZn(pairing.NewGT().Pair(PK.W, PK.G), TW1)
	fmt.Println(TW3)
	return TW{
		TW1: TW1,
		TW2: TW2,
		TW3: TW3,
	}
}
