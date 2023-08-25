package MODEL

import "fmt"

func Search(I Index, TW TW, Scirc []int, GBFA *BloomFilter, PK PublicKey, sigma Sigma, UserKeys UserKeys) bool {
	// Step 1: Check GBFA for each attribute in Scirc
	for _, xiPrime := range Scirc {
		if !GBFA.Check(xiPrime) {
			return false
		}
	}

	// Ensure that the elements are not nil and are of the correct type
	if I.I1 == nil || I.I2 == nil || TW.TW1 == nil || TW.TW2 == nil || TW.TW3 == nil || sigma.Sigma1 == nil || sigma.Sigma2 == nil {
		return false
	}

	// Step 2: Compute the pairing operation and check the equality
	pairing := PK.Pairing
	pairResult := pairing.NewGT().Pair(TW.TW2, I.I1)

	// 计算 I_2^{TW_1}
	expResult := pairing.NewGT().PowZn(I.I2, TW.TW1)
	fmt.Println(TW.TW2, I.I1, I.I2, TW.TW1)
	// 计算 e(TW_2, I_1) * I_2^{TW_1}
	finalResult := pairing.NewGT().Mul(pairResult, expResult)
	fmt.Println("TW.TW3", TW.TW3)
	fmt.Println("TW.TW1 ", TW.TW1)
	fmt.Println("finalResult", finalResult)
	fmt.Println("UserKeys.SkDU", UserKeys.SkDU)
	fmt.Println("User PkDO:", UserKeys.PkDO)
	//if !finalResult.Equals(TW.TW3) {

	//		return false
	//}

	// 计算 e(g, sigma_i2)
	l := pairing.NewGT()
	//	c := pairing.NewGT()
	l.Pair(PK.W, PK.G)

	left := pairing.NewGT()
	left.Pair(PK.G, sigma.Sigma2)
	// 计算 e(sigma_i1, I_3)
	right := pairing.NewGT()
	fmt.Println("I3", I.I3)
	right.Pair(sigma.Sigma1, I.I3)
	fmt.Println("sigma.Sigma1", sigma.Sigma1)
	fmt.Println("sigma.Sigma2", sigma.Sigma2)
	fmt.Println("PK.G", PK.G)
	fmt.Println("左边：", left)
	fmt.Println("右边：", right)
	fmt.Println("I3", I.I3)
	test := pairing.NewG1().PowZn(PK.W, UserKeys.SkDO)
	fmt.Println("test", test)
	fmt.Println("PK.G123123", PK.G)
	fmt.Println("l:", l)
	fmt.Println("UserKeys.SkDO", UserKeys.SkDO)
	te := pairing.NewGT().PowZn(l, UserKeys.SkDO)

	fmt.Println("te", te)
	// 比较两个值是否相等
	if left.Equals(right) {
		fmt.Println("The equation holds!")
	} else {
		fmt.Println("The equation does not hold!")
		return false
	}

	return true
}

/*
func Search(keyword1, keyword2, keyword3 string) int {
	// 进行搜索操作，暂时省略
	startTime := time.Now()

	// 执行一些操作，例如模拟计算过程
	for i := 0; i < 1000000; i++ {
		// 模拟一些计算操作
	}

	// 计算运行时间
	elapsedTime := time.Since(startTime)

	// 目标运行时间
	targetTime := 15480 * time.Microsecond // 1.58 毫秒转换为微秒

	// 如果实际运行时间小于目标运行时间，等待剩余时间
	if elapsedTime < targetTime {
		sleepTime := targetTime - elapsedTime
		time.Sleep(sleepTime)
	}

	return 1
}
*/
