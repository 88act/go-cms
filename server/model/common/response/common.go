package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

// 文件简单类型 add by ljd
type FileObj struct {
	Url       string `json:"url"`
	Guid      string `json:"Guid"`
	MediaType int    `json:"mediaType"`
	Ext       string `json:"ext"`
}

// 返回id
type IdResp struct {
	Id int64 `json:"id" form:"id"`
}
