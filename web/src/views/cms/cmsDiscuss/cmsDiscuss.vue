<template>
	<div> 	 
		<!----------查询form--------- -->
		<div class="gocms-table-box bg-bg_color"> 
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
							<el-button class="el-btn-save" type="primary" :icon="searchToggle?useRenderIcon('ep:arrow-up-bold'):useRenderIcon('ep:arrow-down-bold')" @click="searchToggle=!searchToggle">筛选</el-button>					
							<el-button class="el-btn-save" type="primary" @click="goEditForm(0)">新增</el-button>
							<el-button class="el-btn-save" type="primary" @click="deleteMultiRow">删除</el-button>	
						 </el-form-item>
					</div>
				</div>

				<div v-if="searchToggle" class="gocms-box-search"> 
					<el-form-item label="创建时间">
						<el-date-picker v-model="searchInfo.createdAtBetween" type="datetimerange"
							format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :shortcuts="shortcuts" range-separator="至"
							start-placeholder="开始日期" end-placeholder="结束日期" />
					</el-form-item> 
							<el-form-item label="id">
								<el-input placeholder="搜索id" v-model="searchInfo.id" />
							</el-form-item>
								<el-form-item label="父ID">
									<el-input placeholder="搜索条件" v-model="searchInfo.pid" clearable />
								</el-form-item>
								<el-form-item label="id">
									<el-input placeholder="搜索条件" v-model="searchInfo.artId" clearable />
								</el-form-item>
								<el-form-item label="用户id">
									<el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
								</el-form-item>
								<el-form-item label="用户At">
									<el-input placeholder="搜索条件" v-model="searchInfo.userIdAt" clearable />
								</el-form-item> 
								<el-form-item label="标题">
								<el-input placeholder="搜索条件" v-model="searchInfo.title" clearable />
								</el-form-item> 
								<el-form-item label="内容">
								<el-input placeholder="搜索条件" v-model="searchInfo.detail" clearable />
								</el-form-item>
								<el-form-item label="总赞">
									<el-input placeholder="搜索条件" v-model="searchInfo.totalGood" clearable />
								</el-form-item>
								<el-form-item label="总踩">
									<el-input placeholder="搜索条件" v-model="searchInfo.totalPoor" clearable />
								</el-form-item>
								<el-form-item label="状态" prop="status">                
									<el-select v-model="searchInfo.status" placeholder="请选择" clearable>
									<el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value"></el-option>
									</el-select>
								</el-form-item>  
                    </div> 
				</el-form>
		 
			<!----------数据表------------------ -->
			<el-table row-key="id" ref="multipleTable" border  style="width: 100%" tooltip-effect="dark" :data="tableData" @selection-change="handleSelectionChange" @sort-change="sortChange" >
 				<el-table-column type="selection" width="55" />
				<el-table-column label="序号" width="80" prop="id" sortable="custom" /> 
					<el-table-column label="id" prop="artId" min-width="120"   sortable="custom"  /> 
					<el-table-column label="用户id" prop="userId" min-width="120"   sortable="custom"  /> 
					<el-table-column label="用户At" prop="userIdAt" min-width="120"   sortable="custom"  /> 
					<el-table-column label="标题" prop="title" min-width="120"   sortable="custom"  />
					<el-table-column label="图片" prop="fileList" min-width="120"   sortable="custom" >
						<template #default="scope"> 						    
							 <FileListView :objList="getFileByGuidStr(scope.row.fileList,scope.row.fileObjList)" />  
						</template>
					</el-table-column> 
					<el-table-column label="总赞" prop="totalGood" min-width="120"   sortable="custom"  /> 
					<el-table-column label="总踩" prop="totalPoor" min-width="120"   sortable="custom"  />
					<el-table-column label="状态" prop="status" min-width="120"  sortable="custom" >
					<template #default="scope">  
					<el-popover trigger="click" placement="top"  width = "280">  
						<el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.id,scope.row.status,scope)">
							<el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value"></el-option>
						</el-select> 
						<template #reference>
							<div class="quickEdit" > {{filterDict(scope.row.status,status_options)}} </div>
						</template>
						</el-popover>
					</template>  
					</el-table-column> 

				<el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
					<template #default="scope">{{formatDate(scope.row.createdAt,1)}}</template>
				</el-table-column> 

				<el-table-column label="编辑" width="80"  fixed="right">
					<template #default="scope">
			              <el-button :icon="useRenderIcon('ep:edit')" type="primary" link @click="goEditForm(scope.row.id)"/>
                          <el-button :icon="useRenderIcon('ep:delete')" type="primary" link @click="deleteRow(scope.row)" />
					</template>
				</el-table-column>
			</el-table>
		</div>
		<el-pagination class="gocms-pagination" layout="total, prev, pager, next, jumper, sizes" :current-page="page"
			:page-size="pageSize" :page-sizes="[10,20,50, 100]" :total="total" @current-change="handleCurrentChange"
			@size-change="handleSizeChange" />  
		 
	</div>
</template>
 
<script setup lang="tsx">
   import {
     ref,
     onMounted,
     computed,
     h,
     createVNode,
     type CSSProperties,
   } from "vue";
   import {
     message
   } from "@/utils/message";
   import {
     storageLocal
   } from "@pureadmin/utils";
   import {
     useUserStoreHook
   } from "@/store/modules/user";
   import {
     usePermissionStoreHook
   } from "@/store/modules/permission";
   import {
     useRenderIcon
   } from "@/components/ReIcon/src/hooks";
   import {
     ElMessage,
     ElMessageBox
   } from 'element-plus'

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
     substr,
     filterDict,
     formatDate,
     formatBoolean
   } from '@/utils/utils'
   

   import {
     addDialog,
     closeDialog,
     updateDialog,
     closeAllDialog
   } from "@/components/ReDialog";


 import {
	createCmsDiscuss,
	deleteCmsDiscussByIds,
	updateCmsDiscuss,
	findCmsDiscuss,
	getCmsDiscussList,
	quickEdit,
	excelList
 } from '@/api/cmsDiscuss'

 import CmsDiscussForm from './cmsDiscussForm.vue'


   const page = ref(1)
   const total = ref(0)
   const pageSize = ref(20)
   const tableData = ref([])
   const searchInfo = ref({})
   const searchToggle = ref(false)
   const beChange = ref(false) 
   const editId = ref(0)
   const multipleSelection = ref([])  

    // 字典 
		const status_options = ref([])
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
				const res = await deleteCmsDiscussByIds(data)
				if (res.code === 200) {
					message(res.msg, { type: "success" })
					if (tableData.value.length === 1 && page.value > 1) {
						page.value--
					}
					getTableData()
				} else message(res.msg, { type: "error" })
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
			message("请先选择删除项", { type: "error" })			 
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
				const res = await deleteCmsDiscussByIds(data)
				if (res.code === 200) {
					message(res.msg, { type: "success" })
					if (tableData.value.length === 1 && page.value > 1) {
						page.value--
					}
					getTableData()
				} else message(res.msg, { type: "error" })
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
			message(res.msg, { type: "success" })
		else
			message(res.msg, { type: "error" })
		// if (scope._self.$refs[`popover-${scope.$index}`])
		// scope._self.$refs[`popover-${scope.$index}`].doClose();
	}

	//编辑
	const goEditForm = (id) => {
	   editId.value = id
	   addDialog({
			title: "编 辑",
			fullscreenIcon: true,
			hideFooter: true,
			contentRenderer: ({ options, index }) =>
				h(CmsDiscussForm, {
				editId: editId.value,
				beChange: beChange.value,
				index: index,
				options: options,
				"onUpdate:editId": val => (editId.value = val),
				"onUpdate:beChange": val => (beChange.value = val)
				}),
			closeCallBack: () => {
				if (beChange.value) {
				    getTableData()
				}
			}
		}); 
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
		const res = await getCmsDiscussList(paramData)
		if (res.code === 200) {
			tableData.value = res.data.list
			total.value = res.data.total
			page.value = res.data.page
			pageSize.value = res.data.pageSize
		} else message(res.msg, { type: "error"})
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
			status_options.value = await getDict('status') 
		//sexOptions.value = await getDict('sex')
		//statusOptions.value = await getDict('status')
		//usersafe_typeOptions.value = await getDict('usersafe_type')
	} 
	const init = () => {
		getOptionsData()
		//getTreeData()
		getTableData()
	}
	onMounted(() => {
     init()
   })
</script>

 <style lang="scss" scoped>
 </style>
 