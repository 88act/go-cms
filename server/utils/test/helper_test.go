package utils

import (
	"fmt"
	"testing"

	"github.com/golang-module/dongle"
)

func Test_IsEmpty(t *testing.T) {
	// var i int
	// var i2 int
	// i2 = 1
	// var s string = "xiaochuan"
	// var s2 string = ""
	// i3 := 0
	// var i4 *int
	// fmt.Println(IsEmpty(i4))
	// i4 = &i2
	// fmt.Println(IsEmpty(i4))
	// fmt.Println(IsEmpty(i3))
	// fmt.Println(IsEmpty(i))
	// fmt.Println(IsEmpty(i2))
	// fmt.Println(IsEmpty(s))
	// fmt.Println(IsEmpty(s2))

	// var arr [2]string
	// fmt.Println(IsEmpty(arr))

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)                // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)           // No、Zero、PKCS5、PKCS7
	cipher.SetKey("123456781234567812345678") // key 长度必须是 24
	cipher.SetIV("123456788")                 // iv 长度必须是 8

	str2 := dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
	fmt.Println("str2 ============= ")
	fmt.Println(str2)
	// 对字符串进行 3des 加密，输出经过 hex 编码的字符串
	str := dongle.Encrypt.FromString("hello12121221 world").By3Des(cipher).ToBase64String() // 0b2a92e81fb49ce1a43266aacaea7b81
	fmt.Println("str ============= ")
	fmt.Println(str)
	// 对经过 hex 编码的字符串进行 3des 解密，输出字符串
	res := dongle.Decrypt.FromBase64String(str).By3Des(cipher).ToString() // hello world
	fmt.Println("res ============= ")
	fmt.Println(res)

}

func Test_3DES(t *testing.T) {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)                // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)           // No、Zero、PKCS5、PKCS7
	cipher.SetKey("1dcrm4goRY8KODsgV1PPuHLB") // key must be 24 bytes
	cipher.SetIV("QCsJ2SKR")                  // iv must be 8 bytes

	// Encrypt by 3des from string and output string with hex encoding
	res := dongle.Encrypt.FromString("123456789|wqqweqw2321312312|2312312321321").By3Des(cipher).ToBase64String() // 0b2a92e81fb49ce1a43266aacaea7b81
	fmt.Println("res ============= ")
	fmt.Println(res)
	// Decrypt by des from string with hex encoding and output string
	res2 := dongle.Decrypt.FromBase64String(res).By3Des(cipher).ToString() // hello world

	fmt.Println("res2 ============= ")
	fmt.Println(res2)

}

func Test_AEC(t *testing.T) {

	// cipher := dongle.NewCipher()
	// cipher.SetMode(dongle.CBC)                // CBC、ECB、CFB、OFB、CTR、GCM
	// cipher.SetPadding(dongle.PKCS7)           // No、Zero、PKCS5、PKCS7
	// cipher.SetKey("123456781234567812345678") // key 长度必须是 24
	// cipher.SetIV("123456788")                 // iv 长度必须是 8

	// str2 := dongle.Encrypt.FromString("hello world").ByMd5().ToHexString() // 5eb63bbbe01eeed093cb22bb8f5acdc3
	// fmt.Println("str2 ============= ")
	// fmt.Println(str2)
	// // 对字符串进行 3des 加密，输出经过 hex 编码的字符串
	// str := dongle.Encrypt.FromString("hello12121221 world").By3Des(cipher).ToBase64String() // 0b2a92e81fb49ce1a43266aacaea7b81
	// fmt.Println("str ============= ")
	// fmt.Println(str)
	// // 对经过 hex 编码的字符串进行 3des 解密，输出字符串
	// res := dongle.Decrypt.FromBase64String(str).By3Des(cipher).ToString() // hello world
	// fmt.Println("res ============= ")
	// fmt.Println(res)

	cipher2 := dongle.NewCipher()
	cipher2.SetMode(dongle.CBC)        // CBC、ECB、CFB、OFB、CTR、GCM
	cipher2.SetPadding(dongle.PKCS7)   // No、Zero、PKCS5、PKCS7
	cipher2.SetKey("12345678AbcdEfgh") // key 长度必须是 16、24 或 32
	cipher2.SetIV("www.jiquxr.com88")  // iv 长度必须是 16、24 或 32

	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	str3 := dongle.Encrypt.FromString("hello worleqweqwe| weqweqwe| 2342d").ByAes(cipher2).ToHexString() // 65d823bdf1c581a1ded1cba42e03ca52
	fmt.Println("str3 ============= ")
	fmt.Println(str3)

	// 对经过 hex 编码的字符串进行 aes 解密，输出字符串
	str3333 := dongle.Decrypt.FromHexString(str3).ByAes(cipher2).ToString() // hello world
	fmt.Println("str3333 ============= ")
	fmt.Println(str3333)

	// cipher4 := dongle.NewCipher()
	// cipher4.SetMode(dongle.CBC)      // CBC、ECB、CFB、OFB、CTR、GCM
	// cipher4.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
	// cipher4.SetKey("12345678")       // key 长度必须是 8
	// cipher4.SetIV("123456788")       // iv 长度必须是 8

	// // 对字符串进行 des 加密，输出经过 hex 编码的字符串
	// str4 := dongle.Encrypt.FromString("hello world").ByDes(cipher4).ToHexString() // 0b2a92e81fb49ce1a43266aacaea7b81
	// // 对经过 hex 编码的字符串进行 des 解密，输出字符串
	// //dongle.Decrypt.FromHexString("0b2a92e81fb49ce1a43266aacaea7b81").ByDes(cipher).ToString() // hello world

	// fmt.Println("str4 ============= ")
	// fmt.Println(str4)

}
