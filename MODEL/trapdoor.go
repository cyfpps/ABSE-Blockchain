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

func Trapdoor(KW []string, PK PublicKey, sk_DU *pbc.Element, pk_DO *pbc.Element, UserKeys UserKeys) TW {
	// 直接从PublicKey结构体中获取配对
	pairing := PK.Pairing

	// 计算Sec = -Σ H2(kwi')
	Sec := pairing.NewZr().Set0()
	for _, kw := range KW {
		h2 := hashToBigInt(PK.H2, kw)
		h2Element := pairing.NewZr().SetBig(h2)
		Sec = Sec.Sub(Sec, h2Element)
	}
	fmt.Println("Sec:", Sec)
	Secte := pairing.NewZr().Add(Sec, UserKeys.SkDO)
	fmt.Println("Secte:", Secte)
	// 选择随机数Ω
	l := pairing.NewGT()
	l.Pair(PK.W, PK.G)
	fmt.Println("l:", l)
	Omega := pairing.NewZr().Rand()
	fmt.Println("Omega", Omega)
	// 计算TW1, TW2, TW3
	TW1 := pairing.NewZr().Mul(sk_DU, Omega)
	fmt.Println("TW1 ", TW1)
	temp := pairing.NewG1().Mul(pk_DO, pairing.NewG1().PowZn(PK.G, Sec))
	TW2 := pairing.NewG1().PowZn(temp, TW1)
	lS := pairing.NewGT().PowZn(l, TW1)
	fmt.Println("lS:", lS)
	TW3 := pairing.NewGT().PowZn(pairing.NewGT().Pair(PK.W, PK.G), TW1)
	fmt.Println("Omega", Omega)
	return TW{
		TW1: TW1,
		TW2: TW2,
		TW3: TW3,
	}
}
