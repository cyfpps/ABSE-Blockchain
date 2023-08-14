package main

import (
	"fmt"
	"time"

	"github.com/cyfpps/ABSE-Blockchain/model" // 修改为你的实际路径
)

func main() {
	// 初始化全局参数
	startTime := time.Now()
	globalParams, err := model.Setup(256) // 使用256位参数，可以根据需要调整
	if err != nil {
		fmt.Println("Error in setup:", err)
		return
	}
	setupDuration := time.Since(startTime)
	fmt.Printf("Setup 运行时间：%s\n", setupDuration)

	// 密钥生成
	startTime = time.Now()
	S := []string{"attr1", "attr2"} // 数据使用者属性集合
	SK, pkDO, pkDU, SStar, err := model.KeyGen(globalParams.PK, globalParams.MK, S)
	if err != nil {
		fmt.Println("Error in key generation:", err)
		return
	}
	keyGenDuration := time.Since(startTime)
	fmt.Printf("KeyGen 运行时间：%s\n", keyGenDuration)

	// 加密
	startTime = time.Now()
	KW := "keyword"
	ck := model.GenerateCiphertextKey(globalParams.PK, SK, SStar)
	CT, sigma, I, Addr, err := model.Encrypt(KW, globalParams.PK, pkDO, S, ck, "plaintext")
	if err != nil {
		fmt.Println("Error in encryption:", err)
		return
	}
	encryptionDuration := time.Since(startTime)
	fmt.Printf("Encrypt 运行时间：%s\n", encryptionDuration)

	// 陷门生成
	startTime = time.Now()
	KWPrime := []string{"kw1", "kw2"}
	TW, err := model.GenerateTrapdoor(KWPrime, pkDO, pkDU)
	if err != nil {
		fmt.Println("Error in trapdoor generation:", err)
		return
	}
	trapdoorDuration := time.Since(startTime)
	fmt.Printf("GenerateTrapdoor 运行时间：%s\n", trapdoorDuration)

	// 搜索
	startTime = time.Now()
	result := model.Search(I, TW, SStar)
	searchDuration := time.Since(startTime)
	fmt.Printf("Search 运行时间：%s\n", searchDuration)

	if result == 1 {
		fmt.Println("Search successful, CT and Addr sent to user.")
	} else {
		fmt.Println("Search failed.")
	}

	// 解密
	startTime = time.Now()
	m, err := model.Decrypt(SK, CT, Addr)
	if err != nil {
		fmt.Println("Error in decryption:", err)
		return
	}
	decryptionDuration := time.Since(startTime)
	fmt.Printf("Decrypt 运行时间：%s\n", decryptionDuration)
	fmt.Printf("成功解密，解密数据是：%s\n", m)

	// 撤销
	startTime = time.Now()
	overlineS := []string{"attr1"}
	overlineUK, overlineSI, overlineCy, err := model.Revocation(globalParams.PK, pkDO, overlineS)
	if err != nil {
		fmt.Println("Error in revocation:", err)
		return
	}
	revocationDuration := time.Since(startTime)
	fmt.Printf("Revocation 运行时间：%s\n", revocationDuration)
	fmt.Printf("撤销后的公钥组件：%v\n", overlineUK)
	fmt.Printf("撤销后的私钥组件：%v\n", overlineSI)
	fmt.Printf("撤销后的密文组件：%v\n", overlineCy)

	fmt.Println("运行成功！")
}
