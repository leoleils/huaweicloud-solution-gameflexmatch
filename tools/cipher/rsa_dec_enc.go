package main
import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"os"
	"encoding/pem"
	"encoding/base64"
	"fmt"
)

// RSA加密，使用标准的base64编码输出
func RSA_Encrypt(plaintextStr string, path string) (string, error) {
	// 打开文件
	plaintext := []byte(plaintextStr)
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open public key file failed: %s, %s", path, err.Error())
	}
	defer file.Close()
	// 读取文件的内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	// pem解码
	block, _ := pem.Decode(buf)
	// x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("parse public key file failed: %s, %s", path, err.Error())
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plaintext)
	if err != nil {
		return "", fmt.Errorf("encode message using public key file failed: %s", err.Error())
	}
	//返回密文
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

//RSA解密
func RSA_Decrypt(cipherTextStr string, path string) (string, error) {
	//打开文件
	ciphertext, err := base64.StdEncoding.DecodeString(cipherTextStr)

	if err != nil {
		return "", fmt.Errorf("decode std base64 message failed: %s", err.Error())
	}
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("open private key file failed: %s, %s", path, err.Error())
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		
		return "", fmt.Errorf("parse private key file failed: %s, %s", path, err.Error())
	}
	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode message using private key file failed: %s", err.Error())
	}
	//返回明文
	return string(plainText), nil
}