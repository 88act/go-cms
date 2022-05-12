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
                <el-form-item label="标题">
                  <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
                </el-form-item> 
                <el-form-item label="简介">
                  <el-input placeholder="搜索条件" v-model="searchInfo.desc" clearable />
                </el-form-item> 
                <el-form-item label="手机">
                  <el-input placeholder="搜索条件" v-model="searchInfo.mobile" clearable />
                </el-form-item>
                  <el-form-item label="vip等级">
                      <el-input placeholder="搜索条件" v-model="searchInfo.vipLev" clearable />
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
          <!--userId  BeQuickEdit -->  
        <el-table-column label="用户id" prop="userId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.userId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('user_id',scope.row.ID,scope.row.userId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.userId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--name  BeQuickEdit -->  
        <el-table-column label="标题" prop="name" width="120"   sortable="custom" >
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
      <!--desc BeHide --> 
      <!--detail BeHide --> 
      <!--avater BeHide --> 
      <!--mediaList BeHide -->
          <!--address  BeQuickEdit -->  
        <el-table-column label="地址" prop="address" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.address"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('address',scope.row.ID,scope.row.address,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.address}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--areaId  BeQuickEdit -->  
        <el-table-column label="地区id" prop="areaId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.areaId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('area_id',scope.row.ID,scope.row.areaId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.areaId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--lng BeHide --> 
      <!--lat BeHide -->
          <!--email  BeQuickEdit -->  
        <el-table-column label="邮件" prop="email" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.email"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('email',scope.row.ID,scope.row.email,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.email}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--mobile  BeQuickEdit -->  
        <el-table-column label="手机" prop="mobile" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.mobile"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('mobile',scope.row.ID,scope.row.mobile,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.mobile}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--vipLev BeHide --> 
      <!--vipTime BeHide --> 
      <!--totalWhole BeHide --> 
      <!--totalShare BeHide --> 
      <!--totalFav BeHide --> 
      <!--totalJoin BeHide --> 
      <!--totalDiscuss BeHide --> 
      <!--totalClick BeHide --> 
      <!--totalStar BeHide -->
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
        <el-form-item label="标题:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="简介:">
              <editor ref="editor_desc" :value="formData.desc" placeholder="请输入简介" />
       </el-form-item>
        <el-form-item label="详细内容:">
              <editor ref="editor_detail" :value="formData.detail" placeholder="请输入详细内容" />
       </el-form-item>
        <el-form-item label="缩略图:"> 
              <el-input v-model="formData.avater" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体列表:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.address" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地区id:">
                 <el-input v-model.number="formData.areaId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经度:">
                 <el-input v-model.number="formData.lng" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="纬度:">
                 <el-input v-model.number="formData.lat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮件:"> 
              <el-input v-model="formData.email" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:"> 
              <el-input v-model="formData.mobile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip等级:">
                 <el-input v-model.number="formData.vipLev" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip结束时间:">
                <el-date-picker v-model="formData.vipTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="综合指数:">
                 <el-input v-model.number="formData.totalWhole" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总分享:">
                 <el-input v-model.number="formData.totalShare" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总收藏:">
                 <el-input v-model.number="formData.totalFav" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总报名人数:">
                 <el-input v-model.number="formData.totalJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总评论:">
                 <el-input v-model.number="formData.totalDiscuss" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总点击:">
                 <el-input v-model.number="formData.totalClick" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总评:">
                 <el-input v-model.number="formData.totalStar" clearable placeholder="请输入" />
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
  createActShop, 
  deleteActShopByIds,
  updateActShop,
  findActShop,
  getActShopList,
  quickEdit,
  excelList
} from '@/api/actShop' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'ActShop',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getActShopList,   
      excelListApi: excelList,
      statusOptions: [],
      formData: {
            userId: 0,
           name: '',
           desc: '',
           detail: '',
           avater: '',
           mediaList: '',
           address: '',
            areaId: 0,
            lng: 0,
            lat: 0,
           email: '',
           mobile: '',
            vipLev: 0,
            vipTime: new Date(),
            totalWhole: 0,
            totalShare: 0,
            totalFav: 0,
            totalJoin: 0,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
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
        //this.deleteActShop(row)
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
     const res = await deleteActShopByIds({ ids })
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
			this.$router.push({ name:'actShopForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'actShopForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findActShop({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.actShop 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() { 
      this.formData.desc = this.$refs.editor_desc.getContent(); 
      this.formData.detail = this.$refs.editor_detail.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createActShop(this.formData);
          break
        case "update": 
          res = await updateActShop(this.formData);
          break
        default: 
          res = await createActShop(this.formData);
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
