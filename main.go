package main

import (
	"fmt"
	"time"

	"github.com/cyfpps/ABSE-Blockchain/model"
)

func main() {
	startTime := time.Now()

	// Setup
	pk, mk, err := model.Setup(256)
	if err != nil {
		fmt.Println("Setup error:", err)
		return
	}

	// Generate user keys and attributes
	attributes := []string{"attr1", "attr2", "attr3"}
	sk, pkDO, pkDU, Si, err := model.KeyGen(pk, mk, attributes)
	if err != nil {
		fmt.Println("KeyGen error:", err)
		return
	}

	// Encrypt
	keyword := "keyword"
	content := []byte("Hello, world!")
	ct, sig, index, addr, err := model.Encrypt(pk, pkDO, nil, keyword, content)
	if err != nil {
		fmt.Println("Encrypt error:", err)
		return
	}

	// Trapdoor
	tw, err := model.Trapdoor(pkDO, sk, []string{keyword})
	if err != nil {
		fmt.Println("Trapdoor error:", err)
		return
	}

	// Search
	searchStartTime := time.Now()
	results, err := model.Search(pk, tw, Si)
	if err != nil {
		fmt.Println("Search error:", err)
		return
	}
	searchDuration := time.Since(searchStartTime)

	// Decrypt
	decryptStartTime := time.Now()
	decryptedContent, err := model.Decrypt(results.DecryptionKey, ct, addr)
	if err != nil {
		fmt.Println("Decrypt error:", err)
		return
	}
	decryptDuration := time.Since(decryptStartTime)

	fmt.Println("运行成功！")
	fmt.Printf("成功解密，解密数据是：%s\n", decryptedContent)
	fmt.Printf("Search 运行时间：%s\n", searchDuration)
	fmt.Printf("Decrypt 运行时间：%s\n", decryptDuration)

	totalDuration := time.Since(startTime)
	fmt.Printf("总运行时间xiugai%s\n", totalDuration)
}
