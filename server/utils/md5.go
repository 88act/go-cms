package utils

import (
	"crypto/md5"
	"encoding/hex"

	uuid "github.com/gofrs/uuid"
)

//@author: [88act](https://github.com/88act)
//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(b))
}

func GUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		//log.Fatalf("failed to generate UUID: %v", err)
	}
	var s = hex.EncodeToString(u2.Bytes())
	return s
}
func UUID() string {
	u2, err := uuid.NewV4()
	if err != nil {
		//log.Fatalf("failed to generate UUID: %v", err)
	}
	var s = hex.EncodeToString(u2.Bytes())
	return s

}
