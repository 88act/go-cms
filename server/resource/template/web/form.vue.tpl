<template>
	<div>
		<div :class="beComponent==true?'':'gocms-form-box'">
			<el-form ref="editForm" :model="formData" :rules="editRules"  label-position="right" label-width="80px" > 

 {{- range .Fields }}
        <el-form-item label="{{.FieldDesc}}:"  prop="{{.FieldJson}}">
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

                {{- if eq .FieldType "json"}} 
                      <el-input type="textarea" rows="4" v-model="formData.{{.FieldJson}}" clearable placeholder="请输入" /> 
                {{- end}}  
                {{- if eq .FieldType "image"}} 
                  <div style="display: block;"> 
                      <div> 
						    <FileListEdit ref="fileListEdit_{{.FieldJson}}" :multi="true" :mediaType="1" :max="1024" :objList="getFileByGuidStr(formData.{{.FieldJson}},formData.fileObjList)" /> 
                       </div>
                       <div >
                          <span style="font-family: SourceHanSansCN-Regular; font-size: 12px; font-weight: normal; line-height: 20px;letter-spacing: -0.01px;" >只能上传jpg/png文件，且不超过1024kb</span>
                       </div>
                  </div>  
               {{- end}}  
              {{- if eq .FieldType "int"}}
                    {{- if .DictType}}
                 <el-select v-model="formData.{{ .FieldJson }}" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in {{ .DictType }}_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
                    {{- else}}
                 <el-input v-model.number="formData.{{ .FieldJson }}" clearable placeholder="请输入" />
                    {{- end}}
              {{- end}}
              {{- if eq .FieldType "time.Time"}}
                <el-date-picker v-model="formData.{{ .FieldJson }}" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
              {{- end}}
              {{- if eq .FieldType "float64"}}
               <el-input-number v-model="formData.{{ .FieldJson }}" :precision="2" clearable />
              {{- end}}
       </el-form-item>
       {{- end}} 
			</el-form>
			<div class="btn-save">
				<el-button class="el-btn-save" type="primary" @click="save">保存</el-button>
				<el-button v-if="!beComponent" class="el-btn-save" type="primary" @click="back">返回</el-button>
			</div>
		</div>
	</div>
</template>
 
<script setup>
	import {
		ref,		 
		defineEmits
	} from 'vue'
	import {
		emitter
	} from '@/utils/bus.js'
	import {
		getDict,
		getPidData,
		getPidTreeData,
		getDictNew,
		getDictNew2,
		getDictTreeNew,
		getTreeName,
		getTreeFullPath,
		getOptLabel,
	} from '@/utils/dictionary'
	import {	 
		 isEmpty,
		obj2Num,
		removeNullAttr,
		get7days,
		getShortcuts,	
		getDataTimeStr,
		getFileByGuid,
		getFileByGuidStr,
		getFileByGuidList,
	} from '@/utils/utils'
	import {
		formatDate,
		filterDict,
		getDictFunc,
		formatBoolean
	} from '@/utils/format'
	import {
		required,
		username,
		password,
		range,
		len,
		okChat
	} from "@/utils/formRules"

	import {
		toSQLLine
	} from '@/utils/stringFun'
		
	import WarningBar from '@/components/warningBar/warningBar.vue'

	import {
		ElMessage,
		ElMessageBox
	} from 'element-plus'

	import {
		create{{.StructName}},
		update{{.StructName}},
		find{{.StructName}}
	} from '@/api/{{.Abbreviation}}'
 

	const emit = defineEmits(['closeFormDialog'])
	import {
		useRoute,
		useRouter
	} from 'vue-router'
	const router = useRouter()
	const route = useRoute()
    import MyStore from "@/pinia/modules/myStore"
	const myStore = MyStore()

	const id = ref(0)

    // 字典
	{{- range .Fields}}
		{{- if .DictType }}
		const {{ .DictType }}_options = ref([])
		{{- end}}
	{{- end}}

	import FileListEdit from '@/components/mediaLib/fileListEdit.vue'

	 {{- range .Fields }}
       {{- if eq .FieldType "image"  }}
    const fileListEdit_{{.FieldJson}} = ref(null) 
      {{- end}}   
    {{- end}}  
	
	// 注意属性是只读的，不能修改
	const props = defineProps({
		beComponent: {
			type: Boolean,
			required: false,
			default: false
		},
		editId: {
			type: Number,
			required: false,
			default: 0
		}
	})

	const formData = ref({
		   id:0,
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
	})  
 
    const editRules = ({
		 {{- range .Fields}}
          {{- if .BeRequired }} 
				{{.FieldJson}}: [required],
		   {{- end }}
        {{- end}}
	})  
    //const fileObjList = ref([])
	const defaultProps = ref({
		checkStrictly: true,
		expandTrigger: "hover"
	})

	const el_tree_props = ref({
		children: 'children',
		label: 'label',
	})

	// 查询
	const getData = async () => { 
		if (id.value <= 0) {
			console.log("getData id.value <=0")
			return;
		}
		const res = await find{{.StructName}}({
			id: id.value
		})
		if (res.code === 200) {
			formData.value = res.data 
			//console.log("=========getData=========")  
			//formData.value.pidList = getTreeFullPath(treeOptions.value, formData.value.branchId);
		    //console.log(formData.value) 
		} else  ElMessage.error(res.msg)
	}
	//保存
	const editForm = ref(null) 
	const save = async () => {
		//console.log(editForm.value)
		let resValid = await editForm.value.validate((valid, fields) => {
			if (valid) {
				console.log("验证正常。。。")
				return true
			} else {
				console.log("验证失败。。。")
				return false
			}
		})
		if (resValid) { 
			delete formData.value.mapData;
			delete formData.value.createdAt;
			delete formData.value.updatedAt;
            //图片 
			{{- range .Fields }}
			{{- if eq .FieldType "image"  }}
				let guidList_{{.FieldJson}} = fileListEdit_{{.FieldJson}}.value.getGuidList()
				formData.value.{{.FieldJson}} =  guidList_{{.FieldJson}}.length >0? guidList_{{.FieldJson}}.join(","):"" 
			{{- else if .BeEditor }} 
					//BeEditor   this.formData.{{.FieldJson}} = this.$refs.editor_{{.FieldJson}}.getContent();
			{{- end}}           
			{{- end}}   

			let res;
			if (id.value > 0) { //update
				res = await update{{.StructName}}(formData.value)
			} else {
				formData.value.status = 1
				res = await create{{.StructName}}(formData.value)
			}
			if (res.code === 200) {
				ElMessage.success(res.msg)
				if (props.beComponent == true)
					emit('closeFormDialog', true)
				else
					emitter.emit('closeThisPage')
			} else {
				ElMessage.error(res.msg)
			}
			
		}
	}

	const back = () => { 
		if (props.beComponent == true)
			emit('closeFormDialog', false)
		else
			emitter.emit('closeThisPage')
	}

	const getOptionsData = async () => {
	   {{- range .Fields}}
		  {{- if .DictType }} 
			 {{.DictType}}_options.value = await getDict("{{.DictType}}")
		   {{- end}}
		{{- end}} 
	}
	const getTreeData = async () => {
		let treeDataReq = {
			table: "memBranch",
			pidField: "id",
			nameField: "title",
			pidValue: 0
		}
		treeOptions.value = await getPidTreeData(treeDataReq)
	}


	const init = async () => {
		//console.log("props.beComponent = ", props.beComponent)
		//console.log("props.editId = ", props.editId)
		if (props.editId && props.editId > 0)
			id.value = props.editId
		else if (route && route.params && route.params.id > 0) {
			id.value = Number(route.params.id)
		}
		getOptionsData()
		//getTreeData()
		if (id.value > 0) {
			getData()
		}
	}
	init()
</script>
<style>
</style>
