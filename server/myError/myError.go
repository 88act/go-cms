package myError

import (
	"fmt"
)

//自定义异常结构体
type MyError struct {
	Type ErrorType //异常操作
	Msg  string    //异常信息
	//ErrorTime string    //异常时间
}

func (e MyError) Error() string {
	errInfo := fmt.Sprintf(" msg:%s ", e.Msg)
	return errInfo
}

func New(errorType ErrorType, msg string) MyError {
	//now := time.Now()
	//myErr := MyError{errorType, msg, now.Format("2006-01-02 15:04:05")}
	myErr := MyError{errorType, msg}
	return myErr
}
func NewP(errorType ErrorType, msg string) *MyError {
	//now := time.Now()
	myErr := MyError{errorType, msg}
	return &myErr
}

// func (myError *MyError) New(errorType ErrorType, msg string) error {
// 	myErr := new(MyError)
// 	myErr.Type = errorType
// 	myErr.Msg = msg
// 	now := time.Now()
// 	myErr.ErrorTime = now.Format("2006-01-02 15:04:05")
// 	return myErr
// }
