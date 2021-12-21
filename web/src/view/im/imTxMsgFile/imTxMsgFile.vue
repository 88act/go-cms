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
                <el-form-item label="消息类型">
                  <el-input placeholder="搜索条件" v-model="searchInfo.chatType" clearable />
                </el-form-item> 
                <el-form-item label="消息时间">
                  <el-input placeholder="搜索条件" v-model="searchInfo.msgTime" clearable />
                </el-form-item>
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_downloadOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
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
          <el-table-column label="消息类型" prop="chatType" width="120"   sortable="custom"  /> 
          <el-table-column label="消息时间" prop="msgTime" width="120"   sortable="custom"  /> 
      <!--url BeHide --> 
          <el-table-column label="过期时间" prop="expireTime" width="120"   sortable="custom"  /> 
          <el-table-column label="文件大小" prop="fileSize" width="120"   sortable="custom"  /> 
          <el-table-column label="文件md5" prop="fileMd5" width="120"   sortable="custom"  /> 
          <el-table-column label="压缩大小" prop="gzipSize" width="120"   sortable="custom"  /> 
          <el-table-column label="压缩md5" prop="gzipMd5" width="120"   sortable="custom"  /> 
      <!--localFile BeHide -->
        <el-table-column label="状态" prop="status" width="120"  sortable="custom" >
          <template #default="scope">
            {{filterDict(scope.row.status,"status_download")}}
          </template>
        </el-table-column> 
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
        <el-form-item label="消息类型:"> 
              <el-input v-model="formData.chatType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消息时间:"> 
              <el-input v-model="formData.msgTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下载路径:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="过期时间:">
                <el-date-picker v-model="formData.expireTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="文件大小:">
                 <el-input v-model.number="formData.fileSize" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件md5:">
                 <el-input v-model.number="formData.fileMd5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="压缩大小:">
                 <el-input v-model.number="formData.gzipSize" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="压缩md5:">
                 <el-input v-model.number="formData.gzipMd5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="本地路径:"> 
              <el-input v-model="formData.localFile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_downloadOptions" :key="key" :label="item.label" :value="item.value" />
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
  createImTxMsgFile, 
  deleteImTxMsgFileByIds,
  updateImTxMsgFile,
  findImTxMsgFile,
  getImTxMsgFileList,
  quickEdit,
  excelList
} from '@/api/imTxMsgFile' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'ImTxMsgFile',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:false,//是否在新窗口打开编辑器
      listApi: getImTxMsgFileList,   
      excelListApi: excelList,
      status_downloadOptions: [],
      formData: {
           chatType: '',
           msgTime: '',
           url: '',
            expireTime: new Date(),
            fileSize: 0,
            fileMd5: 0,
            gzipSize: 0,
            gzipMd5: 0,
           localFile: '',
            status: 0,
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('status_download')
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
        //this.deleteImTxMsgFile(row)
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
     const res = await deleteImTxMsgFileByIds({ ids })
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
			this.$router.push({ name:'imTxMsgFileForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'imTxMsgFileForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findImTxMsgFile({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.imTxMsgFile 
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
          res = await createImTxMsgFile(this.formData);
          break
        case "update": 
          res = await updateImTxMsgFile(this.formData);
          break
        default: 
          res = await createImTxMsgFile(this.formData);
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
