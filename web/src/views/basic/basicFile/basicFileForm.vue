<template>	 
		<div class="gocms-form-box">
			<el-form ref="editForm" :model="formData" :rules="editRules"  label-position="right" label-width="80px" >
        
        <el-form-item label="文件名:"  prop="name">   
                <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        
   
        <el-form-item label="媒体类型:"  prop="mediaType">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_type_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="存储位置:"  prop="driver">
                 <el-input v-model.number="formData.driver" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件路径:"  prop="path">   
                <el-input v-model="formData.path" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件类型:"  prop="fileType">   
                <el-input v-model="formData.fileType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件类型:"  prop="ext">   
                <el-input v-model="formData.ext" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件大小[k]:"  prop="size">
                 <el-input v-model.number="formData.size" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="md5值:"  prop="md5">   
                <el-input v-model="formData.md5" clearable placeholder="请输入" />
       </el-form-item> 
        <el-form-item label="状态:"  prop="status">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item> 
			</el-form>
			<div class="btn-save">
				<el-button class="el-btn-save" type="primary" @click="save">保存</el-button>				
			</div>		 
	</div>
</template>
 


<script setup lang="ts">
  import { message } from "@/utils/message";
  import { useVModel } from "@vueuse/core";
  import { ref, onMounted, computed, type CSSProperties, } from "vue";
  import { storageLocal } from "@pureadmin/utils";
  import { useUserStoreHook } from "@/store/modules/user";
  import { usePermissionStoreHook } from "@/store/modules/permission";
  import { useRenderIcon } from "@/components/ReIcon/src/hooks";
  import {
    getDict,
    getPidData,
    getPidTreeData,
    getDictNew,
    getDictNew2,
    getDictTreeNew,
    getTreeName,
    getTreeFullPath,
    getOptLabel
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
    substr,
    filterDict,
    formatDate,
    formatBoolean
  } from '@/utils/utils'
 

  import {
    required,
    username,
    password,
    range,
    len,
    okChat
  } from "@/utils/formRules"
  import {
    addDialog,
    closeDialog,
    updateDialog,
    closeAllDialog
  } from "@/components/ReDialog";

  import type {
    DialogOptions
  } from "@/components/ReDialog/type";

 import {
	createBasicFile,
	updateBasicFile,
	findBasicFile
} from '@/api/basicFile'
	

  // 声明 props 类型
  export interface FormProps {
    editId : number;
    beChange : boolean,
    index : number,
    options ?: DialogOptions;
  }
  // 声明 props 默认值 
  const props = withDefaults(defineProps<FormProps>(), {
    editId: () => 0,
    beChange: () => false,
    index: () => 0
  });
  // 使用 vueuse 的双向绑定工具
  const emit = defineEmits(["update:editId", "update:beChange"]);
  const id = useVModel(props, "editId", emit);
  const beChange = useVModel(props, "beChange", emit);
  const editForm = ref(null) 
   // 字典
		const media_type_options = ref([])
		const status_options = ref([])

  const treeOptions = ref([])
  //图片处理
  import FileListEdit from '@/components/mediaLib/fileListEdit.vue'  

  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findBasicFile({
      id: id.value
    })
    if (res.code === 200) {
      formData.value = res.data
      //formData.value.pidList = getTreeFullPath(treeOptions.value, formData.value.branchId);
    } else {
       message(res.msg, { type: "error" })
    }
  }
  //保存

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
      let res;
      if (id.value > 0) { //update
        res = await updateBasicFile(formData.value)
      } else {
        formData.value.status = 1
        res = await createBasicFile(formData.value)
      }
      if (res.code === 200) {
        console.log(res)
        beChange.value = true;
        if (res.data && res.data.id) {
          id.value = res.data.id
        }
        message(res.msg, { type: "success" })
        closeDialog(props.options, props.index)
      } else {
        message(res.msg, { type: "error" })
      }
    }
  }

const getOptionsData = async () => { 
			media_type_options.value = await getDict("media_type") 
			status_options.value = await getDict("status") 
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
    //console.log("props.data = ", props.data)
    console.log("id = ", id.value)
    getOptionsData()
    //getTreeData()
    if (id.value > 0) {
      getData()
    }
  }
  onMounted(() => {
    init()
  })
 const formData = ref({
		   id:0,
           guid: '',
            userId: 0,
            userIdSys: 0,
           name: '',
            catId: 0,
            module: 0,
            mediaType: 0,
            driver: 0,
           path: '',
           fileType: '',
           ext: '',
            size: 0,
           md5: '',
           sha1: '',
            sort: 0,
            download: 0,
            usedTime: 0,
            status: 0,
	})  
 
    const editRules = ({
	})  

	  //const fileObjList = ref([])
  // const defaultProps = ref({
  // 	checkStrictly: true,
  // 	expandTrigger: "hover"
  // })

  // const el_tree_props = ref({
  // 	children: 'children',
  // 	label: 'label',
  // })

</script>



 