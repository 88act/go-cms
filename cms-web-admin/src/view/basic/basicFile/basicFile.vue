 <!--修改 by ljd 20210725， bool datatime DictType字段 的查询填充数据 --> 

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
      <el-form-item label="唯一id">
                <el-input placeholder="搜索条件" v-model="searchInfo.guid" clearable />
              </el-form-item>
            
          <el-form-item label="用户id">
                      <el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
                  </el-form-item>
                
          <el-form-item label="文件名">
                <el-input placeholder="搜索条件" v-model="searchInfo.name" clearable />
              </el-form-item>
            
          <el-form-item label="模块名" prop="module">                
                    <el-select v-model="searchInfo.module" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
          <el-form-item label="媒体类型" prop="mediaType">                
                    <el-select v-model="searchInfo.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
          <el-form-item label="存储位置" prop="driver">                
                    <el-select v-model="searchInfo.driver" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in driverOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
                
            
            
            
            
            
            
            
            
          
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSubmit">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="openDialog">新增</el-button>
         
           <el-button size="mini" type="primary" icon="el-icon-plus" @click="excel">导出</el-button>
        
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
         <!-- add by ljd 20210709,增加id 排序功能等  -->
       <el-table-column label="ID" min-width="60" prop="ID" sortable="custom" />                
	  <el-table-column label="唯一id" prop="guid" width="120"   sortable="custom"   />  
	  <el-table-column label="用户id" prop="userId" width="120"   sortable="custom"  />  
                 <!-- BeQuickEdit -->  
                    <el-table-column label="文件名" prop="name" width="120"   sortable="custom" >
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
                 
           
        
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="模块名" prop="module" width="120"  sortable="custom" >
                    <template #default="scope">  
					 <div > {{filterDict(scope.row.module,"module")}} </div>
                   <!-- <el-popover trigger="click" placement="top"  width = "280">  
                          <el-select v-model="scope.row.module" placeholder="请选择"  @change="quickEdit_do('module',scope.row.ID,scope.row.module,scope)">
                              <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                          </el-select> 
                          <template #reference>
                              <div class="quickEdit" > {{filterDict(scope.row.module,"module")}} </div>
                          </template>
                       </el-popover> -->
                    </template>  
                    </el-table-column> 
           
        
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="媒体类型" prop="mediaType" width="120"  sortable="custom" >
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
           
        
             
                 <!-- BeQuickEdit -->
                    <el-table-column label="存储位置" prop="driver" width="120"  sortable="custom" >
                    <template #default="scope">  
					 <div   > {{filterDict(scope.row.driver,"driver")}} </div>
                    <!-- <el-popover trigger="click" placement="top"  width = "280">  
                          <el-select v-model="scope.row.driver" placeholder="请选择"  @change="quickEdit_do('driver',scope.row.ID,scope.row.driver,scope)">
                              <el-option v-for="(item,key) in driverOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                          </el-select> 
                          <template #reference>
                              <div class="quickEdit" > {{filterDict(scope.row.driver,"driver")}} </div>
                          </template>
                       </el-popover> -->
                    </template>  
                    </el-table-column> 
           
        
           <el-table-column label="配图" prop="path" width="120"   >
               <template #default="scope">
                 <ImageView  :url="scope.row.path"  />
               </template>
           </el-table-column>  
           
        
              
                  <el-table-column label="文件类型" prop="ext" width="80"   sortable="custom"  >
					  <template #default="scope">
					    <el-tag
					      :type="scope.row.ext === 'jpg' ? 'primary' : 'success'"
					      disable-transitions
					    >{{ scope.row.ext }}</el-tag>
					  </template>
                </el-table-column>
           
        
              
                  <el-table-column label="文件大小[k]" prop="size" width="120"   sortable="custom"  />  
               
           
        
           <!-- add by ljd 20210720, 隐藏字段   md5 -->         
          
        
           <!-- add by ljd 20210720, 隐藏字段   sha1 -->         
          
        
             
                 <!-- BeQuickEdit -->  
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
                 
           
        
           <!-- add by ljd 20210720, 隐藏字段   download -->         
          
        
           <!-- add by ljd 20210720, 隐藏字段   usedTime -->         
          
         

      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="updateBasicFile(scope.row)">编辑</el-button>
          <el-button plain size="mini" type="danger" icon="el-icon-delete"  @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10,30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!---------- 编辑弹窗------------------ -->
    <el-dialog :before-close="closeDialog" v-model="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="唯一id:"> 
                  <el-input v-model="formData.guid" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="用户id:"> 
                        <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="文件名:">
              
                  <el-input v-model="formData.name" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="模块名:">
              
                        <el-select v-model="formData.module" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
        <el-form-item label="媒体类型:">
              
                        <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
        <el-form-item label="存储位置:">
              
                        <el-select v-model="formData.driver" placeholder="请选择" clearable>
                          <el-option v-for="(item,key) in driverOptions" :key="key" :label="item.label" :value="item.value" />
                        </el-select>
                    </el-form-item>
        <el-form-item label="文件路径:">
              
                  <el-input v-model="formData.path" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="文件类型:">
              
                  <el-input v-model="formData.ext" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="文件大小[k]:">
              
                        <el-input v-model.number="formData.size" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="md5值:">
              
                  <el-input v-model="formData.md5" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="sha散列值:">
              
                  <el-input v-model="formData.sha1" clearable placeholder="请输入" />
              </el-form-item>
        <el-form-item label="排序:">
              
                        <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="下载次数:">
              
                        <el-input v-model.number="formData.download" clearable placeholder="请输入" />
                    </el-form-item>
        <el-form-item label="使用次数:">
              
                        <el-input v-model.number="formData.usedTime" clearable placeholder="请输入" />
                    </el-form-item>
     </el-form>
      <div slot="footer" class="el-dialog__footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createBasicFile,
  deleteBasicFile,
  deleteBasicFileByIds,
  updateBasicFile,
  findBasicFile,
  getBasicFileList,
  quickEdit
} from '@/api/basicFile' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
import ImageView from '@/components/mediaLib/imageView.vue'
import MediaLib  from '@/components/mediaLib/mediaLib.vue'
export default {
  name: 'BasicFile',
  mixins: [infoList],
  components: {
    ImageView,
   	MediaLib
  },
  data() {
    return {
      listApi: getBasicFileList,
      dialogFormVisible: false,
      type: '',
      deleteVisible: false,
      multipleSelection: [],
      
      moduleOptions: [],
          
      media_typeOptions: [],
          
      driverOptions: [],
          
      formData: {
         guid: '',
          userId: 0,
          name: '',
          module: 0,
          mediaType: 0,
          driver: 0,
          path: '',
          ext: '',
          size: 0,
          md5: '',
          sha1: '',
          sort: 0,
          download: 0,
          usedTime: 0,
          
      }, 
      shortcuts: [
                {
                  text: '最近一周',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
                    return [start, end]
                  },
                },
                {
                  text: '最近一个月',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
                    return [start, end]
                  },
                },
                {
                  text: '最近三个月',
                  value: () => {
                    const end = new Date()
                    const start = new Date()
                    start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
                    return [start, end]
                  },
           },
      ],
    }
  },
  
  async created() {
    await this.getTableData()
    
    await this.getDict('module')
      
    await this.getDict('media_type')
      
    await this.getDict('driver')
      
  },
  methods: {
  // 条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 10              
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
        this.deleteBasicFile(row)
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
      const res = await deleteBasicFileByIds({ ids })
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
    async updateBasicFile(row) {
      const res = await findBasicFile({ ID: row.ID })
      this.type = 'update'
      if (res.code === 200) {
		//console.log(" res.data =",res.data);
        this.formData = res.data.basicFile
        this.dialogFormVisible = true
      }
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.formData = {
        guid: '',
          userId: 0,
          name: '',
          module: 0,
          mediaType: 0,
          driver: 0,
          path: '',
          ext: '',
          size: 0,
          md5: '',
          sha1: '',
          sort: 0,
          download: 0,
          usedTime: 0,
          
      }
    },
    async deleteBasicFile(row) {
      const res = await deleteBasicFile({ ID: row.ID })
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === 1 && this.page > 1 ) {
          this.page--
        }
        this.getTableData()
      }
    },
    async enterDialog() {
      let res
      switch (this.type) {
        case "create":
          res = await createBasicFile(this.formData)
          break
        case "update":
          res = await updateBasicFile(this.formData)
          break
        default:
          res = await createBasicFile(this.formData)
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
    openDialog() {
      this.type = 'create'
      this.dialogFormVisible = true
    },
    //  add by ljd 20210709, 排序 
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.orderDesc = order === 'descending'
      }
      this.getTableData()
    },
    quickEdit_do(field,id,value,scope) {    
	  let value2 = value;
	  if (typeof(value)==="boolean")
		   value2 = value?"1":"0"
	  value2 =  value2+"";   
	  let obj = {field:field,id:id,value:value2}	
	 // console.log("quickEdit_do2 obj 1 =",obj);
      this.quickEdit(obj);	  
	    // if (scope._self.$refs[`popover-${scope.$index}`])
		  // scope._self.$refs[`popover-${scope.$index}`].doClose();
    },
    async quickEdit(obj) { 
    //console.log("quickEdit_do2 res 2 =",obj);
      const res =  await quickEdit(obj)	 
     // console.log("quickEdit_do2 res 3=",res);
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '修改成功'
        }) 
        // this.getTableData()
      }
    }
  },
}
</script>

<style>
</style>
