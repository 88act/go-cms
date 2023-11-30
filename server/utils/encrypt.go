package utils

import "github.com/golang-module/dongle"

//3des  加密 ToBase64String
func Des3Encode(str string, key string) string {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
	cipher.SetKey(key)              // key must be 24 bytes
	cipher.SetIV("jiquxr88")        // iv must be 8 bytes
	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	return dongle.Encrypt.FromString(str).By3Des(cipher).ToBase64String() // 65d823bdf1c581a1ded1cba42e03ca52
}

//3des  解密 ToBase64String
func Des3Decode(str string, key string) string {
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7) // No、Zero、PKCS5、PKCS7
	cipher.SetKey(key)              // key must be 24 bytes
	cipher.SetIV("jiquxr88")        // iv must be 8 bytes
	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	return dongle.Decrypt.FromBase64String(str).By3Des(cipher).ToString()
}

//aec  加密
func AecEncode(str string, key string) string {

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)       // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)  // No、Zero、PKCS5、PKCS7
	cipher.SetKey(key)               //"12345678AbcdEfgh"  // key 长度必须是 16、24 或 32
	cipher.SetIV("www.jiquxr.com88") // iv 长度必须是 16、24 或 32
	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	return dongle.Encrypt.FromString(str).ByAes(cipher).ToBase64String() // 65d823bdf1c581a1ded1cba42e03ca52

}

//aec  解密
func AecDecode(str string, key string) string {

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)       // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)  // No、Zero、PKCS5、PKCS7
	cipher.SetKey(key)               //"12345678AbcdEfgh"  // key 长度必须是 16、24 或 32
	cipher.SetIV("www.jiquxr.com88") // iv 长度必须是 16、24 或 32
	// 对字符串进行 aes 加密，输出经过 hex 编码的字符串
	return dongle.Decrypt.FromBase64String(str).ByAes(cipher).ToString()
}
