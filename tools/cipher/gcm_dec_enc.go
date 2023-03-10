package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

func AESGCMEncrypt(plaintextStr string, key string, nonce string) (string, error) {
	// 将明文和密钥转换为字节切片
	plaintext := []byte(plaintextStr)
	keyByte := []byte(key)

	// 创建加密分组
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", fmt.Errorf("key 长度必须 16/24/32长度: %s", err.Error())
	}

	// 创建 GCM 模式的 AEAD
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 生成密文

	nonceByte, _ := base64.RawURLEncoding.DecodeString(nonce)
	ciphertext := aesgcm.Seal(nil, nonceByte, plaintext, nil)
	// 返回密文及随机数的 base64 编码
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

func AESGCMDecrypt(ciphertextStr string, key string, nonce string) (string, error) {
	// 将密文,密钥和生成的随机数转换为字节切片
	ciphertext, _ := base64.RawURLEncoding.DecodeString(ciphertextStr)
	nonceByte, _ := base64.RawURLEncoding.DecodeString(nonce)
	keyByte := []byte(key)

	// 创建加密分组
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	// 创建 GCM 模式的 AEAD
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 明文内容
	plaintext, err := aesgcm.Open(nil, nonceByte, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
