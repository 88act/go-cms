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
                <el-form-item label="文章类型" prop="mediaType">                
                    <el-select v-model="searchInfo.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item> 
                <el-form-item label="文章标题">
                  <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
                </el-form-item> 
                <el-form-item label="文章摘要">
                  <el-input placeholder="搜索条件" v-model="searchInfo.sketch" clearable />
                </el-form-item> 
                <el-form-item label="文章内容">
                  <el-input placeholder="搜索条件" v-model="searchInfo.detail" clearable />
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
          <el-table-column label="类别ID" prop="catId" width="120"   sortable="custom"  /> 
          <el-table-column label="系统类别ID" prop="catIdSys" width="120"   sortable="custom"  />
          <!--mediaType  BeQuickEdit -->
        <el-table-column label="文章类型" prop="mediaType" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.mediaType" placeholder="请选择"  @change="quickEdit_do('media_type',scope.row.ID,scope.row.mediaType,scope)">
                  <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.mediaType,"media_type")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--name  BeQuickEdit -->  
        <el-table-column label="文章标题" prop="name" width="120"   >
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
      <!--sketch BeHide --> 
      <!--detail BeHide -->
          <!--author  BeQuickEdit -->  
        <el-table-column label="文章作者" prop="author" width="120"   >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.author"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('author',scope.row.ID,scope.row.author,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.author}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--tagList BeHide --> 
      <!--authorEmail BeHide --> 
      <!--referer BeHide -->
          <el-table-column label="插图" prop="thumb" width="120"   >
              <template #default="scope">
                <ImageView :url="getMapData(scope.row.thumb,scope.row.mapData)" />
              </template>
          </el-table-column> 
      <!--qrcode BeHide --> 
      <!--imageList BeHide --> 
      <!--mediaList BeHide --> 
      <!--link BeHide -->
          <!--isTop  BeQuickEdit -->
        <el-table-column label="置顶" prop="isTop" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.isTop" @change="quickEdit_do('is_top',scope.row.ID,scope.row.isTop,scope)"/></template> 
        </el-table-column>
          <!--isHot  BeQuickEdit -->
        <el-table-column label="热门" prop="isHot" width="120"   sortable="custom"  >                        
            <template #default="scope" ><el-switch v-model="scope.row.isHot" @change="quickEdit_do('is_hot',scope.row.ID,scope.row.isHot,scope)"/></template> 
        </el-table-column>
          <!--totalDiscuss  BeQuickEdit -->  
        <el-table-column label="总评论" prop="totalDiscuss" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalDiscuss"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_discuss',scope.row.ID,scope.row.totalDiscuss,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalDiscuss}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalClick  BeQuickEdit -->  
        <el-table-column label="总点击" prop="totalClick" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalClick"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_click',scope.row.ID,scope.row.totalClick,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalClick}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--totalStar  BeQuickEdit -->  
        <el-table-column label="总评" prop="totalStar" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.totalStar"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('total_star',scope.row.ID,scope.row.totalStar,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.totalStar}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--totalStar1 BeHide --> 
      <!--totalStar2 BeHide --> 
      <!--totalStar3 BeHide --> 
      <!--totalStar4 BeHide --> 
      <!--totalStar5 BeHide -->
          <!--sort  BeQuickEdit -->  
        <el-table-column label="排序" prop="sort" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sort"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sort',scope.row.ID,scope.row.sort,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sort}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--pid  BeQuickEdit -->  
        <el-table-column label="父id" prop="pid" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.pid"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('pid',scope.row.ID,scope.row.pid,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.pid}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column> 
      <!--chapter BeHide -->
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
      <!--verifyMsg BeHide --> 
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
        <el-form-item label="类别ID:">
                 <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统类别ID:">
                 <el-input v-model.number="formData.catIdSys" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="文章标题:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章摘要:"> 
              <el-input v-model="formData.sketch" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章内容:">
              <editor ref="editor_detail" :value="formData.detail" placeholder="请输入文章内容" />
       </el-form-item>
        <el-form-item label="文章作者:"> 
              <el-input v-model="formData.author" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标签列表:"> 
              <el-input v-model="formData.tagList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="作者邮箱:"> 
              <el-input v-model="formData.authorEmail" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="来源:"> 
              <el-input v-model="formData.referer" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="插图:">
               <ImageView ref="imageView_thumb" be-edit :url="getMapData(formData.thumb,formData.mapData)" :guid="formData.thumb" />
       </el-form-item>
        <el-form-item label="二维码图片:"> 
              <el-input v-model="formData.qrcode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片列表:"> 
              <el-input v-model="formData.imageList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体列表:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="链接地址:"> 
              <el-input v-model="formData.link" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="置顶:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="热门:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isHot" clearable ></el-switch>
              
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
        <el-form-item label="总星评1:">
                 <el-input v-model.number="formData.totalStar1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评2:">
                 <el-input v-model.number="formData.totalStar2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评3:">
                 <el-input v-model.number="formData.totalStar3" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评4:">
                 <el-input v-model.number="formData.totalStar4" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评5:">
                 <el-input v-model.number="formData.totalStar5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父id:">
                 <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="章节集合:"> 
              <el-input v-model="formData.chapter" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="审核信息:"> 
              <el-input v-model="formData.verifyMsg" clearable placeholder="请输入" />
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
  createCmsArticle, 
  deleteCmsArticleByIds,
  updateCmsArticle,
  findCmsArticle,
  getCmsArticleList,
  quickEdit,
  excelList
} from '@/api/cmsArticle' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'CmsArticle',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getCmsArticleList,   
      excelListApi: excelList,
      media_typeOptions: [],
      statusOptions: [],
      formData: {
            userId: 0,
            catId: 0,
            catIdSys: 0,
            mediaType: 0,
           name: '',
           sketch: '',
           detail: '',
           author: '',
           tagList: '',
           authorEmail: '',
           referer: '',
            thumb: "",
           qrcode: '',
           imageList: '',
           mediaList: '',
           link: '',
           isTop: false,
           isHot: false,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            totalStar1: 0,
            totalStar2: 0,
            totalStar3: 0,
            totalStar4: 0,
            totalStar5: 0,
            sort: 0,
            pid: 0,
           chapter: '',
            status: 0,
           verifyMsg: '',
            mapData: {}
      } 
    }
  },
  
  async created() {
    await this.getDict('media_type')
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
        //this.deleteCmsArticle(row)
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
     const res = await deleteCmsArticleByIds({ ids })
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
			this.$router.push({ name:'cmsArticleForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'cmsArticleForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findCmsArticle({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.cmsArticle 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() { 
      this.formData.detail = this.$refs.editor_detail.getContent();
      this.formData.thumb = this.$refs.imageView_thumb.myGuid;  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createCmsArticle(this.formData);
          break
        case "update": 
          res = await updateCmsArticle(this.formData);
          break
        default: 
          res = await createCmsArticle(this.formData);
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
