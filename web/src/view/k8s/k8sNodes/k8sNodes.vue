<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item> 
                <el-form-item label="节点名称">
                  <el-input placeholder="搜索条件" v-model="searchInfo.nodeName" clearable />
                </el-form-item> 
                <el-form-item label="ip地址">
                  <el-input placeholder="搜索条件" v-model="searchInfo.ip" clearable />
                </el-form-item> 
                <el-form-item label="节点状态">
                  <el-input placeholder="搜索条件" v-model="searchInfo.status" clearable />
                </el-form-item> 
                <el-form-item label="节点角色">
                  <el-input placeholder="搜索条件" v-model="searchInfo.roles" clearable />
                </el-form-item> 
                <el-form-item label="节点标签">
                  <el-input placeholder="搜索条件" v-model="searchInfo.label" clearable />
                </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSearch">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="goEditForm(0)">新增</el-button>
           <el-button size="mini" type="success"  @click="onExcel">导出当前</el-button>
         <el-button size="mini" type="success"  @click="onExcelAll">导出全部</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
   <!----------数据表------------------ -->
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
      @sort-change="sortChange" 
    >
      <el-table-column type="selection" width="55" />      
       <el-table-column label="ID" min-width="60" prop="ID" sortable="custom" />
          <!--nodeName  BeQuickEdit -->  
        <el-table-column label="节点名称" prop="nodeName" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nodeName"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('node_name',scope.row.ID,scope.row.nodeName,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nodeName}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
          <el-table-column label="ip地址" prop="ip" width="120"   sortable="custom"  /> 
          <el-table-column label="节点状态" prop="status" width="120"   sortable="custom"  /> 
          <el-table-column label="节点角色" prop="roles" width="120"   sortable="custom"  /> 
          <el-table-column label="创建时间" prop="createTime" width="120"   sortable="custom"  /> 
          <el-table-column label="节点版本" prop="version" width="120"   sortable="custom"  /> 
          <el-table-column label="节点资源" prop="resource" width="120"   sortable="custom"  /> 
          <el-table-column label="节点标签" prop="label" width="120"   sortable="custom"  /> 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="goEditForm(scope.row.ID)">编辑</el-button>
          <el-button plain size="mini" type="danger" icon="el-icon-delete"  @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10,20,50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!---------- 编辑弹窗------------------ -->
    <el-dialog  v-if="dialogFormVisible"  :before-close="closeDialog" v-model="dialogFormVisible" title="编辑资料">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="节点名称:"> 
              <el-input v-model="formData.nodeName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ip地址:"> 
              <el-input v-model="formData.ip" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点状态:"> 
              <el-input v-model="formData.status" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点角色:"> 
              <el-input v-model="formData.roles" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:"> 
              <el-input v-model="formData.createTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点版本:"> 
              <el-input v-model="formData.version" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点资源:"> 
              <el-input v-model="formData.resource" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点标签:"> 
              <el-input v-model="formData.label" clearable placeholder="请输入" />
       </el-form-item>
     </el-form>
      <div slot="footer" class="el-dialog__footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="saveEditForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createK8sNodes, 
  deleteK8sNodesByIds,
  updateK8sNodes,
  findK8sNodes,
  getK8sNodesList,
  quickEdit,
  excelList
} from '@/api/k8sNodes' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'K8sNodes',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:false,//是否在新窗口打开编辑器
      listApi: getK8sNodesList,   
      excelListApi: excelList,
      formData: {
           nodeName: '',
           ip: '',
           status: '',
           roles: '',
           createTime: '',
           version: '',
           resource: '',
           label: '',
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getTableData()
  },
  methods: { 
    onSearch() {
      this.page = 1
      this.pageSize = 20 
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
         const ids = [row.ID] 
         this.doDelete(ids); 
        //this.deleteK8sNodes(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      this.doDelete(ids); 
    },
  	async doDelete(ids) { 
     const res = await deleteK8sNodesByIds({ ids })
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      } 
	},  
   async goEditForm(id) { 
	  if (this.beNewWindow) {
		  if (id >0) {
			this.$router.push({ name:'k8sNodesForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'k8sNodesForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findK8sNodes({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.k8sNodes 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createK8sNodes(this.formData);
          break
        case "update": 
          res = await updateK8sNodes(this.formData);
          break
        default: 
          res = await createK8sNodes(this.formData);
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },   
    quickEdit_do(field,id,value,scope) {    
      let value2 = value;
      if (typeof(value)==="boolean")
         value2 = value?"1":"0"
      value2 =  value2+"";   
      let obj = {field:field,id:id,value:value2}	 
      this.quickEdit_do2(obj);	  
	    // if (scope._self.$refs[`popover-${scope.$index}`])
		  // scope._self.$refs[`popover-${scope.$index}`].doClose();
    },
    async quickEdit_do2(obj) {  
      const res =  await quickEdit(obj)	  
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '修改成功'
        }) 
        // this.getTableData()
      }
    },
   onExcel(){
        this.getExcelList(this.page,this.pageSize)    
    },
    onExcelAll(){
        this.getExcelList(1,1000)  
    }
  },
}
</script>
<style>
</style>
