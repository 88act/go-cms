package genuid

import (
	"fmt"
	"testing"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/speps/go-hashids/v2"
)

func TestGenSn(t *testing.T) {
	fmt.Println("11111111111")
	sn := GenSn(SN_PREFIX_ACT)
	fmt.Println(sn)
}

func TestNanoId(t *testing.T) {

	id, err := gonanoid.New()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated  1 id: %s\n", id)
	// Custom length
	id, err = gonanoid.New(5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated 2 id: %s\n", id)

	// Custom alphabet
	id, err = gonanoid.Generate("abcdefg1234", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated 3 id: %s\n", id)

	// Custom non ascii alphabet
	id, err = gonanoid.Generate("こちんにабдежиклмнは你好喂שלום😯😪🥱😌😛äöüß", 10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Generated 4 id: %s\n", id)

	fmt.Printf("Generated 5 id: %s\n", gonanoid.Must())
	fmt.Printf("Generated 6 id: %s\n", gonanoid.MustGenerate("🚀💩🦄🤖", 4))
}

func TestHashIds(t *testing.T) {
	hd := hashids.NewData()
	hd.Salt = "gocms"
	hd.MinLength = 2
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{90022212, 1223213})
	fmt.Println(e)
	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}

// Generate ID

// id, err := gonanoid.New()
// Generate ID with a custom alphabet and length

// id, err := gonanoid.Generate("abcde", 54)

// hd := hashids.NewData()
// 	hd.Salt = "this is my salt"
// 	hd.MinLength = 30
// 	h, _ := hashids.NewWithData(hd)
// 	e, _ := h.Encode([]int{45, 434, 1313, 99})
// 	fmt.Println(e)
// 	d, _ := h.DecodeWithError(e)
// 	fmt.Println(d)
