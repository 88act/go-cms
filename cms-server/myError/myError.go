package myError

import (
	"fmt"
	"time"
)

type ErrorType int

//在一个 const 声明块中，iota 的初始值为0，每声明一个变量，自增1。以上的代码可以简化成
const (
	//成功=0
	ErrOK ErrorType = iota
	//失败=1
	ErrFail

	//DB更新失败
	ErrDbUpdate
	//DB新增失败
	ErrDbInsert
	//DB删除失败
	ErrDbDel

	//File删除失败
	ErrFileDel
	//File写入失败
	ErrFileWrite
	//File读取失败
	ErrFileRead

	//启动收集器错误
	ErrRunCollectErr
	//停止收集器错误
	ErrStopCollectErr
)

/**
*  自定义异常结构体
 */
type MyError struct {
	Type      ErrorType //异常操作
	Msg       string    //异常信息
	ErrorTime string    //异常时间
}

func (myError *MyError) Error() string {
	errInfo := fmt.Sprintf(" msg:%s,time:%s\n", myError.Msg, myError.ErrorTime)
	return errInfo
}

// func (myError *MyError) New(errorType ErrorType, msg string) error {
// 	myErr := new(MyError)
// 	myErr.Type = errorType
// 	myErr.Msg = msg
// 	now := time.Now()
// 	myErr.ErrorTime = now.Format("2006-01-02 15:04:05")
// 	return myErr
// }

func New(errorType ErrorType, msg string) MyError {
	now := time.Now()
	myErr := MyError{errorType, msg, now.Format("2006-01-02 15:04:05")}
	return myErr
}
