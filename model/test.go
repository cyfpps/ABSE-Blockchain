package model

import (
	"fmt"
	"time"

	"github.com/Nik-U/pbc"
)

// GenerateAndComputeTimes 生成群元素 g，并计算相关时间
func GenerateAndComputeTimes(p *pbc.Element) error {
	// 初始化 pbc.Pairing
	params := pbc.GenerateA(160, 512)
	pairing := pbc.NewPairing(params)

	// 生成群元素 g
	startTime := time.Now()
	g := pairing.NewG1().Rand()
	generationDuration := time.Since(startTime)
	fmt.Printf("生成群元素 g 的时间：%s\n", generationDuration)

	// 计算 g 的指数运算时间
	expStartTime := time.Now()
	expResult := pairing.NewG1().PowZn(g, p)
	expDuration := time.Since(expStartTime)
	fmt.Printf("Result of g's Exponential Operation: %s\n", expResult)
	fmt.Printf("计算 g 的指数运算时间：%s\n", expDuration)

	// 进行一次双线性配对 e(g, g)
	pairingStartTime := time.Now()
	gt := pairing.NewGT().Pair(g, g)
	pairingDuration := time.Since(pairingStartTime)
	fmt.Printf("进行双线性配对 e(g, g) 的时间：%s\n", pairingDuration)

	// 计算 gt 的指数运算时间
	gtExpStartTime := time.Now()
	gtExpResult := pairing.NewGT().PowZn(gt, p)
	gtExpDuration := time.Since(gtExpStartTime)
	fmt.Printf("Result of gt's Exponential Operation: %s\n", gtExpResult)
	fmt.Printf("计算 gt 的指数运算时间：%s\n", gtExpDuration)

	return nil
}
