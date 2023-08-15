package model

import (
	"github.com/Nik-U/pbc"
)

// SearchAlgorithm 实现搜索算法
func SearchAlgorithm(I Index, TW *Trapdoor, SStar []*pbc.Element) int {
	n := len(SStar)
	gt := pbc.NewGT()
	z := gt.NewZr()

	// 条件（1）：检查 S* 是否满足访问策略
	validSStar := true
	for i := 0; i < n; i++ {
		if !IsAttributeSatisfied(SStar[i]) {
			validSStar = false
			break
		}
	}

	// 条件（2）：验证等式是否成立
	if validSStar {
		TW2 := pbc.NewG2().SetBytes(TW.TW2.Bytes())
		I1 := pbc.NewG1().SetBytes(I.I1.Bytes())
		I2 := pbc.NewG2().SetBytes(I.I2.Bytes())
		TW1 := z.SetFromHash(I1).PowZn(I2)
		TW3 := z.Mul(TW1, TW2)

		if TW3.Equals(TW.TW3) {
			return 1
		}

		for i := 0; i < n; i++ {
			sigma2 := pbc.NewG1().SetBytes(SStar[i].Bytes())
			sigma1 := pbc.NewG1().SetBytes(I.I3.Bytes())
			g := pbc.NewG1().SetBytes(I.I3.Bytes())

			if gt.Pair(sigma2, g).Equals(gt.Pair(sigma1, I2)) {
				return 1
			}
		}
	}

	return 0
}

// IsAttributeSatisfied 检查属性是否满足访问策略
func IsAttributeSatisfied(attr *pbc.Element) bool {
	// 在此添加逻辑以检查属性是否满足访问策略
	// 返回 true 表示满足，返回 false 表示不满足
	return true
}
