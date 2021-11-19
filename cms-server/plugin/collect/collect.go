package plugin

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/88act/go-cms/server/global"
	"github.com/88act/go-cms/server/model/business"
	"github.com/88act/go-cms/server/service"
	"github.com/88act/go-cms/server/utils"
	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/util/gconv"
)

var once sync.Once = sync.Once{}
var instanse *CollectManager

type News2 struct {
	Title   string
	Media   string
	Url     string
	PubTime string
	Content string
}

/**
* 数据采集管理器，单例
 */
type CollectManager struct {
	List []business.ColCollect
}

/**
*获取单例
 */
func GetCollectManager() *CollectManager {
	once.Do(func() {
		instanse = new(CollectManager)
		instanse.List = []business.ColCollect{}
		//instanse.Name = "这是一个单例"
	})
	return instanse
}

func (m *CollectManager) Start(collect business.ColCollect, opt int) (err error) {
	collect_old := business.ColCollect{}
	var idx int = -1
	for i, v := range m.List {
		if v.ID == collect.ID {
			collect_old = v
			idx = i
			break
		}
	}
	fmt.Println("idx === ", idx)
	if idx >= 0 {
		fmt.Println("*collect_old.StatusRun === ", *collect_old.StatusRun)
	}
	if opt == 1 {
		//启动
		if idx >= 0 && 1 == *collect_old.StatusRun {
			fmt.Println(collect.Name + "已经在运行了")
			return nil
		} else {
			//加入队列

			var statusRun int = 1
			collect.StatusRun = &statusRun
			m.List = append(m.List, collect)
			fmt.Println(collect.Name + "加入队列")
			fmt.Println("List 长度=====")
			fmt.Println(len(m.List))
			return nil
		}
	} else if opt == 0 {
		//停止
		if idx >= 0 {
			fmt.Println(collect.Name + " 已经在运行了 删除队列1个 ")
			//删除队列1个
			m.List = append(m.List[:idx], m.List[idx+1:]...)
			fmt.Println("List 长度=====")
			fmt.Println(len(m.List))
			return nil
		} else {
			fmt.Println(collect.Name + "原来没启动。。。")
			return nil
		}

	}
	return nil
	// //
	// url := collect.Url
	// resp, err := http.Get(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()
	// if resp.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	// }

	// html, err := goquery.NewDocumentFromReader(resp.Body)
	// //global.LOG.Debug(html.Text())
	// var newsList []string
	// newsList = getUrlList(html, newsList)
	// //fmt.Println(newsList)
	// var wg sync.WaitGroup
	// for i := 0; i < len(newsList); i++ {
	// 	wg.Add(1)
	// 	go getDetail(newsList[i], &wg, collect)
	// }
	// wg.Wait()
	// //开始下一页
	// global.LOG.Debug("完成 ---Wait--1-")
	// return nil
}

func (m *CollectManager) start_do(collect business.ColCollect, opt int) (err error) {

	//
	m.List = append(m.List, collect)
	//
	url := collect.Url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	//global.LOG.Debug(html.Text())
	var newsList []string
	newsList = getUrlList(html, newsList)
	//fmt.Println(newsList)
	var wg sync.WaitGroup
	for i := 0; i < len(newsList); i++ {
		wg.Add(1)
		go getDetail(newsList[i], &wg, collect)
	}
	wg.Wait()
	//开始下一页
	global.LOG.Debug("完成 ---Wait--1-")
	return nil
}

func (m *CollectManager) Stop(collect business.ColCollect) (err error) {

	//
	url := collect.Url
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	//global.LOG.Debug(html.Text())
	var newsList []string
	newsList = getUrlList(html, newsList)
	//fmt.Println(newsList)
	var wg sync.WaitGroup
	for i := 0; i < len(newsList); i++ {
		wg.Add(1)
		go getDetail(newsList[i], &wg, collect)
	}
	wg.Wait()
	//开始下一页
	global.LOG.Debug("完成 ---Wait--1-")
	return nil
}

func getUrlList(html *goquery.Document, newsList []string) []string {
	//global.LOG.Debug("getNewsList2==========")
	// '//a[@class="tt"]/@href'
	//html.Find("a[class=tt]").Each(func(i int, selection *goquery.Selection) {
	html.Find("li>h3>a").Each(func(i int, selection *goquery.Selection) {
		url, _ := selection.Attr("href")
		url = "http://www.cjhdj.com.cn/hdfw/hdtg/" + strings.Replace(url, "./", "", 1)
		//	global.LOG.Debug("html.Find------")
		global.LOG.Debug(url)
		newsList = append(newsList, url)
	})
	return newsList
}

func getDetail(url string, wg *sync.WaitGroup, collect business.ColCollect) {
	resp, err := http.Get(url)
	if err != nil {
		global.LOG.Error(err.Error())
		wg.Done()
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		global.LOG.Error("Error: status code =" + strconv.Itoa(resp.StatusCode))
		wg.Done()
		return
	}

	html, err := goquery.NewDocumentFromReader(resp.Body)
	//fmt.Println(html.Text())
	news := business.ColHsj{}

	news.Url = url
	news.PubUnit = collect.Name
	html.Find("p[class=biaoti]").Each(func(i int, selection *goquery.Selection) {
		news.Title = selection.Text()
	})

	//global.LOG.Debug("news.Title = " + news.Title)
	if news.Title == "" {
		wg.Done()
		return
	}

	html.Find("p[style=\"white-space: pre-wrap;\"]").Each(func(i int, selection *goquery.Selection) {
		news.Content = news.Content + selection.Text()
	})

	var tmpTime string
	html.Find("div[class=xl_tit2]>div").Each(func(i int, selection *goquery.Selection) {
		tmpTime = selection.Text()
	})
	reg := regexp.MustCompile(`\d+`)
	timeString := reg.FindAllString(tmpTime, -1)
	strTime := fmt.Sprintf("%s-%s-%s %s:%s:00", timeString[0], timeString[1], timeString[2], timeString[3], timeString[4])
	colId := gconv.Int(collect.ID)
	news.ColId = &colId
	news.PubTime = utils.Str2Time(strTime)
	status := 1
	news.Status = &status
	err = service.ServiceGroupApp.BusinessServiceGroup.ColHsjService.CreateColHsj(news)
	global.LOG.Debug("完成 ---")

	// fmt.Println("news.Title=", news.Title)
	// fmt.Println("news.Url=", news.Url)
	// fmt.Println("news.PubUnit=", news.PubUnit)
	// fmt.Println("news.content=", news.Content)
	// fmt.Println("news.PubTime=", news.PubTime)

	wg.Done()
}
