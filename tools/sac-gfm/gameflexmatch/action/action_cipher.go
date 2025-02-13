package action

import (
	"crypto/aes"
	"crypto/cipher"
	cryrand "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"solutionbuild/gameflexmatch/common"

	"github.com/urfave/cli/v2"
)

type CipherAction struct {
	c *cli.Context
}

const (
	akKey = "AK"
	skKey = "SK"
	fmRedisPwdKey = "FleetManagerRedisPassword"
	fmDbPwdKey = "FleetManagerDBPassword"
	agRedisPwdKey = "AppGatewayRedisPassword"
	agDbPwdKey = "AppGatewayDBPassword"
	aassRedisPwdKey = "AASSRedisPassword"
	aassDbPwdKey = "AASSDBPassword"
	influxDBPwdKey = "InfluxDBPassword"
)

const (
	Create = "create"
	Mode = "mode"
	Method = "method"
	Text = "text"
	InitConf = "init-conf"
)

const (
	jwtKeyLen = 48
	gcmKeyLen = 24
	gcmNonceLen = 16
	rsaBits = 2048
)

func NewCipherAction(c *cli.Context) *CipherAction {
	return &CipherAction{
		c: c,
	}
}

// Create GCM or RSA info
func (a *CipherAction) Create(initConf string) error {
	create := a.c.String(Create)
	switch create {
	case "gcm":
		return a.createGcm(initConf)

	case "rsa":
		return a.createRsa(initConf)

	case "tls":
		return a.createTls(initConf)
	case "jwt": 
		return a.createJwt(initConf)
	default:
		fmt.Printf("unavailable method: %s\n", create)
		fmt.Print("create gcm key and nonce or create rsa key file, defalut using auth configure file, options: [gcm, rsa]")
		return fmt.Errorf("unavailable method: %s", create)
	}
}

func (a *CipherAction) createGcm(initConf string) error {
	key := common.StringWithCharset(gcmKeyLen)
	fmt.Printf("new gcm key: %s", key)
	// 修改key
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherGcmKey, key), initConf)); err != nil {
		return err
	}
	nonce := common.StringWithCharset(gcmNonceLen)
	fmt.Printf("new gcm nonce: %s", nonce)
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherGcmNonce, nonce), initConf)); err != nil {
		return err
	}
	return nil
}

func (a *CipherAction) createJwt(initConf string) error {
	key := common.StringWithCharset(jwtKeyLen)
	fmt.Printf("new jwt key: %s", key)
	// 修改key
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.FleetmanagerJwtKey, key), initConf)); err != nil {
		return err
	}
	return nil
}

func (a *CipherAction) createRsa(initConf string) error {
	absDir, err := filepath.Abs(initConf)
	if err != nil {
		return err
	}
	parentDir := filepath.Dir(absDir)
	if err := a.generateRSAKey(parentDir); err != nil {
		return err
	}
	fmt.Printf("succeed to generate new rsa private key and public key into %s", parentDir)
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherRsaPublicKey, parentDir + "/public.pem"), initConf)); err != nil {
		return err
	}
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherRsaPrivateKey, parentDir + "/private.pem"), initConf)); err != nil {
		return err
	}
	return nil
}

func (a *CipherAction) createTls(initConf string) error {
	absDir, err := filepath.Abs(initConf)
	if err != nil {
		return err
	}
	parentDir := filepath.Dir(absDir)
	if err := a.generateTlsKey(parentDir); err != nil {
		return err
	}
	fmt.Printf("succeed to generate new tls.key / tls.csr / tls.crt into %s", parentDir)
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherTlsKey, parentDir + "/tls.key"), initConf)); err != nil {
		return err
	}
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherTlsCsr, parentDir + "/tls.csr"), initConf)); err != nil {
		return err
	}
	if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(common.CipherTlsCrt, parentDir + "/tls.crt"), initConf)); err != nil {
		return err
	}
	return nil
}

func (a *CipherAction) GcmEncode(initConf string) error {
	text := a.c.String(Text)
	key, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherGcmKey)
	if err != nil {
		return err
	}
	nonce, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherGcmNonce)
	if err != nil {
		return err
	}
	fmt.Print("cipher encode use gcm\n")
	fmt.Printf("gcm key: %s\n", key)
	fmt.Printf("gcm nonce: %s\n", nonce)
	if text == "" {
		return fmt.Errorf("you must provide the text that would encode")
	}
	cipherText, err := a.AESGCMEncrypt(text, key, nonce)
	if err != nil {
		return err
	}
	fmt.Printf("cipher encode result: %s for %s \n", cipherText, text)
	return nil
}

func (a *CipherAction) GcmDecode(initConf string) error {
	text := a.c.String(Text)
	key, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherGcmKey)
	if err != nil {
		return err
	}
	nonce, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherGcmNonce)
	if err != nil {
		return err
	}
	fmt.Print("cipher decode use gcm\n")
	fmt.Printf("gcm key: %s\n", key)
	fmt.Printf("gcm nonce: %s\n", nonce)
	if text == "" {
		return fmt.Errorf("you must provide the text that would decode")
	}
	cipherText, err := a.AESGCMDecrypt(text, key, nonce)
	if err != nil {
		return err
	}
	fmt.Printf("cipher decode result: %s for %s \n", cipherText, text)
	return nil
}

func (a *CipherAction) RsaEncode(initConf string) error {
	text := a.c.String(Text)
	publicKey, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherRsaPublicKey)
	if err != nil {
		return err
	}
	fmt.Printf("cipher encode use rsa public key: %s\n", publicKey)
	if text != "" {
		cipherText, err := a.RSAEncrypt(text, publicKey)
		if err != nil {
			return err
		}
		fmt.Printf("cipher result: %s for %s \n", cipherText, text)
		return nil
	}
	cipherMap := a.getKeyInfo()

	for section, cipherList := range(cipherMap) {
		for _, textKey := range(cipherList) {
			textValue, err := common.GetConfKey(initConf, section, textKey)
			if err != nil {
				return err
			}
			if textValue != "" {
				cipherText, err := a.RSAEncrypt(textValue, publicKey)
				if err != nil {
					return err
				}
				fmt.Printf("[cipher result] %s : %s for %s \n", textKey, cipherText, textValue)
				if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(textKey, cipherText), initConf)); err != nil {
					return err
				}
				continue
			}
			fmt.Printf("the text key: %s is none.\n", textKey)
		}
	}
	
	return nil
}

func (a *CipherAction) RsaDecode(initConf string) error {
	text := a.c.String(Text)
	privateKey, err := common.GetConfKey(initConf, common.SectionCipher, common.CipherRsaPrivateKey)
	if err != nil {
		return err
	}
	fmt.Printf("cipher decode use rsa private key: %s\n", privateKey)
	if text != "" {
		cipherText, err := a.RSADecrypt(text, privateKey)
		if err != nil {
			return err
		}
		fmt.Printf("cipher result: %s for %s \n", cipherText, text)
		return nil
	}
	cipherMap := a.getKeyInfo()

	for section, cipherList := range(cipherMap) {
		for _, textKey := range(cipherList) {
			textValue, err := common.GetConfKey(initConf, section, textKey)
			if err != nil {
				return err
			}
			if textValue != "" {
				cipherText, err := a.RSADecrypt(textValue, privateKey)
				if err != nil {
					return err
				}
				fmt.Printf("[cipher result] %s : %s for %s \n", textKey, cipherText, textValue)
				if err := common.ExecCmd(exec.Command("sed", "-i", common.SedSprintf(textKey, cipherText), initConf)); err != nil {
					return err
				}
				continue
			}
			fmt.Printf("the text key: %s is none.\n", textKey)
		}
	}
	
	return nil
}

func (a *CipherAction) generateTlsKey(path string) error {
	// 生成tls.key
	key, err := rsa.GenerateKey(cryrand.Reader, 3072)
	if err != nil {
		return err
	}
	keyFile, err := os.Create(path + "/tls.key")
	if err != nil {
		return err
	}
	defer keyFile.Close()
	keyBytes := x509.MarshalPKCS1PrivateKey(key)
	pemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}
	err = pem.Encode(keyFile, pemBlock)
	if err != nil {
		return err
	}
	// 生成tls.csr
	template := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName: "GameFlexMatch",
		},
	}
	csrBytes, err := x509.CreateCertificateRequest(cryrand.Reader, &template, key)
	if err != nil {
		return err
	}
	csrFile, err := os.Create(path + "/tls.csr")
	if err != nil {
		return err
	}
	defer csrFile.Close()
	pem.Encode(csrFile, &pem.Block{
		Type:  "CERTIFICATE REQUEST",
		Bytes: csrBytes,
	})

	// 生成tls.crt
	templateCrt := x509.Certificate{
        SerialNumber: big.NewInt(1658),
        Subject: pkix.Name{
            CommonName:   template.Subject.CommonName,
            Organization: template.Subject.Organization,
        },
        NotBefore:             time.Now(),
        NotAfter:              time.Now().AddDate(1, 0, 0),
        BasicConstraintsValid: true,
        IsCA:                  false,
        KeyUsage: x509.KeyUsageDigitalSignature |
            x509.KeyUsageKeyEncipherment |
            x509.KeyUsageCertSign,
        ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    }
	certificate, err := x509.CreateCertificate(cryrand.Reader, &templateCrt, &templateCrt, &key.PublicKey, key)
	if err != nil {
		return err
	}
	certificateFile, err := os.Create(path + "/tls.crt")
    if err != nil {
       return err
    }
	pem.Encode(certificateFile, &pem.Block{Type: "CERTIFICATE", Bytes: certificate})
    certificateFile.Close()
	return nil
}

//生成RSA私钥和公钥，保存到文件中
func (a *CipherAction) generateRSAKey(path string) error {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(cryrand.Reader, rsaBits)
	if err != nil{
		return err
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create(path + "/private.pem")
	if err!=nil{
		return err
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock:= pem.Block{Type: "RSA Private Key",Bytes:X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey:=privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey,err:=x509.MarshalPKIXPublicKey(&publicKey)
	if err!=nil{
		return err
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create(path + "/public.pem")
	if err!=nil{
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock:= pem.Block{Type: "RSA Public Key",Bytes:X509PublicKey}
	//保存到文件
	pem.Encode(publicFile,&publicBlock)
	return nil
}

//RSA加密
// RSA加密，使用标准的base64编码输出
func (a *CipherAction) RSAEncrypt(plaintextStr string, path string) (string, error) {
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
	cipherText, err := rsa.EncryptPKCS1v15(cryrand.Reader, publicKey, plaintext)
	if err != nil {
		return "", fmt.Errorf("encode message using public key file failed: %s", err.Error())
	}
	//返回密文
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

//RSA解密
func (a *CipherAction) RSADecrypt(cipherTextStr string, path string) (string, error) {
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
	plainText, err := rsa.DecryptPKCS1v15(cryrand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", fmt.Errorf("decode message using private key file failed: %s", err.Error())
	}
	//返回明文
	return string(plainText), nil
}

func (a *CipherAction) AESGCMEncrypt(plaintextStr string, key string, nonce string) (string, error) {
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

func (a *CipherAction) AESGCMDecrypt(ciphertextStr string, key string, nonce string) (string, error) {
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

func (a *CipherAction) getKeyInfo() map[string][]string {
	res := make(map[string][]string)
	res[common.SectionCloud] = []string{
		akKey,
		skKey,
	}
	res[common.SectionFleetManager] = []string{
		fmDbPwdKey,
		fmRedisPwdKey,
	}
	res[common.SectionAppGateway] = []string {
		agDbPwdKey,
		agRedisPwdKey,
	}
	res[common.SectionAASS] = []string {
		aassDbPwdKey,
		aassRedisPwdKey,
	}
	res[common.SectionInfluxDB] = []string {
		influxDBPwdKey,
	}
	return res
}
