package mycache

import (
	"fmt"
	"testing"
)

func Test2(t *testing.T) {
	fmt.Println("Test2")
	//byteData := []byte("ssss112112qweqweqwe2")
	byteData := "ssss112112qweqweqwe2"
	GetCache().Set("ssss", byteData, 0)
	entry, err := GetCache().Get("ssss")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print(entry.(string))
	}
}
