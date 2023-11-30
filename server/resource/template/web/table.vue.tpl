<template>
	<div> 	 
		<!----------查询form------------------ -->
		<div class="gocms-table-box"> 
				<el-form :inline="true" :model="searchInfo" class="demo-form-inline">
				<div class="gocms-box-search-button">  
					<div>
						<el-form-item label="ID">
							<el-input placeholder="搜索ID" v-model="searchInfo.id"  clearable />
						</el-form-item>
						<el-form-item label="名称">
							<el-input placeholder="搜索名称/标题" v-model="searchInfo.title"  style="width:300px;" clearable />
						</el-form-item>
					</div>
					<div>
						<el-form-item  >
							<el-button class="el-btn-save" type="primary" @click="onSearch">查询</el-button>			
							<el-button class="el-btn-save" type="primary" :icon="searchToggle?'ArrowUp':'ArrowDown'" @click="searchToggle=!searchToggle">筛选</el-button>						
							<el-button class="el-btn-save" type="primary" @click="goEditForm(0)">新增</el-button>
							<el-button class="el-btn-save" type="primary" @click="deleteMultiRow">删除</el-button>	
						 </el-form-item>
					</div>
				</div>

				<div v-if="searchToggle" class="gocms-box-search">					 
				{{- if .SearchCreate}} 
					<el-form-item label="创建时间">
						<el-date-picker v-model="searchInfo.createdAtBetween" type="datetimerange"
							format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :shortcuts="shortcuts" range-separator="至"
							start-placeholder="开始日期" end-placeholder="结束日期" />
					</el-form-item>		
					{{- end }} 
						{{- if .SearchId}} 
							<el-form-item label="id">
								<el-input placeholder="搜索id" v-model="searchInfo.id" />
							</el-form-item>
					{{- end -}} 
					
				{{- range .Fields}} 
					{{- if .FieldSearchType}}   
						{{- if eq .FieldType "bool" }}
							<el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
							<el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="请选择">
								<el-option
									key="true"
									label="是"
									value="true">
								</el-option>
								<el-option
									key="false"
									label="否"
									value="false">
								</el-option>
							</el-select>
						</el-form-item> 
						{{- else if eq .FieldType "int" }}
							{{- if .DictType }}
								<el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">                
									<el-select v-model="searchInfo.{{ .FieldJson }}" placeholder="请选择" clearable>
									<el-option v-for="(item,key) in {{ .DictType }}_options" :key="key" :label="item.label" :value="item.value"></el-option>
									</el-select>
								</el-form-item>
							{{- else }}
								<el-form-item label="{{.FieldDesc}}">
									<el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}" clearable />
								</el-form-item>
							{{- end }}      
							{{- else if eq .FieldType "time.Time" }}
							<el-form-item label="{{.FieldDesc}}"> 
								<el-date-picker
								v-model="searchInfo.{{ .FieldJson }}"  
								type="datetimerange"
								format="YYYY-MM-DD HH:mm:ss"
								value-format="YYYY-MM-DD HH:mm:ss" :shortcuts="shortcuts" range-separator="至"
								start-placeholder="开始日期"
								end-placeholder="结束日期"
							/> 
							</el-form-item>
							{{- else }} 
								<el-form-item label="{{.FieldDesc}}">
								<el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}" clearable />
								</el-form-item>
							{{- end }}
						{{- end }}  
						{{- end }}  
                    </div> 
				</el-form>
		 
			<!----------数据表------------------ -->
			<el-table row-key="id" ref="multipleTable" border  style="width: 100%" tooltip-effect="dark" :data="tableData" @selection-change="handleSelectionChange" @sort-change="sortChange" >
 				<el-table-column type="selection" width="55" />
				<el-table-column label="序号" width="80" prop="id" sortable="custom" /> 
			{{- range .Fields}} 
			{{- if .BeHide }} 				         
			{{- else}}  
				{{- if .BeQuickEdit}} 
					{{- if .DictType}}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120" {{if .OrderBy}} sortable="custom"{{end}} >
					<template #default="scope">  
					<el-popover trigger="click" placement="top"  width = "280">  
						<el-select v-model="scope.row.{{.FieldJson}}" placeholder="请选择"  @change="quickEdit_do('{{.ColumnName}}',scope.row.id,scope.row.{{.FieldJson}},scope)">
							<el-option v-for="(item,key) in {{.DictType}}_options" :key="key" :label="item.label" :value="item.value"></el-option>
						</el-select> 
						<template #reference>
							<div class="quickEdit" > {{"{{"}}filterDict(scope.row.{{.FieldJson}},{{.DictType}}_options){{"}}"}} </div>
						</template>
						</el-popover>
					</template>  
					</el-table-column>
					{{- else if eq .FieldType "bool" }}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120"  {{if .OrderBy}} sortable="custom"{{end}}  >                        
						<template #default="scope" ><el-switch v-model="scope.row.{{.FieldJson}}" @change="quickEdit_do('{{.ColumnName}}',scope.row.id,scope.row.{{.FieldJson}},scope)"/></template> 
					</el-table-column> 
					{{- else if eq .FieldType "image" }}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}} >						 
							<template #default="scope">
						        <FileListView :objList="getFileByGuidStr(scope.row.{{.FieldJson}},scope.row.fileObjList)" />
					        </template>					 
					</el-table-column>
					{{-  else  if eq .FieldType "time.Time"}}
						<el-table-column label="{{.FieldDesc}}" width="180" prop="{{.FieldJson}}"  {{if .OrderBy}} sortable="custom"{{end}} >
							<template #default="scope">{{"{{formatDate(scope.row."}}{{.FieldJson}}{{")}}"}}</template>
						</el-table-column> 
					{{- else }} 
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120"  {{if .OrderBy}} sortable="custom"{{end}} >
					<template #default="scope">
						<el-popover trigger="click" placement="top" width="300">  
						<el-row :gutter="4">
						<el-col :span="19">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.{{.FieldJson}}"></el-input></el-col>
						<el-col :span="5"> <el-button size="small" type="primary"  class="el-btn-save" @click="quickEdit_do('{{.ColumnName}}',scope.row.id,scope.row.{{.FieldJson}},scope)">保存</el-button> </el-col> 
						</el-row>  
						<template #reference>
							<div  class="quickEditTxt"  > {{"{{"}}scope.row.{{.FieldJson}}{{"}}"}} </div>
						</template>
						</el-popover>
					</template>
					</el-table-column>              
					{{- end -}}
				{{- else}}  
					{{- if .DictType}}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120" {{if .OrderBy}} sortable="custom"{{end}} >
					<template #default="scope">  
						{{"{{"}}filterDict(scope.row.{{.FieldJson}},{{.DictType}}_options){{"}}"}}
					</template>
					</el-table-column>
					{{- else if eq .FieldType "bool" }}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120"  {{if .OrderBy}} sortable="custom"{{end}}  >
					<template #default="scope">{{"{{formatBoolean(scope.row."}}{{.FieldJson}}{{")}}"}}</template>
					</el-table-column>
					{{- else if eq .FieldType "image" }}
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120"  {{if .OrderBy}} sortable="custom"{{end}} >
						<template #default="scope"> 						    
							 <FileListView :objList="getFileByGuidStr(scope.row.{{.FieldJson}},scope.row.fileObjList)" />  
						</template>
					</el-table-column>
					{{-  else  if eq .FieldType "time.Time"}}
						<el-table-column label="{{.FieldDesc}}" width="180" prop="{{.FieldJson}}"  {{if .OrderBy}} sortable="custom"{{end}} >
							<template #default="scope">{{"{{formatDate(scope.row."}}{{.FieldJson}}{{")}}"}}</template>
						</el-table-column> 
					{{- else }} 
					<el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" min-width="120"  {{if .OrderBy}} sortable="custom"{{end}}  />
					{{- end }}  
				{{- end}} 
			{{- end}} 
			{{- end}} 

				<el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
					<template #default="scope">{{`{{formatDate(scope.row.createdAt,1)}}`}}</template>
				</el-table-column> 

				<el-table-column label="编辑" width="100">
					<template #default="scope">
						<el-button icon="delete" type="primary" link @click="deleteRow(scope.row)"></el-button>
						<el-button icon="edit" type="primary" link @click="goEditForm(scope.row.id)"></el-button>
					</template>
				</el-table-column>

			</el-table>
		</div>
		<el-pagination class="gocms-pagination" layout="total, prev, pager, next, jumper, sizes" :current-page="page"
			:page-size="pageSize" :page-sizes="[10,20,50, 100]" :total="total" @current-change="handleCurrentChange"
			@size-change="handleSizeChange" /> 

		<!---------- 编辑弹窗------------------ -->
		<el-dialog width="600px" sappend-to-body="true" v-if="dialogFormVisible" v-model="dialogFormVisible"
			:show-close="false">
			<template #header="{ close, titleId, titleClass }">
				<div class="my-dialog-header">
					<span>编辑</span>
					<el-icon size="40" @click="closeDialog" class="my-dialog-closeIcon">
						<CircleCloseFilled />
					</el-icon>
				</div>
			</template>
			<{{.StructName}}Form :beComponent="true" :editId="editId" @closeFormDialog="closeFormDialog" />
		</el-dialog>
	</div>
</template>
 
<script setup>
	import { 
		ref
	} from 'vue'
	import {
		getDict,
		getPidData,
		getPidTreeData,
		getDictNew,
		getDictNew2,
		getDictTreeNew,
		getTreeName,
		getTreeFullPath,
		getOptLabel
	} from '@/utils/dictionary' 
	import {
        isEmpty,
		obj2Num,
		removeNullAttr,
		get7days,
		getShortcuts,	
		getDataTimeStr,
		getFileByGuid,
		getFileByGuidStr,
		getFileByGuidList,
		substr
	} from '@/utils/utils'   
	import {
		formatDate,
		filterDict,
		getDictFunc,
		formatBoolean
	} from '@/utils/format'
	import {
		updatePidSort,
	} from '@/api/common_db'

	import {
		toSQLLine
	} from '@/utils/stringFun'
	import WarningBar from '@/components/warningBar/warningBar.vue'

	import {
		ElMessage,
		ElMessageBox
	} from 'element-plus'

	import {
		useRoute,
		useRouter
	} from 'vue-router'
	const router = useRouter()
	const route = useRoute()


	import {{.StructName}}Form from '@/view/{{.Module}}/{{.Abbreviation}}/{{.Abbreviation}}Form.vue'

	import MyStore from "@/pinia/modules/myStore"
	const myStore = MyStore()

	import {
		create{{.StructName}},
		delete{{.StructName}}ByIds,
		update{{.StructName}},
		find{{.StructName}},
		get{{.StructName}}List,
		quickEdit,
		excelList
	} from '@/api/{{.Abbreviation}}'

    const shortcuts = getShortcuts()
	const page = ref(1)
	const total = ref(0)
	const pageSize = ref(20)
	const tableData = ref([])
	const searchInfo = ref({})
    const searchToggle = ref(false)
 
	const dialogFormVisible = ref(false)
	const beNewWindow = ref(false) //是否在新窗口打开编辑器  
	const opt = ref("") //opt =add /del 操作类型
	const editId = ref(0)
	const multipleSelection = ref([]) 
 
    // 字典
	{{- range .Fields}}
		{{- if .DictType }} 
		const {{ .DictType }}_options = ref([])
		{{- end}}
	{{- end}}
	// 搜索 
	const onSearch = () => {
		page.value = 1
		pageSize.value = 20
		getTableData()
	}

	// 分页
	const handleSizeChange = (val) => {
		pageSize.value = val
		getTableData()
	}

	const handleCurrentChange = (val) => {
		page.value = val
		getTableData()
	}

	// 排序
	const sortChange = ({
		prop,
		order
	}) => {
		if (prop) {
			if (prop === 'id')
				prop = 'id';
			searchInfo.value.orderKey = toSQLLine(prop)
			searchInfo.value.orderDesc = order === 'descending'
		}
		getTableData()
	}

	//删除   
	const deleteRow = async (row) => {
		ElMessageBox.confirm('确认删除?', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				//type: 'warning'
			})
			.then(async () => {
				let data = {
					"ids": [row.id]
				}
				const res = await delete{{.StructName}}ByIds(data)
				if (res.code === 200) {
					ElMessage.success(res.msg)
					if (tableData.value.length === 1 && page.value > 1) {
						page.value--
					}
					getTableData()
				} else ElMessage.error(res.msg)
			})
	}
	//批量多选
	const handleSelectionChange = (val) => {
		multipleSelection.value = val
	}
	const deleteMultiRow = async () => {
		const ids = []
		multipleSelection.value && multipleSelection.value.forEach(item => {
			if (item.id > 0)
				ids.push(item.id)
		})
		if (ids.length == 0) {
			ElMessage.error("请先选择删除项");
			return;
		}
		ElMessageBox.confirm('确认删除?', '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				//type: 'warning'
			})
			.then(async () => {
				let data = {
					"ids": ids
				}
				const res = await delete{{.StructName}}ByIds(data)
				if (res.code === 200) {
					ElMessage.success(res.msg)
					if (tableData.value.length === 1 && page.value > 1) {
						page.value--
					}
					getTableData()
				} else ElMessage.error(res.msg)
			})
	}

	//快速编辑 
	const quickEdit_do = async (field, id, value, scope) => {
		let value2 = value;
		if (typeof(value) === "boolean")
			value2 = value ? "1" : "0"
		value2 = value2 + "";
		let obj = {
			field: field,
			id: id,
			value: value2
		}
		const res = await quickEdit(obj)
		if (res.code === 200)
			ElMessage.success(res.msg)
		else
			ElMessage.error(res.msg)
		// if (scope._self.$refs[`popover-${scope.$index}`])
		// scope._self.$refs[`popover-${scope.$index}`].doClose();
	}

	//编辑
	const goEditForm = (id) => {
		editId.value = id
		if (beNewWindow.value == true) {
			router.push({
				name: '{{.Abbreviation}}Form',
				params: {
					id: id
				}
			})
		} else {
			dialogFormVisible.value = true
		}
	}
	//关闭dialog
	const closeDialog = () => {
		dialogFormVisible.value = false
	}
	const closeFormDialog = (beGetData) => {
		console.log("closeFormDialog   beGetData=", beGetData)
		dialogFormVisible.value = false
		getTableData()
	}

	// 查询
	const getTableData = async () => {
        // 时间范围
		if (searchInfo.value.createdAtBetween && searchInfo.value.createdAtBetween.length >= 2) {
			searchInfo.value.createdAtBegin = searchInfo.value.createdAtBetween[0]
			searchInfo.value.createdAtEnd = searchInfo.value.createdAtBetween[1]
		}else {
			searchInfo.value.createdAtBegin = null
			searchInfo.value.createdAtEnd = null
		}	
		removeNullAttr(searchInfo.value)
		let paramData = {
			page: page.value,
			pageSize: pageSize.value,
			...searchInfo.value
		}
		if (paramData.createdAtBetween)
			delete paramData.createdAtBetween 
		const res = await get{{.StructName}}List(paramData)
		if (res.code === 200) {
			tableData.value = res.data.list
			total.value = res.data.total
			page.value = res.data.page
			pageSize.value = res.data.pageSize
		} else ElMessage.error(res.msg)
	}
	const getTreeData = async () => {
		let treeDataReq = {
			table: "memBranch",
			pidField: "id",
			nameField: "title",
			pidValue: 0
		}
		treeOptions.value = await getPidTreeData(treeDataReq)
	}

	const getOptionsData = async () => { 
		{{- range .Fields}}
		{{- if .DictType }} 
			{{.DictType}}_options.value = await getDict('{{.DictType}}')
		{{- end}}
		{{- end}} 
		//sexOptions.value = await getDict('sex')
		//statusOptions.value = await getDict('status')
		//usersafe_typeOptions.value = await getDict('usersafe_type')
	} 
	const init = () => {
		getOptionsData()
		//getTreeData()
		getTableData()
	}
	init()
</script>
<style>
</style>
