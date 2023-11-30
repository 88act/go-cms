import service from '@/utils/request'

// @Tags {{.StructName}}
// @Summary 创建{{.StructName}}
export const create{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/create{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
export const delete{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 删除{{.StructName}}
export const delete{{.StructName}}ByIds = (data) => {
  return service({
    url: '/{{.Abbreviation}}/delete{{.StructName}}ByIds',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 更新{{.StructName}}
export const update{{.StructName}} = (data) => {
  return service({
    url: '/{{.Abbreviation}}/update{{.StructName}}',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 用id查询{{.StructName}}
export const find{{.StructName}} = (params) => {
  return service({
    url: '/{{.Abbreviation}}/find{{.StructName}}',
    method: 'get',
    params
  })
}

// @Tags {{.StructName}}
// @Summary 分页获取{{.StructName}}列表
export const get{{.StructName}}List = (params) => {
  return service({
    url: '/{{.Abbreviation}}/get{{.StructName}}List',
    method: 'get',
    params
  })
}


// @Tags {{.StructName}}
// @Summary 快速编辑
export const quickEdit = (data) => {
  return service({
    url: '/{{.Abbreviation}}/quickEdit',
    method: 'post',
    data
  })
}

// @Tags {{.StructName}}
// @Summary 导出excel{{.StructName}}列表
export const excelList = (params) => {
  return service({
    url: '/{{.Abbreviation}}/excelList',
    method: 'get',
    params
  })
}