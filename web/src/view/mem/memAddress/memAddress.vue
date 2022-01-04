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
                  <el-form-item label="用户id">
                      <el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
                  </el-form-item> 
                <el-form-item label="收货人">
                  <el-input placeholder="搜索条件" v-model="searchInfo.realname" clearable />
                </el-form-item> 
                <el-form-item label="手机">
                  <el-input placeholder="搜索条件" v-model="searchInfo.phone" clearable />
                </el-form-item> 
                <el-form-item label="邮箱地址">
                  <el-input placeholder="搜索条件" v-model="searchInfo.email" clearable />
                </el-form-item> 
                <el-form-item label="地区编码">
                  <el-input placeholder="搜索条件" v-model="searchInfo.areaCode" clearable />
                </el-form-item> 
                <el-form-item label="详细地址">
                  <el-input placeholder="搜索条件" v-model="searchInfo.address" clearable />
                </el-form-item> 
                <el-form-item label="邮政编码">
                  <el-input placeholder="搜索条件" v-model="searchInfo.zipCode" clearable />
                </el-form-item>
              <el-form-item label="默认地址" prop="isDefault">
              <el-select v-model="searchInfo.isDefault" clearable placeholder="请选择">
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
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
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
          <el-table-column label="用户id" prop="userId" width="120"   sortable="custom"  /> 
          <el-table-column label="收货人" prop="realname" width="120"    /> 
          <el-table-column label="手机" prop="phone" width="120"   sortable="custom"  /> 
          <el-table-column label="邮箱地址" prop="email" width="120"    /> 
          <el-table-column label="地区编码" prop="areaCode" width="120"    /> 
      <!--address BeHide --> 
          <el-table-column label="邮政编码" prop="zipCode" width="120"    />
          <!--isDefault  BeQuickEdit -->
        <el-table-column label="默认地址" prop="isDefault" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.isDefault" @change="quickEdit_do('is_default',scope.row.ID,scope.row.isDefault,scope)"/></template> 
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
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="收货人:"> 
              <el-input v-model="formData.realname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:"> 
              <el-input v-model="formData.phone" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮箱地址:"> 
              <el-input v-model="formData.email" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地区编码:"> 
              <el-input v-model="formData.areaCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="详细地址:"> 
              <el-input v-model="formData.address" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮政编码:"> 
              <el-input v-model="formData.zipCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="默认地址:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isDefault" clearable ></el-switch>
              
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
  createMemAddress, 
  deleteMemAddressByIds,
  updateMemAddress,
  findMemAddress,
  getMemAddressList,
  quickEdit,
  excelList
} from '@/api/memAddress' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'MemAddress',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:false,//是否在新窗口打开编辑器
      listApi: getMemAddressList,   
      excelListApi: excelList,
      statusOptions: [],
      formData: {
            userId: 0,
           realname: '',
           phone: '',
           email: '',
           areaCode: '',
           address: '',
           zipCode: '',
           isDefault: false,
            status: 0,
            mapData: {}
      } 
    }
  },
  
  async created() {
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
        //this.deleteMemAddress(row)
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
     const res = await deleteMemAddressByIds({ ids })
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
			this.$router.push({ name:'memAddressForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'memAddressForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findMemAddress({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.memAddress 
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
          res = await createMemAddress(this.formData);
          break
        case "update": 
          res = await updateMemAddress(this.formData);
          break
        default: 
          res = await createMemAddress(this.formData);
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
