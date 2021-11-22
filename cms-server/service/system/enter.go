package system

type ServiceGroup struct {
	JwtService
	ApiService
	AuthorityService
 
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
