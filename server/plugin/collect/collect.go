package plugin

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"go-cms/global"
	"go-cms/model/business"
	"go-cms/myError"
	"go-cms/service"
	"go-cms/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/gogf/gf/util/gconv"
)

/**
* 数据采集管理器，单例
 */
type Collect struct {
	Data      business.ColCollect
	statusRun int
}

func (m *Collect) Stop() (err error) {
	m.statusRun = 0
	err = global.DB.Model(&m.Data).Update("status_run", 0).Error
	if err != nil {
		return myError.New(myError.ErrDbUpdate, "停止失败，更新数据库失败")
	} else {
		global.LOG.Debug("更新数据库 StatusRun 成功")
	}
	return myError.New(myError.ErrOK, "停止成功")
}

func (m *Collect) Start() (err error) {
	//更新数据库
	err = global.DB.Model(&m.Data).Update("status_run", 1).Error
	if err != nil {
		return myError.New(myError.ErrDbUpdate, "启动失败，更新数据库失败")
	} else {
		global.LOG.Debug("更新数据库 StatusRun成功")
	}
	global.LOG.Debug(m.Data.Name + " 启动，加入队列id=" + gconv.String(m.Data.ID))
	m.statusRun = 1
	go m.beginCollect()
	return myError.New(myError.ErrOK, "启动成功")
}
func (m *Collect) beginCollect() {
	global.LOG.Debug("beginCollect开始分页获取  PageStart =" + gconv.String(m.Data.PageStart) + "PageEnd=" + gconv.String(m.Data.PageEnd))
	var wg_page sync.WaitGroup
	for i := *m.Data.PageStart; i <= *m.Data.PageEnd; i++ {
		if m.statusRun == 1 {
			wg_page.Add(1)
			go m.getPage(i, &wg_page)
		} else {
			global.LOG.Debug("getPage 中途退出 id=" + gconv.String(m.Data.ID) + " , i=" + gconv.String(i))
			return
		}
	}
	wg_page.Wait()
	global.LOG.Debug("beginCollect完成全部 getPage-----")
	m.statusRun = 0
	m.Data.StatusRun = &m.statusRun
}

func (m *Collect) getPage(page int, wg_page *sync.WaitGroup) {
	url := m.Data.Url
	if page > 0 {
		url = fmt.Sprintf(m.Data.UrlPage, page)
	}
	global.LOG.Debug("page=" + gconv.String(page) + " ,page url = " + url)
	resp, err := http.Get(url)
	if err != nil {
		global.LOG.Error(err.Error())
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		//log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
		global.LOG.Error(fmt.Sprintf("status code error: %d %s", resp.StatusCode, resp.Status))
	}
	html, err := goquery.NewDocumentFromReader(resp.Body)
	//global.LOG.Debug(html.Text())
	var newsList []string
	newsList = getUrlList(html, newsList)
	//fmt.Println(newsList)
	var wg_detail sync.WaitGroup
	for i := 0; i < len(newsList); i++ {
		if m.statusRun == 1 {
			wg_detail.Add(1)
			go getDetail(newsList[i], &wg_detail, m.Data)
		} else {
			global.LOG.Debug("getDetail() 中途退出 id=" + gconv.String(m.Data.ID) + " , i=" + gconv.String(i))
			return
		}
	}
	wg_detail.Wait()
	global.LOG.Debug("getDetail()  正常完成 ")
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

func getDetail(url string, wg *sync.WaitGroup, data business.ColCollect) {
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
	news.PubUnit = data.Name
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
	colId := gconv.Int(data.ID)
	news.ColId = &colId
	news.PubTime = utils.Str2Time(strTime)
	status := 1
	news.Status = &status
	err = service.ServiceGroupApp.BusinessServiceGroup.ColHsjService.CreateColHsj(news)
	global.LOG.Debug("完成一个页面news ---")
	wg.Done()
}
