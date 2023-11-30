package business

type ServiceGroup struct {
	// golang_coder Begin; DO NOT EDIT.
	SysApisService
	SysMenuService

	MemUserService
	MemLogsService
	MemDepartService
	MemUserSafeService
	BasicFileService
	BasicDictService
	BasicFileCatService
	// golang_coder End; DO NOT EDIT.
}
