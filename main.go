package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/Nik-U/pbc"
	"github.com/cyfpps/ABSE-Blockchain/model" // 修改为你的实际路径
)

func main() {
	// 初始化全局参数
	pStr, err := model.GenerateRandomPrime(1024)

	// 将字符串转换为大整数
	pBigInt, ok := new(big.Int).SetString(pStr, 10)
	if !ok {
		log.Fatal("Failed to convert p to big integer")
	}

	// 使用 pbc 库生成椭圆曲线参数
	params := pbc.GenerateA1(pBigInt)
	pairing := pbc.NewPairing(params)
	g1 := pairing.NewG1().Rand()
	g2 := pairing.NewG2().Rand()

	// 计算 g1 的指数运算，并计算运行时间
	startTime := time.Now()
	expZr := pairing.NewZr().Rand()
	expG1 := pairing.NewG1().PowZn(g1, expZr)
	expG1Duration := time.Since(startTime)

	// 计算 g2 的指数运算，并计算运行时间
	startTime = time.Now()
	expG2 := pairing.NewG2().PowZn(g2, expZr)
	expG2Duration := time.Since(startTime)

	// 计算配对运算，并计算运行时间
	startTime = time.Now()
	e := pairing.NewGT().Pair(expG1, expG2)
	pairingDuration := time.Since(startTime)

	// 计算 e 的指数运算，并计算运行时间
	startTime = time.Now()
	expGT := pairing.NewGT().PowZn(e, expZr)
	expGTDuration := time.Since(startTime)

	// 打印结果和运行时间
	fmt.Println("原始 g1 =", g1)
	fmt.Println("原始 g2 =", g2)
	fmt.Println("指数运算后的 g1 =", expG1)
	fmt.Println("指数运算后的 g2 =", expG2)
	fmt.Println("e(指数运算后的 g1, 指数运算后的 g2) =", e)
	fmt.Println("e 的指数运算后的结果 =", expGT)
	fmt.Printf("g1 的指数运算时间：%s\n", expG1Duration)
	fmt.Printf("g2 的指数运算时间：%s\n", expG2Duration)
	fmt.Printf("配对运算时间：%s\n", pairingDuration)
	fmt.Printf("e 的指数运算时间：%s\n", expGTDuration)

	SK := "mySecretKey"
	CT := "encryptedText"
	CK := "encrypasdadtedText"
	Addr := "123.456.789.0"
	CTCK := "123.456.783249.0"
	keyword1 := "2"
	keyword2 := "33"
	keyword3 := "444"

	/*	curveParams := new(model.CurveParam)
		curveParams.Initialize()
		// 使用 CurveParam 实例进行操作，例如生成群元素、双线性配对等

		// 示例：生成一个新的 G1 群元素
		g1 := curveParams.GetNewG1()

		// 示例：生成一个新的 GT 群元素
		gt := curveParams.GetNewGT()

		// 示例：生成一个新的 Zn 群元素
		zn := curveParams.GetNewZn()

		// 示例：打印生成的群元素
		fmt.Printf("Generated G1 element: %v\n", g1)
		fmt.Printf("Generated GT element: %v\n", gt)
		fmt.Printf("Generated Zn element: %v\n", zn)*/

	masterKey, globalParams, err := model.Setup(256) // 使用256位参数，可以根据需要调整
	if err != nil {
		fmt.Println("Error in setup:", err)
		return
	}
	fmt.Println("Global Parameters:", globalParams)
	fmt.Println("Master Key:", masterKey)
	setupDuration := time.Since(startTime)
	fmt.Printf("Setup 运行时间：%s\n", setupDuration)
	//fmt.Println("Global Parameters:", globalParams)

	// 密钥生成1
	startTime = time.Now()
	Sxitong := []string{"cyf", "patient", "male", "female", "ten"} // 数据使用者属性集合
	fmt.Printf("系统属性：%s\n", Sxitong)
	S := []string{"cyf", "patient", "male"}
	fmt.Printf("cyf用户属性：%s\n", S) // 数据使用者属性集合
	P := []string{"('cyf', 'patient', 'male')", "('cyf', 'patient', 'ten')"}
	SAttr := []string{"cyf", "patient", "male"}
	fmt.Printf("访问策略：%s\n", P) // 数据使用者属性集合

	/*	DataOwner, err := model.NewDataOwner(globalParams.G)
		DataUser, err := model.NewDataUser(globalParams.G)
		PKDO := DataOwner.PKDO
		SKDO := DataOwner.PrivateKey
		PKDU := DataUser.PKDU
		SKDU := DataUser.PrivateKey*/

	// 密钥生成2
	// 创建两个不同的布隆过滤器实例
	bf1 := model.NewBloomFilter("filter1", 100, 3)
	bf2 := model.NewBloomFilter("filter2", 100, 3)
	bf1.BloomAdd("cyf")
	bf1.BloomAdd("patient")
	bf1.BloomAdd("male")
	bf1.BloomAdd("ten")
	bf2.BloomAdd("cyf")
	bf2.BloomAdd("patient")
	bf2.BloomAdd("male")
	fmt.Println("cyf 是否可能存在初始过滤器：", bf1.BloomContains("cyf"))         // 应该返回 true
	fmt.Println("male 是否可能存在初始过滤器：", bf1.BloomContains("male"))       // 应该返回 true
	fmt.Println("patient 是否可能存在初始过滤器：", bf1.BloomContains("patient")) // 应该返回 true
	fmt.Println("ten 是否可能存在初始过滤器：", bf1.BloomContains("ten"))         // 应该返回 true
	fmt.Println("cyp 是否可能存在隐藏过滤器：", bf2.BloomContains("cyp"))         // 应该返回 false
	fmt.Println("cyf是否可能存在隐藏过滤器：", bf2.BloomContains("cyf"))          // 应该返回 false
	fmt.Println("male 是否可能存在隐藏过滤器：", bf2.BloomContains("male"))       // 应该返回 true
	fmt.Println("patient 是否可能存在隐藏过滤器：", bf2.BloomContains("patient")) // 应该返回 true

	//fmt.Printf("布隆过滤器 %s", bf2.name)
	//hash = bf2.CollectHashValues()
	//fmt.Printf("布隆过滤器 %s 的已设置位的哈希值：%v\n", bf2.name, hash)
	//HideAttr := make([]uint64, 0)
	//	HideAttrbf2, err := model.CollectHashValues()

	// 密钥生成3

	/*	SK, err := model.KeyGen(globalParams, masterKey, S)
		if err != nil {
			fmt.Println("Error in key generation:", err)
			return
		}*/

	keyGenDuration := time.Since(startTime)
	fmt.Printf("KeyGen 运行时间：%s\n", keyGenDuration)

	// 加密
	startTime = time.Now()
	CTEnc := model.CTEnc(CTCK)
	if CTEnc == 1 {
		fmt.Println("加密成功！")
	} else {
		fmt.Println("加密失败。")
	}
	KW := "研究生二年级，网络安全，复旦大学"
	fmt.Printf("密文的关键词：%v\n", KW)

	/*	ck := model.GenerateCiphertextKey(globalParams.PK, SK, SStar)
		CT, sigma, I, Addr, err := model.Encrypt(KW, globalParams.PK, pkDO, S, ck, "plaintext")
		if err != nil {
			fmt.Println("Error in encryption:", err)
			return
		}*/

	encryptionDuration := time.Since(startTime)
	fmt.Printf("Encrypt 运行时间：%s\n", encryptionDuration)

	// 陷门生成
	startTime = time.Now()
	trapdoor := model.GenerateTrapdoor(CK)
	if trapdoor == 1 {
		fmt.Println("陷门生成成功！")
	} else {
		fmt.Println("陷门生成失败。")
	}
	KW2 := "研究生二年级，网络安全，复旦大学"
	fmt.Printf("医疗机构的关键词：%v\n", KW2)

	//KWPrime := []string{"bone", "tongue"}
	/*	TW, err := model.GenerateTrapdoor(KWPrime, PKDO, PKDU)
		if err != nil {
			fmt.Println("Error in trapdoor generation:", err)
			return
		}*/

	fmt.Println("您已成功生成陷门")
	trapdoorDuration := time.Since(startTime)
	fmt.Printf("GenerateTrapdoor 运行时间：%s\n", trapdoorDuration)

	// 搜索算法
	startTime = time.Now()
	result := model.Search(keyword1, keyword2, keyword3)
	searchDuration := time.Since(startTime)
	fmt.Printf("Search 运行时间：%s\n", searchDuration)

	if result == 1 {
		fmt.Println("搜索成功,密文已经发送！")
	} else {
		fmt.Println("搜索失败。")
	}

	// 解密算法

	/*
		startTime = time.Now()
		m, err := model.Decrypt(SK, CT, Addr)
		if err != nil {
			fmt.Println("Error in decryption:", err)
			return
		}
		decryptionDuration := time.Since(startTime)
		fmt.Printf("Decrypt 运行时间：%s\n", decryptionDuration)
		fmt.Printf("成功解密，解密数据是：%s\n", m)
	*/

	startTime = time.Now()
	m, err := model.Decrypt(SK, CT, Addr)
	if err != nil {
		fmt.Println("Error in decryption:", err)
		return
	}
	decryptionDuration := time.Since(startTime)
	fmt.Printf("Decrypt 运行时间：%s\n", decryptionDuration)
	fmt.Printf("成功解密，解密数据是：%s\n", m)
	/*
		// 撤销
		overlineS := []string{"attr1"}
		overlineUK, overlineSI, overlineCy, err := model.Revocation(globalParams.PK, pkDO, overlineS)
		if err != nil {
			fmt.Println("Error in revocation:", err)
			return
		}
	*/
	startTime = time.Now()
	revoAttr := "ten"
	W := model.Revocation(revoAttr)
	if W == 1 {
		fmt.Println("撤销成功,授权用户已经发送！")
	} else {
		fmt.Println("撤销失败。")
	}

	revocationDuration := time.Since(startTime)
	fmt.Printf("Revocation 运行时间：%s\n", revocationDuration)
	fmt.Printf("撤销属性：%s\n", revoAttr)
	fmt.Printf("撤销后访问策略属性：%s\n", SAttr)

	fmt.Println("系统运行成功！整个过程已经结束。")
	randomElement := model.GenerateRandomElement()

	fmt.Println("Random Element:", randomElement)

	/*	DataOwner, err := model.NewDataOwner(randomElement)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("撤销失败。")
		fmt.Println("DataOwner:")
		fmt.Println("Private Key:", DataOwner.PrivateKey)
		fmt.Println("Public Key (PKDO):", DataOwner.PKDO)*/
}
