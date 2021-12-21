package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
	SuperBuilderService
	SuperBuilderHistoryService
	BaseMenuService
	CasbinService
	DictionaryService
	DictionaryDetailService
	InitDBService
	MenuService
	OperationRecordService
	SystemConfigService
	UserService
}
