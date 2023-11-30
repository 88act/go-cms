package utils

import (
	"testing"

	"github.com/golang-module/dongle"
)

func Test_3DES(t *testing.T) {
	str := "devcode11111122222|APP123|2022-06-29 10:21:49"
	strEncode := Des3Encode(str, "www.jiquxr88.com12345678")
	t.Log(str)
	t.Log(strEncode)

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)                // CBC、ECB、CFB、OFB、CTR、GCM
	cipher.SetPadding(dongle.PKCS7)           // No、Zero、PKCS5、PKCS7
	cipher.SetKey("www.jiquxr88.com12345678") // key must be 24 bytes
	cipher.SetIV("QCsJ2SKR")                  // iv must be 8 bytes

	// Encrypt by 3des from string and output string with hex encoding
	res := dongle.Encrypt.FromString("devcode11111122222|APP123|2022-06-29 10:21:49").By3Des(cipher).ToBase64String() // 0b2a92e81fb49ce1a43266aacaea7b81
	t.Log("res ============= ")
	t.Log(res)
	// Decrypt by des from string with hex encoding and output string
	// res2 := dongle.Decrypt.FromBase64String(res).By3Des(cipher).ToString() // hello world

	// fmt.Println("res2 ============= ")
	// fmt.Println(res2)

}
