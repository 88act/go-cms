<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
      <el-form-item label="创建时间">
            <el-date-picker 
                  v-model="searchInfo.createdAtBetween" 
                  type="datetimerange"
                  format="YYYY-MM-DD HH:mm:ss"
                  :shortcuts="shortcuts"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item> 
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item> 
                <el-form-item label="表名">
                  <el-input placeholder="搜索条件" v-model="searchInfo.tabName" clearable />
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
          <!--tabName  BeQuickEdit -->  
        <el-table-column label="表名" prop="tabName" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tabName"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('tab_name',scope.row.ID,scope.row.tabName,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tabName}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--requestMeta BeHide --> 
      <!--superBuilderPath BeHide --> 
      <!--injectionMeta BeHide -->
          <!--structName  BeQuickEdit -->  
        <el-table-column label="struct名称" prop="structName" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.structName"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('struct_name',scope.row.ID,scope.row.structName,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.structName}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--structCnName  BeQuickEdit -->  
        <el-table-column label="struct中文名" prop="structCnName" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.structCnName"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('struct_cn_name',scope.row.ID,scope.row.structCnName,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.structCnName}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--apiIds  BeQuickEdit -->  
        <el-table-column label="api" prop="apiIds" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.apiIds"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('api_ids',scope.row.ID,scope.row.apiIds,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.apiIds}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--flag  BeQuickEdit -->  
        <el-table-column label="标记" prop="flag" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.flag"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('flag',scope.row.ID,scope.row.flag,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.flag}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tableName  BeQuickEdit -->  
        <el-table-column label="表名" prop="tableName" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tableName"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('table_name',scope.row.ID,scope.row.tableName,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tableName}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini"  type="success"  @click="goSuperBuilder(scope.row.ID)">从新生成</el-button>
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
        <el-form-item label="表名:"> 
              <el-input v-model="formData.tabName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表单:">
              <editor ref="editor_requestMeta" :value="formData.requestMeta" placeholder="请输入表单" />
       </el-form-item>
        <el-form-item label="路径:">
              <editor ref="editor_superBuilderPath" :value="formData.superBuilderPath" placeholder="请输入路径" />
       </el-form-item>
        <el-form-item label="注射:">
              <editor ref="editor_injectionMeta" :value="formData.injectionMeta" placeholder="请输入注射" />
       </el-form-item>
        <el-form-item label="struct名称:"> 
              <el-input v-model="formData.structName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="struct中文名:"> 
              <el-input v-model="formData.structCnName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="api:"> 
              <el-input v-model="formData.apiIds" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标记:">
                 <el-input v-model.number="formData.flag" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表名:"> 
              <el-input v-model="formData.tableName" clearable placeholder="请输入" />
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
  createSysSuperBuilderHistories, 
  deleteSysSuperBuilderHistoriesByIds,
  updateSysSuperBuilderHistories,
  findSysSuperBuilderHistories,
  getSysSuperBuilderHistoriesList,
  quickEdit,
  excelList
} from '@/api/sysSuperBuilderHistories' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'SysSuperBuilderHistories',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getSysSuperBuilderHistoriesList,   
      excelListApi: excelList,
      formData: {
           tabName: '',
           requestMeta: '',
           superBuilderPath: '',
           injectionMeta: '',
           structName: '',
           structCnName: '',
           apiIds: '',
            flag: 0,
           tableName: '',
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
        //this.deleteSysSuperBuilderHistories(row)
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
     const res = await deleteSysSuperBuilderHistoriesByIds({ ids })
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
			this.$router.push({ name:'sysSuperBuilderHistoriesForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'sysSuperBuilderHistoriesForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findSysSuperBuilderHistories({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.sysSuperBuilderHistories 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() { 
      this.formData.requestMeta = this.$refs.editor_requestMeta.getContent(); 
      this.formData.superBuilderPath = this.$refs.editor_superBuilderPath.getContent(); 
      this.formData.injectionMeta = this.$refs.editor_injectionMeta.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createSysSuperBuilderHistories(this.formData);
          break
        case "update": 
          res = await updateSysSuperBuilderHistories(this.formData);
          break
        default: 
          res = await createSysSuperBuilderHistories(this.formData);
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
    },
    goSuperBuilder(ID) {
      if (ID>0) {
        this.$router.push({ name: 'superBuilderEdit', params: {
          id:ID
        }})
      } else {
        this.$router.push({ name: 'superBuilder' })
      }
    }
  },
}
</script>
<style>
</style>
