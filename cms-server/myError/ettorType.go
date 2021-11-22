package myError

type ErrorType int

//在一个 const 声明块中，iota 的初始值为0，每声明一个变量，自增1。以上的代码可以简化成
const (
	//成功=0
	ErrOK ErrorType = iota
	//失败=1
	ErrFail
	//DB没有数据
	ErrDbNoRecord
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
