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
						 
							<el-button class="el-btn-save" type="primary" @click="goEditForm(0)">新增</el-button>
						<!-- 	<el-button class="el-btn-save" type="primary" @click="deleteMultiRow">删除</el-button> -->
						 </el-form-item>
					</div>
				</div>
 
				</el-form>

			<!----------数据表------------------ -->
			<el-table row-key="id" ref="multipleTable" border  style="width: 100%" tooltip-effect="dark" :data="tableData" @selection-change="handleSelectionChange" @sort-change="sortChange" >
 				<!-- <el-table-column type="selection" width="55" /> -->
				<el-table-column label="序号" width="80" prop="id" sortable="custom" />
					<el-table-column label="用户名" prop="username" min-width="120"   sortable="custom"  />
					<el-table-column label="真实名" prop="realname" min-width="120"   sortable="custom"  />
         	<el-table-column label="性别" prop="sex" min-width="120"  sortable="custom" >
            <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">
            	<el-select v-model="scope.row.sex" placeholder="请选择"  @change="quickEdit_do('sex',scope.row.id,scope.row.sex,scope)">
            		<el-option v-for="(item,key) in sex_options" :key="key" :label="item.label" :value="item.value"></el-option>
            	</el-select>
            	<template #reference>
            		<div class="quickEdit" > {{filterDict(scope.row.sex,sex_options)}} </div>
            	</template>
            	</el-popover>
            </template>
            </el-table-column>
				<!-- 	<el-table-column label="角色列表" prop="roleList" min-width="120"   sortable="custom"  /> -->
					<!-- <el-table-column label="角色" prop="role" min-width="120"   sortable="custom"  /> -->
          <el-table-column label="部门" prop="departId" min-width="120"   sortable="custom"  >
            <template #default="scope">
              	{{getTreeName(treeOptions,scope.row.departId)}}
              </template>
            </el-table-column>
        <!--  	<el-table-column label="备注" prop="remark" min-width="120"   sortable="custom"  /> -->
          <el-table-column label="创建人" prop="userId" min-width="150" sortable="custom" >
            <template #default="scope">
            	{{filterDict(scope.row.userId,userName_options)}}
            </template>
          </el-table-column>
          <el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
          	<template #default="scope">{{formatDate(scope.row.createdAt,1)}}</template>
          </el-table-column>

					<el-table-column label="状态" prop="status"  width="90"  sortable="custom" >
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
				<el-table-column label="编辑" width="150">
					<template #default="scope">
            <el-tooltip placement="top-end" effect="dark" content="扫码登录"> <el-button type="primary" link
            	 :icon="useRenderIcon('ep:camera-filled')" 	@click="showQrCodeDo(scope.row.id)">
            	</el-button> </el-tooltip>
            <el-button :icon="useRenderIcon('ep:edit')" type="primary" link @click="goEditForm(scope.row.id)"/>
            <el-tooltip placement="top-end" effect="dark" content="修改密码"> <el-button :icon="useRenderIcon('ep:view')"
            		type="primary" link @click="pwdRow(scope.row.id)"></el-button> </el-tooltip>
            <el-button :icon="useRenderIcon('ep:delete')" type="primary" link @click="deleteRow(scope.row)" />
					</template>
				</el-table-column>

			</el-table>
		</div>
		<el-pagination class="gocms-pagination" layout="total, prev, pager, next, jumper, sizes" :current-page="page"
			:page-size="pageSize" :page-sizes="[10,20,50, 100]" :total="total" @current-change="handleCurrentChange"
			@size-change="handleSizeChange" />
	</div>


  <!-- <el-table-column label="身份证" prop="cardId" min-width="120"   sortable="custom"  />
  <el-table-column label="昵称" prop="nickname" min-width="120"   sortable="custom"  />
	<el-table-column label="手机" prop="mobile" min-width="120"   sortable="custom"  />
  <el-table-column label="邮件" prop="email" min-width="120"   sortable="custom"  />
  <el-table-column label="生日" width="180" prop="birthday"   sortable="custom" >
  			<template #default="scope">{{formatDate(scope.row.birthday)}}</template>
  		</el-table-column>
  	<el-table-column label="头像" prop="avatar" min-width="120"   sortable="custom" >
  		<template #default="scope">
  			 <FileListView :objList="getFileByGuidStr(scope.row.avatar,scope.row.fileObjList)" />
  		</template>
  	</el-table-column>
  	<el-table-column label="岗位" prop="jobId" min-width="120"   sortable="custom"  /> -->

  	<!-- <el-table-column label="验证手机" prop="mobileValidated" min-width="120"   sortable="custom"  >
  	<template #default="scope">{{formatBoolean(scope.row.mobileValidated)}}</template>
  	</el-table-column>
  	<el-table-column label="验证邮箱" prop="emailValidated" min-width="120"   sortable="custom"  >
  	<template #default="scope">{{formatBoolean(scope.row.emailValidated)}}</template>
  	</el-table-column>
  	<el-table-column label="验证实名" prop="cardidValidated" min-width="120"   sortable="custom"  >
  	<template #default="scope">{{formatBoolean(scope.row.cardidValidated)}}</template>
  	</el-table-column> -->

  <!-- 	<el-table-column label="安全状态" prop="statusSafe" min-width="120"  sortable="custom" >
  	<template #default="scope">
  	<el-popover trigger="click" placement="top"  width = "280">
  		<el-select v-model="scope.row.statusSafe" placeholder="请选择"  @change="quickEdit_do('status_safe',scope.row.id,scope.row.statusSafe,scope)">
  			<el-option v-for="(item,key) in usersafe_type_options" :key="key" :label="item.label" :value="item.value"></el-option>
  		</el-select>
  		<template #reference>
  			<div class="quickEdit" > {{filterDict(scope.row.statusSafe,usersafe_type_options)}} </div>
  		</template>
  		</el-popover>
  	</template>
  	</el-table-column>
  	<el-table-column label="注册ip" prop="regIp" min-width="120"   sortable="custom"  />
  	<el-table-column label="登录ip" prop="loginIp" min-width="120"   sortable="custom"  />
  	<el-table-column label="登录次数" prop="loginTotal" min-width="120"   sortable="custom"  />
  		<el-table-column label="最后登录时间" width="180" prop="loginTime"   sortable="custom" >
  			<template #default="scope">{{formatDate(scope.row.loginTime)}}</template>
  		</el-table-column>
  	<el-table-column label="扫码" prop="scan" min-width="120"   sortable="custom"  />
  		<el-table-column label="扫码" width="180" prop="scanTime"   sortable="custom" >
  			<template #default="scope">{{formatDate(scope.row.scanTime)}}</template>
  		</el-table-column>
  	<el-table-column label="设置值" prop="setting" min-width="120"   sortable="custom"  />
  	<el-table-column label="远程协助模式" prop="rtcModel" min-width="120"   sortable="custom"  /> -->
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
	createMemUser,
	deleteMemUserByIds,
	updateMemUser,
	findMemUser,
	getMemUserList,
	quickEdit,
	excelList
 } from '@/api/memUser'

 import MemUserForm from './memUserForm.vue'
 import MemUserFormPwd from './memUserFormPwd.vue'
 import MemUserFormQR from './memUserFormQR.vue'
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
	const sortChange = ({ prop, order }) => {
		if (prop) {
			searchInfo.value.orderKey = prop
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
				const res = await deleteMemUserByIds(data)
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
				const res = await deleteMemUserByIds(data)
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
				h(MemUserForm, {
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
    searchInfo.value.userType = 1
		removeNullAttr(searchInfo.value)
		let paramData = {
			page: page.value,
			pageSize: pageSize.value,
			...searchInfo.value
		}
		if (paramData.createdAtBetween)
			delete paramData.createdAtBetween
		const res = await getMemUserList(paramData)
		if (res.code === 200) {
			tableData.value = res.data.list
			total.value = res.data.total
			page.value = res.data.page
			pageSize.value = res.data.pageSize
		} else message(res.msg, { type: "error"})
	}
  const treeOptions = ref([])
  const getTreeData = async () => {
    let treeDataReq = {
      table: "mem_depart",
      pidField: "id",
      nameField: "title",
      pidValue: 0
    }
    treeOptions.value = await getPidTreeData(treeDataReq)
  }
 const userName_options = ref([])
 const sex_options = ref([])
 const getOptionsData = async () => {
   status_options.value = await getDict('status')
   let dictReq = {
   	table: "mem_user",
   	pidField: "id",
   	nameField: "realname"
   }
   userName_options.value = await getPidData(dictReq)
   sex_options.value =  await getDict('sex')
 }


  const pwdRow = async (id) => {
   editId.value = id
     addDialog({
       title: "编 辑",
       fullscreenIcon: true,
       hideFooter: true,
       contentRenderer: ({
           options,
           index
         }) =>
         h(MemUserFormPwd, {
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


   const showQrCodeDo = async (id) => {
      editId.value = id
      addDialog({
        title: "登录二维码",
        width:400,
        fullscreenIcon: false,
        hideFooter: true,
        contentRenderer: ({
            options,
            index
          }) =>
          h(MemUserFormQR, {
            editId: editId.value,
            beChange: beChange.value,
            index: index,
            options: options,
            "onUpdate:editId": val => (editId.value = val),
            "onUpdate:beChange": val => (beChange.value = val)
          }),
        closeCallBack: () => {
          if (beChange.value) {
            // getTableData()
          }
        }
      });
    }



	const init = () => {
		getOptionsData()
		getTreeData()
		getTableData()
	}
	onMounted(() => {
     init()
   })
</script>

 <style lang="scss" scoped>
 </style>
