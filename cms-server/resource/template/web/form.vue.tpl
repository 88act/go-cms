<template>
  <div>
    <div class="gocms-form-box">
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  create{{.StructName}},
  update{{.StructName}},
  find{{.StructName}}
} from '@/api/{{.PackageName}}' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑{{.StructName}}',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return { 
      {{- range .Fields}}
          {{- if .DictType }}
      {{ .DictType }}Options: [],
          {{- end }}
      {{- end }}
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
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await find{{.StructName}}({ID:id})
      if (res.code === 200) {
        this.formData = res.data.{{.Abbreviation}}
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    {{- range .Fields }}
      {{- if .DictType }}
    await this.getDict('{{.DictType}}')
      {{- end }}
    {{- end }} 
  },
  methods: {
    async save() {
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
        case 'create':
          res = await create{{.StructName}}(this.formData)
          break
        case 'update':
          res = await update{{.StructName}}(this.formData)
          break
        default:
          res = await create{{.StructName}}(this.formData)
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
         emitter.emit('closeThisPage') 
      }
    },
    back() {
      this.$router.go(-1)
      emitter.emit('closeThisPage') 
    }
  }
}
</script>
<style>
</style>
