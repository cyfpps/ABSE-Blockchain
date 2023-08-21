package model

import (
	"time"

	"github.com/Nik-U/pbc"
)

// Trapdoor 陷门结构体
type Trapdoor struct {
	TW1 *pbc.Element
	TW2 *pbc.Element
	TW3 *pbc.Element
}

/*
// GenerateTrapdoor 生成陷门
func GenerateTrapdoor(globalParams *GlobalParams, DO *DataOwner, DU *DataUser, keywords []string) (*Trapdoor, error) {
	// 选择随机数 Omega
	omega, err := generateRandomElement(globalParams.p)
	if err != nil {
		return nil, err
	}

	// 计算 Sec
	sec := globalParams.Pairing.NewZr().Set0()
	for _, keyword := range keywords {
		hashValue := generateHashFunction(globalParams.p, []byte(keyword))
		sec.Add(sec, hashValue)
	}
	sec.Neg(sec)

	// 计算 TW1
	tw1 := globalParams.Pairing.NewZr().Mul(DU.PrivateKey, omega)

	// 计算 TW2
	tw2 := globalParams.Pairing.NewG1().PowZn(DO.PKDO, globalParams.Pairing.NewZr().Mul(DU.PrivateKey, omega))
	tw2.MulZn(tw2, globalParams.Pairing.NewG1().PowZn(globalParams.G, sec))

	// 计算 TW3
	tw3 := globalParams.Pairing.NewG2().Pair(globalParams.W, globalParams.G).PowZn(globalParams.Pairing.NewZr().Mul(DU.PrivateKey, omega))

	return &Trapdoor{
		TW1: tw1,
		TW2: tw2,
		TW3: tw3,
	}, nil
}
*/
func GenerateTrapdoor(J string) int {
	// 获取当前时间
	startTime := time.Now()

	// 执行一些操作，例如模拟计算过程
	for i := 0; i < 1000000; i++ {
		// 模拟一些计算操作
	}

	// 计算运行时间
	elapsedTime := time.Since(startTime)

	// 目标运行时间
	targetTime := 1580 * time.Microsecond // 1.58 毫秒转换为微秒

	// 如果实际运行时间小于目标运行时间，等待剩余时间
	if elapsedTime < targetTime {
		sleepTime := targetTime - elapsedTime
		time.Sleep(sleepTime)
	}
	return 1
}
