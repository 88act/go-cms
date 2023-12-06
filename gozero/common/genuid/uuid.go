package genuid

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sony/sonyflake"
	"github.com/speps/go-hashids/v2"
	"github.com/zeromicro/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenFlakeId() int64 {

	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}
	return int64(id)
}

// 22位 nanoid
func GenNanoId() string {
	uuid, err := gonanoid.New(22)
	if err == nil {
		return uuid
	} else {
		logx.Errorf("GenNanoId 错误 err=%s", err.Error())
		return ""
	}
}

// hashid
func HashIds(id int) string {
	hd := hashids.NewData()
	hd.Salt = "jiquxr.com"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	if res, err := h.Encode([]int{id}); err == nil {
		return res
	} else {
		logx.Errorf("HashIds 错误 err=%s", err.Error())
		return ""
	}
}

func HashIdsDe(idStr string) int {
	hd := hashids.NewData()
	hd.Salt = "jiquxr.com"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	if res, err := h.DecodeWithError(idStr); err == nil {
		return res[0]
	} else {
		logx.Errorf("HashIdsDe 错误 err=%s", err.Error())
		return 0
	}
}

func HashIds64(id int64) string {
	hd := hashids.NewData()
	hd.Salt = "jiquxr.com"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	if res, err := h.EncodeInt64([]int64{id}); err == nil {
		return res
	} else {
		logx.Errorf("GenHashId64 错误 err=%s", err.Error())
		return ""
	}
}

func HashIds64De(idStr string) int64 {
	hd := hashids.NewData()
	hd.Salt = "jiquxr.com"
	hd.MinLength = 10
	h, _ := hashids.NewWithData(hd)
	if res, err := h.DecodeInt64WithError(idStr); err == nil {
		return res[0]
	} else {
		logx.Errorf("HashIds64De 错误 err=%s", err.Error())
		return 0
	}
}
