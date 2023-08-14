// model/encryption.go

package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// EncryptFile 使用AES算法对提供的文件内容进行加密，并返回密文和地址。
func EncryptFile(fileContent []byte) ([]byte, string, error) {
	// 生成一个随机的加密密钥
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return nil, "", err
	}

	// 使用密钥创建新的AES密码
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, "", err
	}

	// 创建一个新的GCM密码
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, "", err
	}

	// 生成一个随机数（nonce）
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, "", err
	}

	// 对文件内容进行加密
	ciphertext := gcm.Seal(nil, nonce, fileContent, nil)

	// 生成用于存储加密文件的地址
	address := generateAddress(ciphertext)

	return ciphertext, address, nil
}

// generateAddress 根据内容生成一个地址。
func generateAddress(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}
