package common

import (
	"context"
	"fmt"
	"sync"

	"go-cms/global"

	"github.com/olivere/elastic/v7"
)

var obj *EsHelper
var once sync.Once

//EsHelper es的连接
type EsHelper struct {
	Client *elastic.Client
	host   string
}

var esUrl string = "http://127.0.0.1:9200"

func GetEsHelper() *EsHelper {
	once.Do(func() {
		obj = &EsHelper{}
		err := obj.connect(esUrl)
		if err != nil {
			fmt.Println("------EsHelper Connect err =  ", err.Error())
		}
	})
	return obj
}

//Connect 连接
func (m *EsHelper) connect(ip string) error {
	//client, err := elastic.NewClient(elastic.SetURL(esUrl), elastic.SetSniff(false))
	//连接客户端
	client, err := elastic.NewClient(elastic.SetURL(ip), elastic.SetSniff(false))
	if err != nil {
		return err
	}
	_, _, err = client.Ping(ip).Do(context.Background())
	if err != nil {
		return err
	}
	//fmt.Printf("EsHelpersearch returned with code %d and version %s\n", code, info.Version.Number)
	ve, err := client.ElasticsearchVersion(ip)
	if err != nil {
		return err
	}
	fmt.Printf("EsHelpersearch  Version " + ve)
	m.Client = client
	m.host = ip
	return nil
}

//CreateIndex 创建一个index, mapping这个你可以不需要
func (Es *EsHelper) CreateIndex(index, mapping string) bool {
	fmt.Printf("a---------1------- ")
	// 判断索引是否存在
	exists, err := Es.Client.IndexExists(index).Do(context.Background())
	if err != nil {
		fmt.Sprintf("<CreateIndex> some error occurred when check exists, index: %s, err:%s", index, err.Error())
		fmt.Printf("a---------1---1---- ")
		fmt.Printf(err.Error())
		return false
	}

	if exists {
		fmt.Printf("a---------1---2---- ")
		fmt.Sprintf("<CreateIndex> index:{%s} is already exists", index)
		fmt.Printf(err.Error())
		return true
	}
	fmt.Printf("a----------2------ ")
	createIndex, err := Es.Client.CreateIndex(index).Body(mapping).Do(context.Background())
	if err != nil {
		fmt.Sprintf("<CreateIndex> some error occurred when create. index: %s, err:%s", index, err.Error())
		return false
	}
	if !createIndex.Acknowledged {
		// Not acknowledged
		fmt.Sprintf("<CreateIndex> Not acknowledged, index: %s", index)
		return false
	}
	fmt.Printf("b---------3------- ")
	return true
}

//DelIndex 删除Index
func (Es *EsHelper) DelIndex(index string) bool {
	// Delete an index.
	deleteIndex, err := Es.Client.DeleteIndex(index).Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Sprintf("<DelIndex> some error occurred when delete. index: %s, err:%s", index, err.Error())
		return false
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
		fmt.Sprintf("<DelIndex> acknowledged. index: %s", index)
		return false
	}
	return true
}

//Put 上传数据
func (Es *EsHelper) Put(index string, typ string, id string, bodyJSON interface{}) bool {
	_, err := Es.Client.Index().
		Index(index).
		Type(typ).
		Id(id).
		BodyJson(bodyJSON).
		Do(context.Background())
	if err != nil {
		// Handle error
		msg := fmt.Sprintf("put 发生错误.  err:%s", err.Error())
		global.LOG.Error(msg)
		return false
	}
	return true
}

//////Get List
func (Es *EsHelper) GetList(index, typ, starttime, endtime, keyword string, size, page int) (*elastic.SearchResult, error) {

	boolQ := elastic.NewBoolQuery()
	keys := fmt.Sprintf("name:*%s*", keyword)
	//boolQ.Filter(elastic.NewRangeQuery("CreateAt").Gte(starttime), elastic.NewRangeQuery("CreateAt").Lte(endtime), elastic.NewQueryStringQuery(keys))
	boolQ.Filter(elastic.NewQueryStringQuery(keys))
	res, err := Es.Client.Search(index).Type(typ).Query(boolQ).Size(size).From((page - 1) * size).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil

	// boolQ := elastic.NewBoolQuery()
	// if keyword == "" {
	// 	boolQ.Filter(elastic.NewRangeQuery("starttime").Gte(starttime), elastic.NewRangeQuery("starttime").Lte(endtime))
	// 	res, err := Es.Client.Search(index).Type(typ).Query(boolQ).Size(size).From((page - 1) * size).Do(context.Background())
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return res, nil
	// }
	// keys := fmt.Sprintf("name:*%s*", keyword)
	// boolQ.Filter(elastic.NewRangeQuery("CreateAt").Gte(starttime), elastic.NewRangeQuery("CreateAt").Lte(endtime), elastic.NewQueryStringQuery(keys))
	// res, err := Es.Client.Search(index).Type(typ).Query(boolQ).Size(size).From((page - 1) * size).Do(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// return res, nil

}

//Del 删除指定id数据
func (Es *EsHelper) Del(index, typ, id string) bool {
	del, err := Es.Client.Delete().
		Index(index).
		Type(typ).
		Id(id).
		Do(context.Background())
	if err != nil {
		// Handle error
		fmt.Sprintf("<Del> some error occurred when put.  err:%s", err.Error())
		return false
	}
	fmt.Sprintf("<Del> success, id: %s to index: %s, type %s\n", del.Id, del.Index, del.Type)
	return true
}

//Update 更新数据
func (Es *EsHelper) Update(index, typ, id string, updateMap map[string]interface{}) bool {
	res, err := Es.Client.Update().
		Index(index).Type(typ).Id(id).
		Doc(updateMap).
		FetchSource(true).
		Do(context.Background())
	if err != nil {
		_ = fmt.Sprintf("<Update> some error occurred when update. index:%s, typ:%s, id:%s err:%s", index, typ, id, err.Error())
		return false
	}
	if res == nil {
		fmt.Sprintf("<Update> expected response != nil. index:%s, typ:%s, id:%s", index, typ, id)
		return false
	}
	if res.GetResult == nil {
		fmt.Sprintf("<Update> expected GetResult != nil. index:%s, typ:%s, id:%s", index, typ, id)
		return false
	}
	//data, _ := json.Marshal(res.GetResult.Source)
	//fmt.Println("<Update> update success. data:%s", data)
	return true
}

//GetTaskLogCount
func (Es *EsHelper) GetTaskLogCount(index, starttime, endtime string) (int, error) {
	boolQ := elastic.NewBoolQuery()
	boolQ.Filter(elastic.NewRangeQuery("time").Gte(starttime), elastic.NewRangeQuery("time").Lte(endtime))
	//统计count
	count, err := Es.Client.Count(index).Type("doc").Query(boolQ).Do(context.Background())
	if err != nil {
		return 0, nil
	}
	return int(count), nil
}

//GetSourceByID
func (Es *EsHelper) GetSourceByID(index, typ, esid string) (*elastic.GetResult, error) {
	source, err := Es.Client.Get().Index(index).Type(typ).Id(esid).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return source, nil
}

//////GetaskMsg
func (Es *EsHelper) GetaskMsg(index, typ, starttime, endtime, keyword string, size, page int) (*elastic.SearchResult, error) {
	boolQ := elastic.NewBoolQuery()
	if keyword == "" {
		boolQ.Filter(elastic.NewRangeQuery("starttime").Gte(starttime), elastic.NewRangeQuery("starttime").Lte(endtime))
		res, err := Es.Client.Search(index).Type(typ).Query(boolQ).Size(size).From((page - 1) * size).Do(context.Background())
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	keys := fmt.Sprintf("name:*%s*", keyword)
	boolQ.Filter(elastic.NewRangeQuery("starttime").Gte(starttime), elastic.NewRangeQuery("starttime").Lte(endtime), elastic.NewQueryStringQuery(keys))
	res, err := Es.Client.Search(index).Type(typ).Query(boolQ).Size(size).From((page - 1) * size).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}
