package utils

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	// fmt.Println("My name is Leaf-------")
	// d := timer.NewDispatcher(10)

	// // timer 1
	// d.AfterFunc(1, func() {
	// 	fmt.Println("My name is Leaf")
	// })

	// timer 2
	// mytimer := d.AfterFunc(1, func() {
	// 	fmt.Println("will not print")
	// })
	// mytimer.Stop()
	// // dispatch
	// (<-d.ChanTimer).Cb()

	// Output:
	// My name is Leaf
}

type Gender int8

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

func Test1(t *testing.T) {
	gender := FEMALE

	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}

func Test2(t *testing.T) {
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}
	// 0 10
	// 1 20
	// 2 30
	// 3 40
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	for key, value := range m2 {
		fmt.Println(key, " = ", value)
	}
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func Test3(t *testing.T) {
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)     // 5 15
	fmt.Println(add(100, 17)) // 117
}

//  try...catch 机制 ==== Go 语言也提供了类似的机制 defer 和 recover
// get 函数中，使用 defer 定义了异常处理的函数，在协程退出前，会执行完 defer 挂载的任务。因此如果触发了 panic，控制权就交给了 defer。
// 在 defer 的处理逻辑中，使用 recover，使程序恢复正常，并且将返回值设置为 -1，在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0。
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("异常处理 Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func Test4(t *testing.T) {
	fmt.Println(get(5))
	fmt.Println("finished")
}

// 结构体(struct) 和方法(methods)
// 结构体类似于其他语言中的 class，可以在结构体中定义多个字段，为结构体实现方法，实例化等。
// 接下来我们定义一个结构体 Student，并为 Student 添加 name，age 字段，并实现 hello() 方法
//实现方法与实现函数的区别在于，func 和函数名hello 之间，加上该方法对应的实例名 stu 及其类型 *Student，
// 可以通过实例名访问该实例的字段name和其他方法了。 调用方法通过 实例名.方法名(参数) 的方式。
type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("Student的hello方法 :  %s, I am %s", person, stu.name)
}

func Test5(t *testing.T) {
	stu := &Student{name: "Tom"}
	msg := stu.hello("学生1jack")
	fmt.Println(msg) // hello Jack, I am Tom
	stu2 := new(Student)
	stu2.name = "Alice"
	fmt.Println(stu2.hello("学生2 Alice")) // hello Alice, I am  , name 被赋予默认值""
}

//接口(interfaces)
//  一般而言，接口定义了一组方法的集合，接口不能被实例化，一个类型可以实现多个接口。
//  举一个简单的例子，定义一个接口 Person和对应的方法 getName() 和 getAge()：
//Go 语言中，并不需要显式地声明实现了哪一个接口，只需要直接实现该接口对应的方法即可
//实例化 Student2后，强制类型转换为接口类型 Person。

type Person interface {
	getName() string
	getAge() int
}

type Student2 struct {
	name string
	age  int
}

func (stu *Student2) getName() string {
	return "学生2返回 = " + stu.name
}
func (stu *Student2) getAge() int {
	return stu.age
}

type Worker2 struct {
	name   string
	age    int
	gender string
}

func (w *Worker2) getName() string {
	return "工人2返回 = " + w.name + w.gender
}
func (w *Worker2) getAge() int {
	return w.age
}

func Test6(t *testing.T) {
	//如何确保某个类型实现了某个接口的所有方法呢？
	//一般可以使用下面的方法进行检测，如果实现不完整，编译期将会报错
	var _ Person = (*Student2)(nil)
	var _ Person = (*Worker2)(nil)

	var p Person = &Student2{
		name: "Tom",
		age:  18,
	}
	var w Person = &Worker2{
		name:   "Tom",
		gender: "男",
	}

	fmt.Println(p.getName()) //
	fmt.Println(w.getName()) //
}

//实例可以强制类型转换为接口，接口也可以强制类型转换为实例
func Test7(t *testing.T) {
	var p Person = &Student2{
		name: "Tom",
		age:  18,
	}
	stu := p.(*Student2) // 接口转为实例
	fmt.Println(stu.getAge())
}

//如果定义了一个没有任何方法的空接口，那么这个接口可以表示任意类型
func Test8(t *testing.T) {
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m) // map[age:18 name:Tom scores:[98 99 85]]
}

//并发编程(goroutine)
// 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发
//希望并发下载 N 个资源，多个并发协程之间不需要通信，那么就可以使用 sync.WaitGroup，等待所有并发协程执行结束
//  strconv.Itoa()：整型转字符串
// strconv.Atoi()：字符串转整型 返回2个值，第一个是值，第二个是错误，下面没有做处理
var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}
func Test9(t *testing.T) {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go download("a.com/" + strconv.Itoa(i))
	}
	wg.Wait()
	fmt.Println("Done!")
}

// channel 信道，可以在协程之间传递消息。阻塞等待并发协程返回消息
//对于一个关闭的channel，如果继续向channel发送数据，会引起panic
//channel关闭之后，仍然可以从channel中读取剩余的数据，直到数据全部读取完成
// value, ok := <- ch  ok是false，就表示已经关闭
//for value := range ch {
//   如果channel被关闭，会跳出循环
//}

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func download2(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

var maxChan int = 10
var nowChan int = 0

func Test10(t *testing.T) {
	for i := 0; i < maxChan; i++ {
		go download2("a.com/" + strconv.Itoa(i))
	}
	for value := range ch {
		nowChan++
		fmt.Println("信道 finish", value)
		fmt.Println("nowChan = ", nowChan)
		if nowChan >= maxChan {
			close(ch)
			fmt.Println("关闭　ｃｈａｎ = ")
			break
		}
	}

	// for i := 0; i < 3; i++ {
	// 	//msg := <-ch // 等待信道返回消息。
	// 	value, ok := <-ch

	// 	fmt.Println("信道 finish", value)
	// 	if !ok {
	// 		break
	// 	}
	// }
	fmt.Println("Done!")
}

// map
func Test11(t *testing.T) {
	//声明并初始化一个map，key是int64类型，value是string类型

	myMap := make(map[string]string)

	myMap["中国1"] = "value1"

	myMap["美国1"] = "value2"

	myMap["中国2"] = "value5"

	myMap["泰国2"] = "value6"

	//声明一个int64数组，然后遍历数组，num是数组中的元素，下划线_代表元素的下标位置

	for _, num := range []string{"中国1", "中国2", "中国3", "中国4", "中国5", "中国6"} {

		//不关心数组中的value，用下划线代替
		//一般都习惯用ok变量表示是否包含，也可以用别的变量名字

		if _, ok := myMap[num]; ok {

			fmt.Printf("myMap中包含key:%s \n", num)

		} else {

			fmt.Printf("myMap中不包含key:%s\n", num)

		}

	}

	fmt.Println("=================分割线=======================")

	for _, num := range []string{"中国1", "中国2", "中国3", "中国4", "中国5", "中国6"} {

		//如果包含key，想知道value，就把返回值赋给一个变量，这儿用变量v
		//这儿用变量s表示是否包含指定的key

		if v, s := myMap[num]; s {

			fmt.Printf("myMap中包含key:%s,value值为:%s\n", num, v)

		} else {

			fmt.Printf("myMap中不包含key:%s\n", num)

		}

	}
}

// 指针(pointer) 即某个值的地址，类型定义时使用符号*，对一个已经存在的变量，使用 & 获取该变量的地址
//指针通常在函数传递参数，或者给某个类型定义新的方法时使用。Go 语言中，参数是按值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，
//对参数的修改并不会影响到外部变量的值。如果参数使用指针，对参数的传递将会影响到外部变量
func Test12(t *testing.T) {

	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变

	num := 100
	addd(num)
	fmt.Println(num) // 100，num 没有变化

	realAddd(&num)
	fmt.Println(num) // 101，指针传递，num 被修改

}
func addd(num int) {
	num += 1
}

func realAddd(num *int) {
	*num += 1
}

// 数组(array)
// 与切片(slice)
func Test13(t *testing.T) {
	//var arr [5]int     // 一维
	//var arr2 [5][5]int // 二维
	//var arr3 = [5]int{1, 2, 3, 4, 5}
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100 //使用 [] 索引/修改数组
	}
	fmt.Println(arr) // [101 102 103 104 105]
	fmt.Println("切片 ----------- ")
	//切片包含三个组件：容量，长度和指向底层数组的指针,
	slice1 := make([]float32, 0)          // 长度为0的切片
	slice2 := make([]float32, 3, 5)       // [0 0 0] 长度为3容量为5的切片
	fmt.Println(len(slice1), cap(slice1)) // 3 5
	fmt.Println(len(slice2), cap(slice2)) // 3 5
	// 添加元素，切片容量可以根据需要自动扩展
	slice2 = append(slice2, 1, 2, 3, 4) // [0, 0, 0, 1, 2, 3, 4]
	fmt.Println(slice2)
	fmt.Println(len(slice2), cap(slice2)) // 7 12
	// 子切片 [start, end)
	sub1 := slice2[3:] // [1 2 3 4]
	sub2 := slice2[:3] // [0 0 0]
	//sub3 := slice2[1:4] // [0 0 1]
	fmt.Println("合并切片 ----------- ")
	// 合并切片
	combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]
	fmt.Println(combined)
	fmt.Println(len(combined), cap(combined)) // 7 12

}

func Test14(t *testing.T) {
	timer := time.NewTimer(1 * time.Second)
	for {
		select {
		case <-timer.C:
			println("1_Tokyo_NewYork_London")
			timer.Reset(5 * time.Second)
			//timer.Reset(2 * time.Second)
			// 若要停止定时器就使用如下代码，这样会释放资源
			//timer.Stop()
		}
		//break
	}

}
func Test15(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	num := 0
	var p *int = &num
	for {
		select {
		case <-ticker.C:

			// 若要停止定时器就使用如下代码，这样会释放资源
			*p++
			println("2_NewYork_London_Tokyo num =", num, *p, p, &num, &p)
			if *p > 5 {
				ticker.Stop()
			}
		}
		// break
	}
}

func Test16(t *testing.T) {
	f := func() { println("Tokyo_London_NewYork") }
	// AfterFunc的内部又启动了一个协程，且这个定时器到了定时时间后，只会执行一次！
	time.AfterFunc(4*time.Second, f)
	time.Sleep(5 * time.Second)

}

type Student4 struct {
	name string
	age  int
}

// 1、适用范围：make 只能创建类型(slice map channel)， new可以对所有类型进行内存分配
// 2、返回值： new 返回指针， make 返回引用
// new返回一个指向已清零内存的指针，而make返回一个复杂的结构。
// make返回复杂的结构为slice时:它是一个包含3个域的结构体：指向slice中第一个元素的指针，slice的长度，以及slice的容量
// 3、填充值： new 填充零值， make 填充非零值
// new(T)会为T类型的新项目，但new它并不初始化内存，只是将其置零
// make(T, args)返回一个初始化的(而不是置零)，类型为T的值（而不是*T）。之所以有所不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构
// 例如
// make([]int, 10, 100) 分配一个有100个int的数组，然后创建一个长度为10，容量为100的slice结构，该slice引用包含前10个元素的数组，对应的，new([]int)返回一个指向新分配的，被置零的slice结构体的指针，即指向值为nil的slice的指针
// 4、直接使用struct{} 来初始化strut时，返回的是一个struct类型的值，而不是指针
func Test17(t *testing.T) {
	var s1 *Student4 = new(Student4)
	s1.age = 22
	s1.name = "tcy"
	var s2 Student4 = Student4{}
	s2.age = 21
	s2.name = "tcy333"
	fmt.Println(s1, s2) //&{tcy 22} {tcy 21}
}

func foo(bar string) {
	fmt.Println("2222")
	fmt.Printf("in foo('%s')", bar)
}

func newFunc(bar string) func() {
	fmt.Printf("creating func with '%s' ", bar)
	fmt.Println("11111")
	return func() {
		foo(bar)
	}
}

func Test18(t *testing.T) {
	somebar := "Here we go!"
	f := newFunc(somebar)
	_ = time.AfterFunc(1*time.Second, f)
	time.Sleep(2 * time.Second)
}
