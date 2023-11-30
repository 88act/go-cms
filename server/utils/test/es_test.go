package utils

import (
	"context"
	"fmt"
	"testing"

	"github.com/olivere/elastic/v7"
)

var ctx = context.Background()
var esUrl string = "http://127.0.0.1:9200"

func Test_es(t *testing.T) {

	//连接客户端
	client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	//Client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200/"))
	fmt.Println(client, err)
	name := "people2"
	client.CreateIndex(name).Do(context.Background())
}

//插入数据
func Test_es2(t *testing.T) {
	Client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	fmt.Println("===========1===========")
	fmt.Println(Client, err)
	fmt.Println("===========2===========")
	name := "people2"
	data := `{
	"name": "wali23332",
	"country": "Chian243432",
	"age": 30,
	"date": "1987-03-07"
	}`
	_, err = Client.Index().Index(name).Type("man1").Id("4").BodyJson(data).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
	}

}

// 查找数据： //通过id查找
func Test_es3(t *testing.T) {
	Client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	fmt.Println(Client, err)
	name := "people2"
	get, err := Client.Get().Index(name).Type("man1").Id("1").Do(context.Background())
	fmt.Println(get, err)
}

//修改
func Test_es4(t *testing.T) {
	Client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	res, err := Client.Update().
		Index("people2").
		Type("employee").
		Id("1").
		Doc(map[string]interface{}{"age": 88}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)

}

//删除数据
func Test_es5(t *testing.T) {
	// Client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	// //使用结构体
	// e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	// //创建
	// put1, err := Client.Index().
	// 	Index("megacorp").
	// 	Type("employee").
	// 	Id("1").
	// 	BodyJson(e1).
	// 	Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// //删除
	// get, err := Client.Get().Index("megacorp").Type("employee").Id("1").Do(context.Background())
	// fmt.Println(get, err)
}
