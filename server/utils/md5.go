package utils

import (
	"crypto/md5"
	"encoding/hex"

	uuid "github.com/gofrs/uuid"
	"github.com/jaevor/go-nanoid"
)

//@author:  [linjd] 10512203@qq.com
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

// 32位
func GUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		//log.Fatalf("failed to generate UUID: %v", err)
	}
	var s = hex.EncodeToString(u2.Bytes())
	return s
}

// 32位
func UUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		//log.Fatalf("failed to generate UUID: %v", err)
	}
	var s = hex.EncodeToString(u2.Bytes())
	return s
}

// f5394eef-e576-4709-9e4b-a7c231bd34a4
func UUID36() string {
	u2, _ := uuid.NewV4()
	return u2.String()

}

// nanoid
func Nanoid(len int) string {
	//str := "0123456789abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := "0123456789abcdefghijklmnopqrstuvwsyz"
	nid, _ := nanoid.CustomASCII(str, len)
	return nid()
}
func NanoidFull(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//str := "0123456789abcdefghijklmnopqrstuvwsyz"
	nid, _ := nanoid.CustomASCII(str, len)
	return nid()
}
