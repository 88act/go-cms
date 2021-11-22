package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	 
	BaseRouter
	CasbinRouter
	DictionaryRouter
	DictionaryDetailRouter
	InitRouter
	JwtRouter
	MenuRouter
	OperationRecordRouter
	SysRouter
	UserRouter
}
