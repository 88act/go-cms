<template>
	<div>
		<!----------查询form--------- -->
		<div class="gocms-table-box">
				<el-form :inline="true" :model="searchInfo" class="demo-form-inline">
				<div class="gocms-box-search-button">
					<div>
						<el-form-item label="GUID">
							<el-input placeholder="搜索guid" v-model="searchInfo.guid"  clearable />
						</el-form-item>
						<el-form-item label="名称">
							<el-input placeholder="搜索名称/标题" v-model="searchInfo.name"  style="width:300px;" clearable />
						</el-form-item>
					</div>
					<div>
						<el-form-item  >
							<el-button class="el-btn-save" type="primary" @click="onSearch">查询</el-button>						 
							<el-button class="el-btn-save" type="primary" @click="deleteMultiRow">删除</el-button>
						 </el-form-item>
					</div>
				</div>

				</el-form>

			<!----------数据表------------------ -->
			<el-table row-key="id" ref="multipleTable" border  style="width: 100%" tooltip-effect="dark" :data="tableData" @selection-change="handleSelectionChange" @sort-change="sortChange" >
 				<el-table-column type="selection" width="55" />
				 <el-table-column label="序号" width="80" prop="id" sortable="custom" />
					<el-table-column label="唯一id" prop="guid" min-width="120"   sortable="custom"  />
					<el-table-column label="用户id" prop="userId" min-width="120"   sortable="custom"  />
					<el-table-column label="文件名" prop="name" min-width="120"   sortable="custom"  />
					<el-table-column label="文件路径" prop="path" min-width="120"   sortable="custom"  />
					<el-table-column label="文件类型" prop="ext" min-width="120"   sortable="custom"  />
					<el-table-column label="文件大小[k]" prop="size" min-width="120"   sortable="custom"  />
					<el-table-column label="md5值" prop="md5" min-width="120"   sortable="custom"  />

				<el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
					<template #default="scope">{{formatDate(scope.row.createdAt,1)}}</template>
				</el-table-column>
				<el-table-column label="编辑" width="100">
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
	createBasicFile,
	deleteBasicFileByIds,
	updateBasicFile,
	findBasicFile,
	getBasicFileList,
	quickEdit,
	excelList
 } from '@/api/basicFile'

 import BasicFileForm from './basicFileForm.vue'


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
		const media_type_options = ref([])
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
				const res = await deleteBasicFileByIds(data)
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
				const res = await deleteBasicFileByIds(data)
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
				h(BasicFileForm, {
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
		const res = await getBasicFileList(paramData)
		if (res.code === 200) {
			tableData.value = res.data.list
			total.value = res.data.total
			page.value = res.data.page
			pageSize.value = res.data.pageSize
		} else message(res.msg, { type: "error"})
	}

	const getOptionsData = async () => {
			media_type_options.value = await getDict('media_type')
			status_options.value = await getDict('status')

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
