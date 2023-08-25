/*package main

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

	//lambda := uint32(256)

	//PK, MK := model.Setup(lambda)

	//fmt.Println("PK:", PK)
	//fmt.Println("MK:", MK)
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

	// Define some sample attributes and policies
	//S := []string{"attr1", "attr2", "attr3"}
	//P := []string{"attr1", "attr2"}

	//	userKeys := model.KeyGen(PK, MK, S, P)

	//	fmt.Println("User SK:", userKeys.SK)
	//	fmt.Println("User PkDO:", userKeys.PkDO)
	//	fmt.Println("User PkDU:", userKeys.PkDU)
	//	fmt.Println("User GBFA:", userKeys.GBFA)
	//	fmt.Println("User SCirc:", userKeys.SCirc)

	// 密钥生成1
	//	startTime = time.Now()
	//	Sxitong := []string{"cyf", "patient", "male", "female", "ten"} // 数据使用者属性集合
	//	fmt.Printf("系统属性：%s\n", Sxitong)
	//	S := []string{"cyf", "patient", "male"}
	//	fmt.Printf("cyf用户属性：%s\n", S) // 数据使用者属性集合
	//	P := []string{"('cyf', 'patient', 'male')", "('cyf', 'patient', 'ten')"}
	//	SAttr := []string{"cyf", "patient", "male"}
	//	fmt.Printf("访问策略：%s\n", P) // 数据使用者属性集合

	//	DataOwner, err := model.NewDataOwner(globalParams.G)
	//	DataUser, err := model.NewDataUser(globalParams.G)
	//	PKDO := DataOwner.PKDO
	//	SKDO := DataOwner.PrivateKey
	//	PKDU := DataUser.PKDU
	//	SKDU := DataUser.PrivateKey

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

	//	SK, err := model.KeyGen(globalParams, masterKey, S)
	//	if err != nil {
	//		fmt.Println("Error in key generation:", err)
	//	}

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

	//	ck := model.GenerateCiphertextKey(globalParams.PK, SK, SStar)
	//	CT, sigma, I, Addr, err := model.Encrypt(KW, globalParams.PK, pkDO, S, ck, "plaintext")
	//	if err != nil {
	//		fmt.Println("Error in encryption:", err)
	//		return
	//	}

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
	//	TW, err := model.GenerateTrapdoor(KWPrime, PKDO, PKDU)
	//	if err != nil {
	//		fmt.Println("Error in trapdoor generation:", err)
	//		return
	//	}

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

	//	startTime = time.Now()
	//	m, err := model.Decrypt(SK, CT, Addr)
	//	if err != nil {
	//		fmt.Println("Error in decryption:", err)
	//		return
	//	}
	//	decryptionDuration := time.Since(startTime)
	//	fmt.Printf("Decrypt 运行时间：%s\n", decryptionDuration)
	//	fmt.Printf("成功解密，解密数据是：%s\n", m)

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
	//	overlineS := []string{"attr1"}
	//	overlineUK, overlineSI, overlineCy, err := model.Revocation(globalParams.PK, pkDO, overlineS)
	//	if err != nil {
	//		fmt.Println("Error in revocation:", err)
	//		return
	//	}

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
	//	fmt.Printf("撤销后访问策略属性：%s\n", SAttr)

	fmt.Println("系统运行成功！整个过程已经结束。")
	randomElement := model.GenerateRandomElement()

	fmt.Println("Random Element:", randomElement)

	//	DataOwner, err := model.NewDataOwner(randomElement)
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		return
	//	}
	//	fmt.Println("撤销失败。")
	//	fmt.Println("DataOwner:")
	//	fmt.Println("Private Key:", DataOwner.PrivateKey)
	//	fmt.Println("Public Key (PKDO):", DataOwner.PKDO)
	fmt.Println("撤销失败。")
}
*/

// main.go
package main

import (
	"fmt"
	"github.com/Nik-U/pbc"
	"github.com/cyfpps/ABSE-Blockchain/MODEL"
	"os"
	"time"
)

func main() {

	lambda := uint32(256)
	startTime := time.Now() // 记录开始时间
	PK, MK := MODEL.Setup(lambda)
	elapsedTime := time.Since(startTime) // 计算经过的时间
	fmt.Printf("Setup took %s\n", elapsedTime)
	fmt.Println("PK:", PK)
	fmt.Println("MK:", MK)
	fmt.Println("PK.G", PK.G)
	// 定义属性集和访问策略
	S := []string{"attr1", "attr2", "attr3"}
	P := []string{"attr1", " attr2", "attr4", "attr3"}

	// 密钥生成
	startTime = time.Now()
	userKeys := MODEL.KeyGen(PK, MK, S, P)
	elapsedTime = time.Since(startTime)
	fmt.Printf("KeyGen took %s\n", elapsedTime)

	// 数据拥有者的私钥
	doPrivateKey := userKeys.SkDO
	duPrivateKey := userKeys.SkDU

	// 打印结果
	fmt.Println("Data Owner Private Key (DO私钥):", doPrivateKey)
	fmt.Println("Data User Private Key (DU私钥):", duPrivateKey)
	fmt.Println("User SK:", userKeys.SK)
	fmt.Println("User PkDO:", userKeys.PkDO)
	fmt.Println("User PkDU:", userKeys.PkDU)
	fmt.Println("User GBFA:", userKeys.GBFA)
	fmt.Println("User SCirc:", userKeys.SCirc)
	fmt.Println("User ck:", MK.Ck)

	/*bf := userKeys.GBFA
	contains := bf.Contains("attr1")
	contain2 := bf.Contains("attr3")
	fmt.Println("BloomFilter contains 'attr1':", contains)
	fmt.Println("BloomFilter contains 'attr3':", contain2)
	*/
	for _, policy := range P {
		root := MODEL.ParsePolicy(policy)
		MODEL.PrintTree(root, 0)
		fmt.Println("Satisfies policy:", MODEL.Satisfies(root, S))
		fmt.Println("------")
	}

	params := pbc.GenerateA(160, 512)
	pairing := params.NewPairing()

	// 创建一个随机的ck
	ck2 := pairing.NewGT().Rand()

	// 调用CTEncrypt函数
	ct := MODEL.CTEncrypt(PK, ck2, MK)
	// 打印结果
	fmt.Println("C:", ct.C)
	fmt.Println("C0:", ct.C0)
	fmt.Println("C1:", ct.C1)
	// 如果您在后续还要使用Cy，也可以打印它
	fmt.Println("Cy:", ct.Cy)

	// 示例CT
	CT := "examplerrrrrCT"

	// 调用SigEncrypt函数
	startTime = time.Now()
	Sigma := MODEL.SigEncrypt(PK, CT, doPrivateKey)
	elapsedTime = time.Since(startTime)
	fmt.Printf("SigEncrypt took %s\n", elapsedTime)

	// 打印结果
	fmt.Println("Sigma1:", Sigma.Sigma1)
	fmt.Println("Sigma2:", Sigma.Sigma2)

	KW := []string{"keyword1", "keyword2", "keyword3"}
	// 测量IndexGen函数的运行时间
	startTime = time.Now()
	Index := MODEL.IndexGen(KW, PK, doPrivateKey)
	elapsedTime = time.Since(startTime)
	fmt.Printf("IndexGen took %s\n", elapsedTime)
	fmt.Println("I1:", Index.I1)
	fmt.Println("I2:", Index.I2)
	fmt.Println("I3:", Index.I3)

	// 测量Trapdoor函数的运行时间
	startTime = time.Now()
	TW := MODEL.Trapdoor(KW, PK, duPrivateKey, userKeys.PkDO)
	elapsedTime = time.Since(startTime)
	fmt.Printf("Trapdoor took %s\n", elapsedTime)
	fmt.Println(TW.TW2, Index.I1, Index.I2, TW.TW1)
	fmt.Println("TW1:", TW.TW1)
	fmt.Println("TW2:", TW.TW2)
	fmt.Println("TW3:", TW.TW3)

	//进行撤销
	revokedAttrs := []string{"attr4"}
	startTime = time.Now()
	userKeys = MODEL.Revocation(userKeys, revokedAttrs)
	elapsedTime = time.Since(startTime)
	fmt.Printf("Revocation took %s\n", elapsedTime)
	fmt.Println("User GBFA:", userKeys.GBFA)

	// 5. 再次验证属性是否在布隆过滤器中

	fmt.Println("\nAfter revocation:")
	for _, attr := range S {
		if userKeys.GBFA.Contains(attr) {
			fmt.Printf("%s is in the Bloom Filter\n", attr)
		} else {
			fmt.Printf("%s is NOT in the Bloom Filter\n", attr)
			fmt.Println("撤销成功")
		}
	}
	startTime = time.Now()
	result := MODEL.Search(Index, TW, userKeys.SCirc, userKeys.GBFA, PK, Sigma, userKeys)
	elapsedTime = time.Since(startTime)
	fmt.Printf("Search took %s\n", elapsedTime)

	// 5. 打印结果
	fmt.Println("搜索结果:", result)
	if !result {
		fmt.Println("搜索失败，程序退出。")
		os.Exit(1) // 退出程序，其中1是退出码，表示程序因为某种错误而终止
	}

	startTime = time.Now()
	CT1, SKPrime := MODEL.ProDec(userKeys, ct, PK, MK)
	elapsedTime = time.Since(startTime)
	fmt.Printf("ProDec took %s\n", elapsedTime)

	// 使用 DecryptDU 函数进行用户解密
	startTime = time.Now()
	f := PK.G.Pairing().NewZr().Rand()
	ck := MODEL.DecryptDU(f, CT1, PK, MK)
	elapsedTime = time.Since(startTime)
	fmt.Printf("DecryptDU took %s\n", elapsedTime)

	// 打印 CK
	fmt.Println("您已经成功解密CK:", ck, SKPrime)

}
