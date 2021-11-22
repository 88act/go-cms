package utils

import (
	"fmt"
	"testing"
)

func Test_IsEmpty(t *testing.T) {
	var i int
	var i2 int
	i2 = 1
	var s string = "xiaochuan"
	var s2 string = ""
	i3 := 0
	var i4 *int
	fmt.Println(IsEmpty(i4))
	i4 = &i2
	fmt.Println(IsEmpty(i4))
	fmt.Println(IsEmpty(i3))
	fmt.Println(IsEmpty(i))
	fmt.Println(IsEmpty(i2))
	fmt.Println(IsEmpty(s))
	fmt.Println(IsEmpty(s2))

	var arr [2]string
	fmt.Println(IsEmpty(arr))

}
