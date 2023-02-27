package main

import (
	"flag"
	"fmt"
)

type flagStr struct {
	mode        string
	str         string
	key         string
	nonce       string
	cipher      string
	publicCert  string
	privateCert string
}

func GCMCipher(mode string, key string, nonce string, str string) (err error) {
	switch mode {
	case "encode":
		ciphertext, err := AESGCMEncrypt(str, key, nonce)
		if err != nil {
			return err
		}
		fmt.Printf("use GCM encode plain text: %s\n, cipher result: %s\n", str, ciphertext)
		return nil
	case "decode":
		plaintext, err := AESGCMDecrypt(str, key, nonce)
		if err != nil {
			return err
		}
		fmt.Printf("use GCM decode cipher text: %s\n, plain result: %s\n", str, plaintext)
		return nil
	default:
		fmt.Printf("unavailable mode: %s\n", mode)
		return nil
	}
}

func RSACipher(mode string, publicCert string, privateCert string, str string) (err error) {
	switch mode {
	case "encode":
		ciphertext, err := RSA_Encrypt(str, publicCert)
		if err != nil {
			return err
		}
		fmt.Printf("use RSA encode plain text: %s\n, cipehr result: %s\n", str, ciphertext)
		return nil
	case "decode":
		plaintext, err := RSA_Decrypt(str, privateCert)
		if err != nil {
			return err
		}
		fmt.Printf("use RSA decode cipher text: %s\n, plain result: %s\n", str, plaintext)
		return nil
	default:
		fmt.Printf("unavailable mode: %s\n", mode)
		return nil
	}
}

func main() {
	// key   := "mTqBnYoyhQm8xOkRFkaUU7X2"
	// nonce := "0xCWIOTotifzJMnD"
	flagParse := &flagStr{}
	flag.StringVar(&flagParse.mode, "mode", "", "gcm or rsa cipher mode, support encode or decode")
	flag.StringVar(&flagParse.str, "str", "", "str need to encode or decode")
	flag.StringVar(&flagParse.key, "key", "", "gcm key")
	flag.StringVar(&flagParse.nonce, "nonce", "", "gcm nonce")
	flag.StringVar(&flagParse.cipher, "cipher-method", "", "cipher method support GCM or RSA")
	flag.StringVar(&flagParse.publicCert, "public-cert", "", "RSA public cert")
	flag.StringVar(&flagParse.privateCert, "private-cert", "", "RSA private cert")

	flag.Parse()
	switch flagParse.cipher {
	case "GCM":
		err := GCMCipher(flagParse.mode, flagParse.key, flagParse.nonce, flagParse.str)
		if err != nil {
			fmt.Printf("GCM cipher error: %+v\n", err)
		}
	case "RSA":
		err := RSACipher(flagParse.mode, flagParse.publicCert, flagParse.privateCert, flagParse.str)
		if err != nil {
			fmt.Printf("RSA cipher error: %+v\n", err)
		}
	default:
		fmt.Printf("GCM/RSA加解密说明\n")
		fmt.Printf("使用RSA加密，需传-public-cert；使用RSA解密，需传-private-cert\n")
		fmt.Printf("使用GCM加密或解密，需传 -key 与 -nonce\n")
		fmt.Printf("-mode: encode or decode\n")
		fmt.Printf("-str: plain or cipher text\n")
		fmt.Printf("-key: gcm key\n")
		fmt.Printf("-nonce: gcm nonce\n")
		fmt.Printf("-cipher-method: cipher method, support GCM or RSA\n")
		fmt.Printf("-public-cert: RSA public cert path\n")
		fmt.Printf("-private-cert: RSA private cert path\n")
	}
}
