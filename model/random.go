package model

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/Nik-U/pbc"
)

// GenerateRandomPrime 生成一个指定位数的随机素数
func GenerateRandomPrime(bits int) (string, error) {
	// 调用 rand.Prime 函数生成一个指定位数的随机素数
	// 参数 rand.Reader 表示使用系统提供的随机性源
	// 参数 bits 指定要生成的素数的位数
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		log.Fatal("Failed to generate random prime:", err)
	}

	// 将生成的素数转换为字符串表示，并返回
	// 这里使用 p.String() 将大整数 p 转换为其十进制表示
	return p.String(), nil
}

func GenerateRandomElement() *pbc.Element {
	pStr, err := GenerateRandomPrime(14)
	if err != nil {
		fmt.Println("Error in setup:", err)
	}
	// 将字符串转换为大整数
	pBigInt, ok := new(big.Int).SetString(pStr, 10)
	if !ok {
		log.Fatal("Failed to convert p to big integer")
	}
	// 使用 pbc 库生成椭圆曲线参数
	params := pbc.GenerateA1(pBigInt) //pBigInt为位数（比特数）
	pairing := pbc.NewPairing(params)

	randomElement := pairing.NewG1().Rand() //创建一个新的 G1 元素
	return randomElement
}
