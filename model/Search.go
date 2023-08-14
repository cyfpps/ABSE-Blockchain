package model

import (
	"math/big"
)

// SearchAlgorithmResult 表示搜索算法的结果
type SearchAlgorithmResult int

const (
	MatchFound SearchAlgorithmResult = 1 // 匹配成功
	NoMatch    SearchAlgorithmResult = 0 // 无匹配
)

// Search 使用搜索算法验证陷门并返回结果
func Search(tw *Trapdoor, sCirc []*big.Int, I1, I2, I3 *big.Int, sigmaI1, sigmaI2 *big.Int) SearchAlgorithmResult {
	// 条件1：验证属性集是否满足访问策略
	matchAttr := true
	for _, s := range sCirc {
		if s.Cmp(big.NewInt(0)) == 0 {
			matchAttr = false
			break
		}
	}
	if !matchAttr {
		return NoMatch
	}

	// 条件2：验证等式
	leftSide1 := new(pairing).MulLine([][2]*big.Int{{tw.TW2, I1}, {I2, tw.TW1}})
	if leftSide1.Cmp(tw.TW3) != 0 {
		return NoMatch
	}

	leftSide2 := new(pairing).Pair(sigmaI2, g)
	rightSide2 := new(pairing).Pair(sigmaI1, I3)
	if leftSide2.Cmp(rightSide2) != 0 {
		return NoMatch
	}

	return MatchFound
}
