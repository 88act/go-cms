package uuid

import (
	"github.com/jaevor/go-nanoid"
	"github.com/zeromicro/go-zero/core/logx"
)

// len 长度 , 大于10 保障 1000万不重复
func Nanoid(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwsyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nid, err := nanoid.CustomASCII(str, len)
	if err != nil {
		logx.Errorf("nanoid错误,%s", err.Error())
	}
	return nid()
}

// nanoid 默认21位
func NanoidDef() string {
	nid, err := nanoid.Standard(21)
	if err != nil {
		logx.Errorf("nanoid错误 %s", err.Error())
	}
	return nid()
}
