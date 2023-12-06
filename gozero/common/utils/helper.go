package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 微秒 0.001.001
func MicrosecondsStr(latency time.Duration) int64 {
	return latency.Microseconds()
}

// struct转map
func Stru2map(content interface{}) map[string]interface{} {
	var name map[string]interface{}
	if marshalContent, err := json.Marshal(content); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&name); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range name {
				fmt.Println(k)
				fmt.Println(v)
				delete(name, k)
				k2 := CamelCaseToUdnderscore(k)
				name[k2] = v
			}
		}
	}
	fmt.Println("name-------")
	fmt.Println(name)
	return name
}

// 三目运算函数
func Ternary(a bool, b, c interface{}) interface{} {
	if a {
		return b
	}
	return c
}

// func If(cond bool, a, b interface{}) {
//     if cond {
//         return a
//     }

//     return b
// }

// age := 20
// val := If(age > 18, "成年人", "未成年人").(string)

// struct 转 切片
func Strct2Slice(f interface{}, sheetFieldsJson []string) []interface{} {
	v := reflect.ValueOf(f)
	ss := []interface{}{} // make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		ss = append(ss, v.Field(i))
		//ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}

// 转为int指针
func StringPtr(s string) *string {
	return &s
}

// 转为string指针
func IntPtr(s int) *int {
	return &s
}

// 判断是否为空
// 0， nil ，"", 空数组 = true
func IsEmpty(params interface{}) bool {

	//初始化变量
	var (
		flag          bool = true
		default_value reflect.Value
	)

	r := reflect.ValueOf(params)
	//获取对应类型默认值
	default_value = reflect.Zero(r.Type())
	//由于params 接口类型 所以default_value也要获取对应接口类型的值 如果获取不为接口类型 一直为返回false
	if !reflect.DeepEqual(r.Interface(), default_value.Interface()) {
		flag = false
	}
	return flag
}

// 空字符串 返回true 否则返回 false
func IsEmptyStr(params string) bool {
	str := strings.TrimSpace(params)
	if str == "" {
		return true
	} else {
		return false
	}
}

// func IsEmptyObj(params *interface{}) bool {
// 	if params == nil {
// 		return true
// 	}
// 	return IsEmpty(*params)
// }

// StrToTime 字符串转time
func Str2Time(str string) *time.Time {
	timeFormat := "2006-01-02 15:04:05"
	// 时区
	Loc, _ := time.LoadLocation("Asia/Shanghai")
	t, e := time.ParseInLocation(timeFormat, str, Loc)
	if e != nil {
		fmt.Println("时间转换错误: str =" + str + "," + e.Error())
		//return nil
	}
	return &t
}

// numToTime 时间戳转time ,注意是否纳秒
func Int2Time(ts int64) *time.Time {
	//timeFormat := "2006-01-02 15:04:05"
	// 时区
	//Loc, _ := time.LoadLocation("Asia/Shanghai")
	//时间戳转化为日期
	t := time.Unix(ts, 0)
	return &t
}

// StrToInt string 转int
func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

// StrToInt string 转int64
func StrToInt64(str string) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// StrToUInt string 转uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

// 阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Ip2long(ipstr string) (ip uint32) {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})`
	reg, err := regexp.Compile(r)
	if err != nil {
		return
	}
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)

	return
}
func Long2ip(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
func GetIp() string {
	ipArr := [15][2]int{
		{607649792, 608174079},
		{975044608, 977272831},
		{999751680, 999784447},
		{1019346944, 1019478015},
		{1038614528, 1039007743},
		{1783627776, 1784676351},
		{1947009024, 1947074559},
		{1987051520, 1988034559},
		{2035023872, 2035154943},
		{2078801920, 2079064063},
		{-1950089216, -1948778497},
		{-1425539072, -1425014785},
		{-1236271104, -1235419137},
		{-770113536, -768606209},
		{-569376768, -564133889},
	}
	randKey := rand.Intn(14)
	ip := Long2ip(uint32(rand.Intn(ipArr[randKey][1]-ipArr[randKey][0]) + ipArr[randKey][0]))
	return ip
}

func UrlQueryStrToMap(urlstr string) (map[string]interface{}, error) {
	formData := make(map[string]interface{})
	if len(urlstr) < 5 {
		return formData, errors.New("url有误")
	}
	u, err := url.Parse(urlstr)
	if err != nil {
		return formData, err
	}

	// 组装map
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return formData, err
	}

	if len(m) > 0 {
		for k, v := range m {
			if len(v) > 0 {
				formData[k] = v[0]
			} else {
				formData[k] = ""
			}
		}
	}

	return formData, nil
}

// GetType : 获取数据类型字符串 (string, int, float64, []int, []string, map[string]int ...)
// GetType : (能不用则不用,由于涉及到使用reflect包,性能堪忧)
func GetType(params interface{}) string {
	//数据初始化
	v := reflect.ValueOf(params)
	//获取传递参数类型
	vT := v.Type()

	//类型名称对比
	return vT.String()
}

// InArray :给定元素值 是否在 指定的数组中
func InArray(needle interface{}, hystack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hystack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hystack.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hystack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
		//panic("needle only support string,int,int64 type")
	}
	//for _, item := range hystack {
	//	if GetType(needle) == GetType(item) {
	//		if needle == item {
	//			return true
	//		}
	//	}
	//}

	return false
}

// 获取 20160102150405  格式
func TimeToStr(t *time.Time) string {
	year := t.Format("2006")
	month := t.Format("01")
	day := t.Format("02")
	hour := t.Format("15")
	//min := t.Format("04")
	// second := time.Now().Format("05")
	str := year + month + day + hour
	return str
}

// date := time.Now().Format("2006-01-02 15:04:05")
// year:=time.Now().Year()
// month := time.Now().Format("01")
// day:=time.Now().Day()
// //或者

// year := time.Now().Format("2006")
// month := time.Now().Format("01")
// day := time.Now().Format("02")
// hour := time.Now().Format("15")
// min := time.Now().Format("04")
// second := time.Now().Format("05")

// var timeTemplates = []string{
// 	"2006-01-02 15:04:05", //常规类型
// 	"2006/01/02 15:04:05",
// 	"2006-01-02",
// 	"2006/01/02",
// 	"15:04:05",
// }

// /* 时间格式字符串转换 */
// func Str2Time(tm string) *time.Time {
// 	for i := range timeTemplates {
// 		t, err := time.ParseInLocation(timeTemplates[i], tm, time.Local)
// 		if nil == err && !t.IsZero() {
// 			return &t
// 		}
// 	}
// 	return &time.Time{}
// }

// func test() {
// 	var tms = []string{
// 		"2021-03-04 08:15:11",
// 		"2021/03/05 08:15:11",
// 		"2021-03-04",
// 		"2021/03/05",
// 		"08:15:11",
// 	}

// 	for _, v := range tms {
// 		fmt.Println("OrgTimeStr: ", v, "; Convert Result: ", Str2Time(v))
// 	}
// 	/* 输出结果:
// 	OrgTimeStr:  2021-03-04 08:15:11 ; Convert Result:  2021-03-04 08:15:11 +0800 CST
// 	OrgTimeStr:  2021/03/05 08:15:11 ; Convert Result:  2021-03-05 08:15:11 +0800 CST
// 	OrgTimeStr:  2021-03-04 ; Convert Result:  2021-03-04 00:00:00 +0800 CST
// 	OrgTimeStr:  2021/03/05 ; Convert Result:  2021-03-05 00:00:00 +0800 CST
// 	OrgTimeStr:  08:15:11 ; Convert Result:  0000-01-01 08:15:11 +0800 CST
// 	*/
// }

// 获取请求中的IP
func RemoteIp(req *http.Request) string {
	var remoteAddr string
	// X-Real-Ip
	remoteAddr = req.Header.Get("X-Real-Ip")
	if remoteAddr != "" {
		return remoteAddr
	} else {
		remoteAddr = "127.0.0.1"
	}
	// RemoteAddr
	remoteAddr = req.RemoteAddr
	if remoteAddr != "" {
		return remoteAddr
	}
	// ipv4
	remoteAddr = req.Header.Get("ipv4")
	if remoteAddr != "" {
		return remoteAddr
	}
	//
	remoteAddr = req.Header.Get("XForwardedFor")
	if remoteAddr != "" {
		return remoteAddr
	}
	// X-Forwarded-For
	remoteAddr = req.Header.Get("X-Forwarded-For")
	if remoteAddr != "" {
		return remoteAddr
	}

	return remoteAddr
}
