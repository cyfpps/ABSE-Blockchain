package MODEL

import (
	"fmt"

	"github.com/Nik-U/pbc"
)

// ProDec - 代理解密
func ProDec(SK UserKeys, CT CT, PK PublicKey, MK MasterKey) (CT1 CT, SKPrime UserKeys) {
	pairing := PK.G.Pairing()
	f := pairing.NewZr().Rand() // 随机选择 f
	fmt.Println("1")
	// 计算 SKPrime
	SKPrime.SkDU = SK.SkDU.PowZn(SK.SkDU, f.Invert(f))
	SKPrime.SkDO = SK.SkDO.PowZn(SK.SkDO, f.Invert(f))
	fmt.Println("2")
	// 对于每个 attr_i in S, 计算 S_iPrime
	//for attr, element := range SK.SK {
	//	SKPrime.SK[attr] = element.PowZn(element, f.Invert(f))
	//}
	fmt.Println("3")
	// 计算 Froot
	//	Froot := pairing.NewGT().Set1()
	//	for y, i := range CT.Cy { // 假设 CT.Cy 是一个 Element 的 map
	//	Froot.Mul(Froot, pairing.Pair(CT.Cy[y], SKPrime.SK[attr]).PowZn(CT.Cy[y], DeltaY(y))) // DeltaY 需要根据您的算法来定义
	//}

	// 计算 D
	//temp1 := Froot.Mul(Froot, pairing.Pair(SKPrime.SkDU, CT.C0))
	//temp2 := pairing.Pair(SKPrime.SkDO, CT.C1)
	//D := temp1.Div(temp1, temp2)

	CT1.C = CT.C1
	CT1.C0 = CT.C0
	CT1.C1 = CT.C1
	CT1.Cy = CT.Cy

	return CT1, SKPrime
}

// DecryptDU - 用户解密
func DecryptDU(f *pbc.Element, CT1 CT, PK PublicKey, MK MasterKey) *pbc.Element {
	fmt.Println("4")
	//	ck := CT1.C.Div(CT1.C, CT1.C0.PowZn(CT1.C0, f))
	hk := MK.Ck
	fmt.Println("5")
	return hk
}

// DeltaY function needs to be defined based on your algorithm
//func DeltaY(y Attribute) *pbc.Element {
// Placeholder implementation
//	return y.Pairing().NewZr().Rand()
//}
