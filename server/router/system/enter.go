package system

type RouterGroup struct {
	ApiRouter
	AuthorityRouter
	SuperBuilderRouter
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
