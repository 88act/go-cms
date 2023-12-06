package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func Test_main() {
	//rsa 密钥文件产生
	// fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	// prvKey, pubKey := GenRsaKey()
	// fmt.Println(string(prvKey))
	// fmt.Println(string(pubKey))

	// fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	// var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	// fmt.Println("对消息进行签名操作...")
	// signData := RsaSignWithSha256([]byte(data), prvKey)
	// fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	// fmt.Println("\n对签名信息进行验证...")
	// if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
	// 	fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	// }

	// fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	// ciphertext := RsaEncrypt([]byte(data), pubKey)
	// fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	// sourceData := RsaDecrypt(ciphertext, prvKey)
	// fmt.Println("私钥解密后的数据：", string(sourceData))
}

// Generate RSA private/public key
func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publickey := &privatekey.PublicKey
	return privatekey, publickey, nil
}

// Dump private key into file
// This has same output as DumpPrivateKeyBuffer(), but dump to a file:
//
//	-----BEGIN RSA PRIVATE KEY-----
//	MIIEoQIBAAKCAQEAuql1lFYgKmKA1x5lQyadktbkeRRO0qrsmAkhvTtiz2p0Y+Ur
//	xWSYqDlmoY6vdkxj0Ex0z4zisoPnI+K89hV69O9v/83Yz7hYkLBHuwGiiSOiPZU7
//	...
//	PfKnburLQLE50wPkglfnGYfqQxtIiqn1hGTQO1xBxu03g+KM/Q==
//	-----END RSA PRIVATE KEY-----
func DumpPrivateKeyFile(privatekey *rsa.PrivateKey, filename string) error {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// Dump public key into file
//
//	-----BEGIN PUBLIC KEY-----
//	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2y8mEdCRE8siiI7udpge
//	5y1hrlSJzV7Xj0UojL/hi9u7s6TjYQQDA4M++/FezwkO5lBby2C+wK8bY7lgphuP
//	...
//	OZPrh/jItinhdzhyIXuYn6ohesPlM9i5TMpeBfpBmCwQQTfsAjBnXTTQzT4m4cmo
//	2QIDAQAB
//	-----END PUBLIC KEY-----
func DumpPublicKeyFile(publickey *rsa.PublicKey, filename string) error {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// RSA公钥私钥产生
func GenRsaKey() (prvkey []byte, pubkey []byte, publicKey *rsa.PublicKey) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	prvkey = pem.EncodeToMemory(block)
	publicKey = &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pubkey = pem.EncodeToMemory(block)
	return
}

// 签名
func RsaSignWithSha256(data []byte, keyBytes []byte) []byte {
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error"))
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		panic(err)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}

	return signature
}

// 验证
func RsaVerySignWithSha256(data, signData, keyBytes []byte) bool {
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	hashed := sha256.Sum256(data)
	//	fmt.Println("======hashed=========", base64.StdEncoding.EncodeToString(hashed))  //hex.EncodeToString(signData) //

	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		panic(err)
	}
	return true
}

// 公钥加密
func RsaEncrypt(data, keyBytes []byte) []byte {
	//解密pem格式的公钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, data)
	if err != nil {
		panic(err)
	}
	return ciphertext
}

// 私钥解密
func RsaDecrypt(ciphertext, keyBytes []byte) []byte {
	//获取私钥
	block, _ := pem.Decode(keyBytes)
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 解密
	data, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		panic(err)
	}
	return data
}

// Dump private key to base64 string
// Compared with DumpPrivateKeyBuffer this output:
//  1. Have no header/tailer line
//  2. Key content is merged into one-line format
//
// The output is:
//
//	MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2y8mEdCRE8siiI7udpge......2QIDAQAB
func DumpPrivateKeyBase64(privatekey *rsa.PrivateKey) (string, error) {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

func DumpPublicKeyBase64(publickey *rsa.PublicKey) (string, error) {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

// Load private key from base64
func LoadPrivateKeyBase64(base64key string) (*rsa.PrivateKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(keybytes)
	if err != nil {
		return nil, errors.New("parse private key error!")
	}

	return privatekey, nil
}

func LoadPublicKeyBase64(base64key string) (*rsa.PublicKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	pubkeyinterface, err := x509.ParsePKIXPublicKey(keybytes)
	if err != nil {
		return nil, err
	}

	publickey := pubkeyinterface.(*rsa.PublicKey)
	return publickey, nil
}

// Load private key from private key file
func LoadPrivateKeyFile(keyfile string) (*rsa.PrivateKey, error) {
	keybuffer, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(keybuffer))
	if block == nil {
		return nil, errors.New("private key error!")
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("parse private key error!")
	}

	return privatekey, nil
}

func LoadPublicKeyFile(keyfile string) (*rsa.PublicKey, error) {
	keybuffer, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keybuffer)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubkeyinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publickey := pubkeyinterface.(*rsa.PublicKey)
	return publickey, nil
}

// Dump private key to buffer.
func DumpPrivateKeyBuffer(privatekey *rsa.PrivateKey) (string, error) {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}

	var keybuffer []byte = pem.EncodeToMemory(block)
	return string(keybuffer), nil
}

func DumpPublicKeyBuffer(publickey *rsa.PublicKey) (string, error) {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keybytes,
	}

	var keybuffer []byte = pem.EncodeToMemory(block)
	return string(keybuffer), nil
}

// encrypt
func Encrypt(plaintext string, publickey *rsa.PublicKey) (string, error) {
	label := []byte("")
	sha256hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(sha256hash, rand.Reader, publickey, []byte(plaintext), label)

	decodedtext := base64.StdEncoding.EncodeToString(ciphertext)
	return decodedtext, err
}

// decrypt
func Decrypt(ciphertext string, privatekey *rsa.PrivateKey) (string, error) {
	decodedtext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	sha256hash := sha256.New()
	decryptedtext, err := rsa.DecryptOAEP(sha256hash, rand.Reader, privatekey, decodedtext, nil)
	if err != nil {
		return "", fmt.Errorf("RSA decrypt failed, error=%s\n", err.Error())
	}

	return string(decryptedtext), nil
}

// func main() {
//     // generate key
//     privatekey, publickey, err := GenerateKey()
//     if err != nil {
//         log.Fatalf("Cannot generate RSA key\n")
//     }

//     // dump private key to file
//     err = DumpPrivateKeyFile(privatekey, "private.pem")
//     if err != nil {
//         log.Fatalf("Cannot dump private key file\n")
//     }
//     // dump public key to file
//     err = DumpPublicKeyFile(publickey, "public.pem")
//     if err != nil {
//         log.Fatalf("Cannot dump public key file\n")
//     }

//     // encrypt message use public key
//     message := "abcd"
//     cipher, err := Encrypt(message, publickey)
//     if err != nil {
//         log.Fatalf("Cannot encrypt message\n")
//     }

//     // load private key
//     privatekey, err = LoadPrivateKeyFile("private.pem")
//     if privatekey == nil {
//         fmt.Printf("Cannot load private key\n");
//     }

//     // decrypt use private
//     plain, err := Decrypt(cipher, privatekey)
//     if err != nil {
//         log.Fatalf("Cannot decrypt message\n")
//     }
//     fmt.Printf("decrypt result is (%s)\n", plain)
// }
