<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      {{- if .SearchCreate}} 
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
      {{- end }} 
      {{- if .SearchId}} 
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item>
      {{- end -}} 

      {{- range .Fields}} 
      {{- if .FieldSearchType}}   
          {{- if eq .FieldType "bool" }}
              <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">
              <el-select v-model="searchInfo.{{.FieldJson}}" clearable placeholder="请选择">
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
          {{- else if eq .FieldType "int" }}
              {{- if .DictType }}
                <el-form-item label="{{.FieldDesc}}" prop="{{.FieldJson}}">                
                    <el-select v-model="searchInfo.{{ .FieldJson }}" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
              {{- else }}
                  <el-form-item label="{{.FieldDesc}}">
                      <el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}" clearable />
                  </el-form-item>
              {{- end }}      
            {{- else if eq .FieldType "time.Time" }}
              <el-form-item label="{{.FieldDesc}}"> 
                <el-date-picker
                v-model="formData.{{ .FieldJson }}"  
                type="datetimerange"
                format="YYYY-MM-DD HH:mm:ss"
                :shortcuts="shortcuts"
                range-separator="至"
                start-placeholder="开始日期"
                end-placeholder="结束日期"
              /> 
             </el-form-item>
            {{- else }} 
                <el-form-item label="{{.FieldDesc}}">
                  <el-input placeholder="搜索条件" v-model="searchInfo.{{.FieldJson}}" clearable />
                </el-form-item>
            {{- end }}
          {{- end }}  
        {{- end }}
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSearch">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="goEditForm(0)">新增</el-button>
         {{- if .BeExcel }}
           <el-button size="mini" type="primary" icon="el-icon-plus" @click="onExcel">导出</el-button>
        {{- end }}
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
{{- range .Fields}} 
  {{- if .BeHide }} 
      <!--{{.FieldJson}} BeHide -->            
  {{- else}}  
      {{- if .BeQuickEdit}}
          <!--{{.FieldJson}}  BeQuickEdit --> 
        {{- if .DictType}}
        <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" {{if .OrderBy}} sortable="custom"{{end}} >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.{{.FieldJson}}" placeholder="请选择"  @change="quickEdit_do('{{.ColumnName}}',scope.row.ID,scope.row.{{.FieldJson}},scope)">
                  <el-option v-for="(item,key) in {{.DictType}}Options" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{"{{"}}filterDict(scope.row.{{.FieldJson}},"{{.DictType}}"){{"}}"}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}}  >                        
            <template #default="scope" ><el-switch v-model="scope.row.{{.FieldJson}}" @change="quickEdit_do('{{.ColumnName}}',scope.row.ID,scope.row.{{.FieldJson}},scope)"/></template> 
        </el-table-column> 
        {{- else }}  
        <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}} >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.{{.FieldJson}}"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('{{.ColumnName}}',scope.row.ID,scope.row.{{.FieldJson}},scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{"{{"}}scope.row.{{.FieldJson}}{{"}}"}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>              
        {{- end -}}
      {{- else}}  
          {{- if .DictType}}
        <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120" {{if .OrderBy}} sortable="custom"{{end}} >
          <template #default="scope">
            {{"{{"}}filterDict(scope.row.{{.FieldJson}},"{{.DictType}}"){{"}}"}}
          </template>
        </el-table-column>
        {{- else if eq .FieldType "bool" }}
        <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}}  >
          <template #default="scope">{{"{{formatBoolean(scope.row."}}{{.FieldJson}}{{")}}"}}</template>
        </el-table-column>
          {{- else if eq .FieldType "image" }}
          <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}} >
              <template #default="scope">
                <ImageView :url="getMapData(scope.row.{{.FieldJson}},scope.row.mapData)" />
              </template>
          </el-table-column>
        {{- else }} 
          <el-table-column label="{{.FieldDesc}}" prop="{{.FieldJson}}" width="120"  {{if .OrderBy}} sortable="custom"{{end}}  />
        {{- end }}  
      {{- end}} 
  {{- end}} 
{{- end}} 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ "{{ formatDate(scope.row.CreatedAt)}}" }}</template>
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
    {{- range .Fields }}
        <el-form-item label="{{.FieldDesc}}:">
              {{- if eq .FieldType "bool"}}
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.{{.FieldJson}}" clearable ></el-switch>
              {{end -}}
              {{- if eq .FieldType "string"}}
                    {{- if .BeEditor }}
              <editor ref="editor_{{.FieldJson}}" :value="formData.{{.FieldJson}}" placeholder="请输入{{.FieldDesc}}" />  
                    {{- else}} 
              <el-input v-model="formData.{{.FieldJson}}" clearable placeholder="请输入" />
                   {{- end}}  
              {{- end}} 
              {{- if eq .FieldType "image"}}
               <ImageView ref="imageView_{{.FieldJson}}" be-edit :url="getMapData(formData.{{.FieldJson}},formData.mapData)" :guid="formData.{{.FieldJson}}" /> 
               {{- end}}  
              {{- if eq .FieldType "int"}}
                    {{- if .DictType}}
                 <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in {{ .DictType }}Options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
                    {{- else}}
                 <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="请输入" />
                    {{- end}}
              {{- end}}
              {{- if eq .FieldType "time.Time"}}
               <el-date-picker type="datetimerange" v-model="formData.{{ .FieldJson }}" format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss" :style="{width: '100%'}" start-placeholder="开始日期"
                  end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
              {{- end}}
              {{- if eq .FieldType "float64"}}
               <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable />
              {{- end}}
       </el-form-item>
       {{- end}}
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
  create{{.StructName}}, 
  delete{{.StructName}}ByIds,
  update{{.StructName}},
  find{{.StructName}},
  get{{.StructName}}List,
  quickEdit
} from '@/api/{{.PackageName}}' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: '{{.StructName}}',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:{{.BeNewWindow}},//是否在新窗口打开编辑器
      listApi: get{{ .StructName }}List,   
 {{- range .Fields}}
      {{- if .DictType }}
      {{ .DictType }}Options: [],
     {{- end}}
  {{- end}}
      formData: {
       {{- range .Fields}}
          {{- if eq .FieldType "bool" }}
           {{.FieldJson}}: false,
          {{- end }}
          {{- if eq .FieldType "string" }}
           {{.FieldJson}}: '',
          {{- end }}
          {{- if eq .FieldType "int" }}
            {{.FieldJson}}: 0,
          {{- end }}
          {{- if eq .FieldType "time.Time" }}
            {{.FieldJson}}: new Date(),
         {{- end}}
         {{- if eq .FieldType "float64" }}
            {{.FieldJson}}: 0,
          {{- end}}
           {{- if eq .FieldType "image" }}
            {{.FieldJson}}: "",
          {{- end }}
        {{- end}}
            mapData: {}
      } 
    }
  },
  
  async created() { 
    {{- range .Fields}}
      {{- if .DictType }}
    await this.getDict('{{.DictType}}')
      {{- end}}
    {{- end}}
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
        //this.delete{{.StructName}}(row)
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
     const res = await delete{{.StructName}}ByIds({ ids })
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
			this.$router.push({ name:'{{.Abbreviation}}Form', params:{id:id}})
		  } else {
			 this.$router.push({ name:'{{.Abbreviation}}Form',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await find{{.StructName}}({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.{{.Abbreviation}} 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() {   
      {{- range .Fields }}
       {{- if eq .FieldType "image"  }}
      this.formData.{{.FieldJson}} = this.$refs.imageView_{{.FieldJson}}.myGuid;           
       {{- else if .BeEditor }} 
      this.formData.{{.FieldJson}} = this.$refs.editor_{{.FieldJson}}.getContent();
       {{- end}}           
      {{- end}}  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await create{{.StructName}}(this.formData);
          break
        case "update": 
          res = await update{{.StructName}}(this.formData);
          break
        default: 
          res = await create{{.StructName}}(this.formData);
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
    }
  },
}
</script>
<style>
</style>
