 <template>
   <div>
     <!----------查询form------------------ -->
     <div class="gocms-table-box bg-bg_color">
       <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
         <div class="gocms-box-search-button">
           <div>
             <el-form-item label="ID">
               <el-input placeholder="搜索ID" v-model="searchInfo.id" clearable />
             </el-form-item>
             <el-form-item label="名称">
               <el-input placeholder="搜索名称/标题" v-model="searchInfo.title" clearable />
             </el-form-item>
           </div>
           <div>
             <el-form-item>
               <el-button class="el-btn-save" type="primary" @click="onSearch">查询</el-button> 
               <el-button class="el-btn-save" type="primary" @click="goEditForm(0)">新增</el-button>
            
             </el-form-item>
           </div>
         </div>

        
       </el-form>

       <!----------数据表------------------ -->
       <el-table row-key="id" ref="multipleTable" border style="width: 100%" tooltip-effect="dark" :data="tableData"
         @selection-change="handleSelectionChange" @sort-change="sortChange">

         <el-table-column label="序号" width="120" prop="id" sortable="custom" />
         <!--  <el-table-column label="商户" prop="cuId" min-width="120" sortable="custom" /> -->
         <!--  <el-table-column label="父菜单ID" prop="pid" min-width="120" sortable="custom" /> -->
         <el-table-column label="标题" prop="title" min-width="120" sortable="custom" />
         <el-table-column label="name" prop="name" min-width="120" sortable="custom" />
         <el-table-column label="path" prop="path" min-width="120" sortable="custom" />
         <!-- <el-table-column label="等级" prop="level" min-width="120" sortable="custom" />
         <el-table-column label="按钮权限" prop="auths" min-width="120" sortable="custom" /> -->

         <!--  <el-table-column label="活动链接" prop="activePath" min-width="120" sortable="custom" />
         <el-table-column label="外链" prop="frameSrc" min-width="120" sortable="custom" /> -->
         <el-table-column label="前端文件" prop="component" min-width="160" sortable="custom">
           <template #default="scope">
             <el-popover trigger="click" placement="top" width="280">
               <el-row :gutter="10">
                 <el-col :span="16">
                   <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.component">
                   </el-input>
                 </el-col>
                 <el-col :span="4">
                   <el-button size="small" type="primary" icon="el-icon-edit" class="table-button"
                     @click="quickEdit_do('component',scope.row.id,scope.row.component,scope)">保存</el-button>
                 </el-col>
               </el-row>
               <template #reference>
                 <div class="quickEditTxt"> {{scope.row.component}} </div>
               </template>
             </el-popover>
           </template>
         </el-table-column>
         <el-table-column label="排序" prop="sort" min-width="120" sortable="custom">
           <template #default="scope">
             <el-popover trigger="click" placement="top" width="280">
               <el-row :gutter="10">
                 <el-col :span="16">
                   <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sort">
                   </el-input>
                 </el-col>
                 <el-col :span="4">
                   <el-button size="small" type="primary" icon="el-icon-edit" class="table-button"
                     @click="quickEdit_do('sort',scope.row.id,scope.row.sort,scope)">保存</el-button>
                 </el-col>
               </el-row>
               <template #reference>
                 <div class="quickEditTxt"> {{scope.row.sort}} </div>
               </template>
             </el-popover>
           </template>
         </el-table-column>

         <!--      <el-table-column label="保持连接" prop="keepAlive" min-width="120" sortable="custom">
           <template #default="scope">{{formatBoolean(scope.row.keepAlive)}}</template>
         </el-table-column> -->
         <el-table-column label="图标" prop="icon" min-width="120" sortable="custom" />
         <el-table-column label="显示" prop="showLink" min-width="120" sortable="custom">
           <template #default="scope">
             <el-switch v-model="scope.row.showLink"
               @change="quickEdit_do('show_link',scope.row.id,scope.row.showLink,scope)" />
           </template>
         </el-table-column>
         <el-table-column label="参数" prop="param" min-width="60" sortable="custom" />
         <el-table-column label="状态" prop="status" min-width="120" sortable="custom">
           <template #default="scope">
             <el-popover trigger="click" placement="top" width="280">
               <el-select v-model="scope.row.status" placeholder="请选择"
                 @change="quickEdit_do('status',scope.row.id,scope.row.status,scope)">
                 <el-option v-for="(item,key) in status_options" :key="key" :label="item.label"
                   :value="item.value"></el-option>
               </el-select>+
               <template #reference>
                 <div class="quickEdit"> {{filterDict(scope.row.status,status_options)}} </div>
               </template>
             </el-popover>
           </template>
         </el-table-column>

         <!--     <el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
           <template #default="scope">{{formatDate(scope.row.createdAt,1)}}</template>
         </el-table-column> -->

         <el-table-column label="编辑" width="100">
           <template #default="scope">
             <el-button :icon="useRenderIcon('ep:edit')" type="primary" link @click="goEditForm(scope.row.id)"/>
             <el-button :icon="useRenderIcon('ep:delete')" type="primary" link @click="deleteRow(scope.row)" />
           </template>
         </el-table-column>

       </el-table>
     </div>
     <!--    <el-pagination class="gocms-pagination" layout="total, prev, pager, next, jumper, sizes" :current-page="page"
       :page-size="pageSize" :page-sizes="[10,20,50, 100]" :total="total" @current-change="handleCurrentChange"
       @size-change="handleSizeChange" /> -->
     <!---------- 编辑弹窗------------------ -->
     <!--   <el-dialog width="60%" sappend-to-body="true" v-if="dialogFormVisible" v-model="dialogFormVisible"
       :show-close="false">
       <template #header="{ close, titleId, titleClass }">
         <div class="my-dialog-header">
           <span>编辑</span>
           <el-icon size="40" @click="closeDialog" class="my-dialog-closeIcon">
             <CircleCloseFilled />
           </el-icon>
         </div>
       </template>
        <SysMenuForm :beComponent="true" :editId="editId" @closeFormDialog="closeFormDialog" />
     </el-dialog> -->
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
     createSysMenu,
     deleteSysMenuByIds,
     updateSysMenu,
     findSysMenu,
     getSysMenuList,
     quickEdit,
     excelList
   } from '@/api/sysMenu'

   import {
     addDialog,
     closeDialog,
     updateDialog,
     closeAllDialog
   } from "@/components/ReDialog";

   import SysMenuForm from './sysMenuForm.vue'

   //const shortcuts = getShortcuts()
   const page = ref(1)
   const total = ref(0)
   const pageSize = ref(1000)
   const tableData = ref([])
   const searchInfo = ref({})
   const searchToggle = ref(false)
   const beChange = ref(false)

   // const dialogFormVisible = ref(false)
   // const beNewWindow = ref(false) //是否在新窗口打开编辑器
   const editId = ref(0)
   const multipleSelection = ref([])

   // 搜索
   const onSearch = () => {
     page.value = 1
     pageSize.value = 1000
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
         const res = await deleteSysMenuByIds(data)
         if (res.code === 200) {
           message(res.msg, {
             type: "success"
           })
           if (tableData.value.length === 1 && page.value > 1) {
             page.value--
           }
           getTableData()
         } else message(res.msg, {
           type: "error"
         })
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
         const res = await deleteSysMenuByIds(data)
         if (res.code === 200) {
           message(res.msg, {
             type: "success"
           })
           if (tableData.value.length === 1 && page.value > 1) {
             page.value--
           }
           getTableData()
         } else message(res.msg, {
           type: "error"
         })
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
     if (res.code == 200)
       message(res.msg, {
         type: "success"
       })
     else
       message(res.msg, {
         type: "error"
       })
     // if (scope._self.$refs[`popover-${scope.$index}`])
     // scope._self.$refs[`popover-${scope.$index}`].doClose();
   }

   //编辑
   const goEditForm = (id: number) => {
     editId.value = id
     addDialog({
       title: "编 辑",
       fullscreenIcon: true,
       hideFooter: true,
       contentRenderer: ({
           options,
           index
         }) =>
         h(SysMenuForm, {
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
     } else {
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
     const res = await getSysMenuList(paramData)
     if (res.code == 200) {
       tableData.value = res.data.list
       total.value = res.data.total
       page.value = res.data.page
       pageSize.value = res.data.pageSize
     } else message(res.msg, {
       type: "error"
     })
   }


   // 字典
   const status_options = ref([])
   const treeOptions = ref([])
   const getOptionsData = async () => {
     status_options.value = await getDict('status')
     //sexOptions.value = await getDict('sex')
     //statusOptions.value = await getDict('status')
     //usersafe_typeOptions.value = await getDict('usersafe_type')
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
