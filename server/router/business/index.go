package business

type RouterGroup struct {
	// golang_coder Begin; DO NOT EDIT.
	SysApisRouter
	SysMenuRouter
	SysLogsRouter
	SysRoleRouter
	MemUserRouter
	MemLogsRouter
	MemDepartRouter
	MemUserSafeRouter
	BasicFileRouter
	BasicDictRouter
	BasicFileCatRouter
	CmsDetailRouter
	CmsVisitorRouter
	CmsDiscussRouter
	CmsCatArtRouter
	CmsCatRouter
	CmsBlogRouter
	CmsGroupRouter
	CmsTagRouter
	CmsArtRouter
	// golang_coder End; DO NOT EDIT.

}
