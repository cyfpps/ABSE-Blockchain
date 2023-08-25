package MODEL

import "fmt"

/*func Revocation(PK PublicKey, SK UserKeys, revokedAttrs []string) (updatedUK map[string]*pbc.Element, updatedS map[string]*pbc.Element, updatedCy map[Attribute]*pbc.Element) {
	pairing := PK.G.Pairing()
	updatedUK = make(map[string]*pbc.Element)
	updatedS = make(map[string]*pbc.Element)
	updatedCy = make(map[Attribute]*pbc.Element)

	for _, attr := range revokedAttrs {
		// 随机选择 h_i
		hBar := pairing.NewZr().Rand()
		Lk := hBar.Div(hBar, PK.UK[attr]) // hBar / h_i

		// 更新 UK_i
		updatedUK[attr] = PK.G.PowZn(PK.G, SK.UK[attr]).PowZn(PK.G.PowZn(PK.G, SK.UK[attr]), Lk)

		// 更新 S_i
		updatedS[attr] = SK.SK[attr].PowZn(SK.SK[attr], Lk.Invert(Lk))

		// 更新 C_y
		// 这里我们假设CT.Cy[attr]存在
		updatedCy[attr] = CT.Cy[attr].PowZn(CT.Cy[attr], Lk)

		GBFA.Remove(attr)
	}

	return updatedUK, updatedS, updatedCy
}
*/

func Revocation(SK UserKeys, revokedAttrs []string) UserKeys {
	for _, attr := range revokedAttrs {
		fmt.Println("撤销:", attr)
		SK.GBFA.Remove(attr)
	}

	//	SK.GBFA.Remove("attr2")
	return SK
}
