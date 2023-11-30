package business

import commSev "go-cms/service/common"

type ApiGroup struct {
	commSev.BaseApi
	// golang_coder Begin; DO NOT EDIT.
	BasicFileCatApi
	BasicDictApi
	BasicFileApi

	MemUserApi
	MemDepartApi
	MemLogsApi
	MemUserSafeApi
	SysMenuApi
	SysApisApi
	SysRoleApi
	SysRoleMenuApi
	SysLogsApi

	CmsDetailApi
	CmsVisitorApi
	CmsDiscussApi
	CmsCatArtApi
	CmsCatApi
	CmsBlogApi
	CmsGroupApi
	CmsTagApi
	CmsArtApi
	// golang_coder End; DO NOT EDIT.
	// 新增

}
