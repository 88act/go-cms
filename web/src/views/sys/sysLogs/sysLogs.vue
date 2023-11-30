<template>
  <div>
    <!----------查询form------------------ -->
    <div class="gocms-table-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <div class="gocms-box-search-button">
          <div>
            <el-form-item label="客户id">
              <el-input placeholder="客户id" v-model="searchInfo.cuId" style="width:100px;" clearable />
            </el-form-item>
            <el-form-item label="用户id">
              <el-input placeholder="用户id" v-model="searchInfo.userId" style="width:100px;" clearable />
            </el-form-item>
            <el-form-item label="请求状态">
              <el-input placeholder="请求状态" v-model="searchInfo.status" style="width:100px;" clearable />
             </el-form-item>

            <el-form-item label="路径">
              <el-input placeholder="路径" v-model="searchInfo.path" style="width:300px;" clearable />
            </el-form-item>
          </div>
          <div>
            <el-form-item>
              <el-button class="el-btn-save" type="primary" @click="onSearch">查询</el-button>
            </el-form-item>
          </div>
        </div>


      </el-form>

      <!----------数据表------------------ -->
      <el-table row-key="id" ref="multipleTable" border style="width: 100%" tooltip-effect="dark" :data="tableData"
        @selection-change="handleSelectionChange" @sort-change="sortChange">
        <el-table-column label="序号" width="80" prop="id" sortable="custom" />
        <el-table-column label="客户id" prop="cuId" min-width="120" sortable="custom" />
        <el-table-column label="用户id" prop="userId" min-width="120" sortable="custom" />
        <el-table-column label="方法" prop="method" min-width="120" sortable="custom" />
        <el-table-column label="路径" prop="path" min-width="120" sortable="custom" />
        <el-table-column label="请求" prop="body" width="80" sortable="custom" >
        <template #default="scope">
            <div>
              <el-popover v-if="scope.row.body" placement="right-end" trigger="hover">
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.body) }}</pre>
                </div>
                <template #reference>
                <!--  <el-icon style="cursor: pointer;" :icon="useRenderIcon('ep:edit')" ></el-icon> -->
                <el-button :icon="useRenderIcon('ep:memo')" type="primary" link ></el-button>
                </template>
              </el-popover>
              <span v-else>无</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="响应" prop="resp"  width="80" sortable="custom" >
          <template #default="scope">
            <div>
              <el-popover v-if="scope.row.resp" placement="right-end" trigger="hover">
                <div class="popover-box">
                  <pre>{{ fmtBody(scope.row.resp) }}</pre>
                </div>
                <template #reference>
                <!--  <el-icon style="cursor: pointer;" :icon="useRenderIcon('ep:edit')" ></el-icon> -->
                <el-button :icon="useRenderIcon('ep:list')" type="primary" link ></el-button>
                </template>
              </el-popover>
              <span v-else>无</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="延迟" prop="latency" min-width="120" sortable="custom" />
   <!--     <el-table-column label="代理" prop="agent" min-width="120" sortable="custom" /> -->
        <el-table-column label="信息" prop="errorMessage" min-width="120" sortable="custom" />

        <el-table-column label="请求ip" prop="ip" min-width="120" sortable="custom" />
        <el-table-column label="地址" prop="ipAddr" min-width="120" sortable="custom" />
        <el-table-column label="状态" prop="status" min-width="120" sortable="custom" />

        <el-table-column label="创建时间" width="120" prop="created_at" sortable="custom">
          <template #default="scope">{{formatDate(scope.row.createdAt,0)}}</template>
        </el-table-column>

        <el-table-column label="编辑" width="100">
          <template #default="scope">
           <!-- <el-button :icon="useRenderIcon('ep:delete')" type="primary" link
              @click="deleteRow(scope.row)"></el-button> -->
            <el-button :icon="useRenderIcon('ep:edit')" type="primary" link
              @click="goEditForm(scope.row.id)"></el-button>
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
    createSysLogs,
    deleteSysLogsByIds,
    updateSysLogs,
    findSysLogs,
    getSysLogsList,
    quickEdit,
    excelList
  } from '@/api/sysLogs'

  import {
    addDialog,
    closeDialog,
    updateDialog,
    closeAllDialog
  } from "@/components/ReDialog";

  import SysLogsForm from './sysLogsForm.vue'

  //const shortcuts = getShortcuts()
  const page = ref(1)
  const total = ref(0)
  const pageSize = ref(20)
  const tableData = ref([])
  const searchInfo = ref({})
  const searchToggle = ref(false)
  const beChange = ref(false)

  // const dialogFormVisible = ref(false)
  // const beNewWindow = ref(false) //是否在新窗口打开编辑器
  const editId = ref(0)
  const multipleSelection = ref([])

const fmtBody = (value:string) => {
  try {
    return JSON.parse(value)
  } catch (err) {
    return value
  }
}


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
        const res = await deleteSysLogsByIds(data)
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
        const res = await deleteSysLogsByIds(data)
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
        h(SysLogsForm, {
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
    const res = await getSysLogsList(paramData)
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
    //getOptionsData()
    //getTreeData()
    getTableData()
  }
  onMounted(() => {
    init()
  })
</script>




<style lang="scss" scoped>
.popover-box {
  background: #112435;
  color: #f08047;
  height: 600px;
  width: 420px;
  overflow: auto;
}
.popover-box::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
</style>
