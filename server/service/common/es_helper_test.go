package common

import (
	"fmt"
	"testing"
)

func Test_es(t *testing.T) {
	es := GetEsHelper()
	fmt.Println("es.host = ", es.host)
	es.CreateIndex("test4", "")
}

func Test_es2(t *testing.T) {
	es := GetEsHelper()
	ok := es.DelIndex("test3")
	fmt.Println(ok)
}

// func Test_es3(t *testing.T) {
// 	es := GetEsHelper()
// 	cmsCat := business.CmsCat{}
// 	cmsCat.ID = 6
// 	cmsCat.Name = "Name6"
// 	cmsCat.CreatedAt = //gconv.Time("2021-12-6")
// 	//cmsCat.Desc = "Desc6"
// 	ok := es.Put("test", "CmsCat", "id6", cmsCat)
// 	fmt.Println(ok)
// }

func Test_es4(t *testing.T) {
	es := GetEsHelper()
	ok := es.Del("test", "CmsCat", "id1")
	fmt.Println(ok)
}

func Test_es5(t *testing.T) {
	es := GetEsHelper()
	myMap := make(map[string]interface{})
	myMap["name"] = "Name11111222"
	myMap["desc"] = "desc111112222"
	ok := es.Update("test", "CmsCat", "id1", myMap)
	fmt.Println(ok)
}

func Test_es6(t *testing.T) {
	es := GetEsHelper()
	list, err := es.GetList("test", "CmsCat", "2021-12-1", "2021-12-13", "Name2", 10, 1)
	if err != nil {
		fmt.Println("-----------err----------------")
		fmt.Println(err.Error())
	}
	fmt.Println("-----------list----------------")
	fmt.Println(list)
	//从搜索结果中取数据的方法
	// for _, item := range list.Each(reflect.TypeOf(&business.CmsCat{})) {
	// 	fmt.Println("搜索结果中取数据1111 ")
	// 	fmt.Println(item)
	// 	if t, ok := item.(business.CmsCat); ok {
	// 		fmt.Println("搜索结果中取数据2222 ")
	// 		fmt.Println(t)
	// 	}
	// }

}
