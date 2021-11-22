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
                <el-form-item label="名称">
                  <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
                </el-form-item> 
                <el-form-item label="对应表">
                  <el-input placeholder="搜索条件" v-model="searchInfo.toTable" clearable />
                </el-form-item> 
                <el-form-item label="当前id">
                  <el-input placeholder="搜索条件" v-model="searchInfo.nowId" clearable />
                </el-form-item>
                  <el-form-item label="当前页码">
                      <el-input placeholder="搜索条件" v-model="searchInfo.pageNow" clearable />
                  </el-form-item>
                <el-form-item label="运行状态" prop="statusRun">                
                    <el-select v-model="searchInfo.statusRun" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSearch">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="goEditForm(0)">新增</el-button>
           <el-button size="mini" type="primary" icon="el-icon-plus" @click="onExcel">导出</el-button>
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
          <el-table-column label="用户id" prop="userid" width="120"   sortable="custom"  />
          <!--name  BeQuickEdit -->  
        <el-table-column label="名称" prop="name" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.name"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('name',scope.row.ID,scope.row.name,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.name}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--url  BeQuickEdit -->  
        <el-table-column label="路径" prop="url" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.url"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('url',scope.row.ID,scope.row.url,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.url}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--urlPage BeHide -->
          <!--toTable  BeQuickEdit -->  
        <el-table-column label="对应表" prop="toTable" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.toTable"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('to_table',scope.row.ID,scope.row.toTable,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.toTable}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--nowId  BeQuickEdit -->  
        <el-table-column label="当前id" prop="nowId" width="120"   >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.nowId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('now_id',scope.row.ID,scope.row.nowId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.nowId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--pageNow  BeQuickEdit -->  
        <el-table-column label="当前页码" prop="pageNow" width="120"   >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.pageNow"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('page_now',scope.row.ID,scope.row.pageNow,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.pageNow}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--pageStart  BeQuickEdit -->  
        <el-table-column label="开始页码" prop="pageStart" width="120"   >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.pageStart"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('page_start',scope.row.ID,scope.row.pageStart,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.pageStart}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--pageEnd  BeQuickEdit -->  
        <el-table-column label="结束页码" prop="pageEnd" width="120"   >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.pageEnd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('page_end',scope.row.ID,scope.row.pageEnd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.pageEnd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--statusRun  BeQuickEdit -->
        <el-table-column label="运行状态" prop="statusRun" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.statusRun" placeholder="请选择"  @change="quickEdit_do('status_run',scope.row.ID,scope.row.statusRun,scope)">
                  <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.statusRun,"status")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--status  BeQuickEdit -->
        <el-table-column label="状态" prop="status" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">
                  <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.status,"status")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column> 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
		  <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="startCollect(scope.row.ID)">启动</el-button>
	      <el-button plain size="mini" type="danger" icon="el-icon-edit" class="table-button" @click="stopCollect(scope.row.ID)">停止</el-button>

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
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路径:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="翻页路径:"> 
              <el-input v-model="formData.urlPage" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对应表:"> 
              <el-input v-model="formData.toTable" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前id:"> 
              <el-input v-model="formData.nowId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前页码:">
                 <el-input v-model.number="formData.pageNow" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开始页码:">
                 <el-input v-model.number="formData.pageStart" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束页码:">
                 <el-input v-model.number="formData.pageEnd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运行状态:">
                 <el-select v-model="formData.statusRun" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
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
  createColCollect, 
  deleteColCollectByIds,
  updateColCollect,
  findColCollect,
  getColCollectList,
  quickEdit,
  startOrStopCollect
} from '@/api/colCollect' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'ColCollect',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:false,//是否在新窗口打开编辑器
      listApi: getColCollectList,
      statusOptions: [],
      statusOptions: [],
      formData: {
            userid: 0,
           name: '',
           url: '',
           urlPage: '',
           toTable: '',
           nowId: '',
            pageNow: 0,
            pageStart: 0,
            pageEnd: 0,
            statusRun: 0,
            status: 0,
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('status')
    await this.getDict('status')
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
        //this.deleteColCollect(row)
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
     const res = await deleteColCollectByIds({ ids })
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
			this.$router.push({ name:'colCollectForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'colCollectForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findColCollect({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.colCollect 
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
          res = await createColCollect(this.formData);
          break
        case "update": 
          res = await updateColCollect(this.formData);
          break
        default: 
          res = await createColCollect(this.formData);
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
       console.log("excel");		  
    },
	async startCollect(id) { 
		  console.log("startCollect id =",id);	 
		  const res = await startOrStopCollect({ID:id,opt:1}) 
		  console.log("startCollect",res);	 
	},
	async stopCollect(id) {
		  console.log("startCollect id =",id);	 
		  const res = await startOrStopCollect({ID:id,opt:0}) 
		  console.log("startCollect",res);	 
	},
  },
}
</script>
<style>
</style>
