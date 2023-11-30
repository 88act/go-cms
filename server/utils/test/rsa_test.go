package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	//rsa 密钥文件产生
	fmt.Println("-------------------------------获取RSA公私钥-----------------------------------------")
	prvKey, pubKey := GenRsaKey()
	fmt.Println(string(prvKey))
	fmt.Println(string(pubKey))

	fmt.Println("-------------------------------进行签名与验证操作-----------------------------------------")
	var data = "卧了个槽，这么神奇的吗？？！！！  ԅ(¯﹃¯ԅ) ！！！！！！）"
	fmt.Println("对消息进行签名操作...")
	signData := RsaSignWithSha256([]byte(data), prvKey)
	fmt.Println("消息的签名信息： ", hex.EncodeToString(signData))
	fmt.Println("\n对签名信息进行验证...")
	if RsaVerySignWithSha256([]byte(data), signData, pubKey) {
		fmt.Println("签名信息验证成功，确定是正确私钥签名！！")
	}

	fmt.Println("-------------------------------进行加密解密操作-----------------------------------------")
	ciphertext := RsaEncrypt([]byte(data), pubKey)
	fmt.Println("公钥加密后的数据：", hex.EncodeToString(ciphertext))
	sourceData := RsaDecrypt(ciphertext, prvKey)
	fmt.Println("私钥解密后的数据：", string(sourceData))
}

// RSA公钥私钥产生
func GenRsaKey() (prvkey, pubkey []byte) {
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
	publicKey := &privateKey.PublicKey
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
