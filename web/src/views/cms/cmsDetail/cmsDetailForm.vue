<template>
		<div class="gocms-form-box bg-bg_color">
			<el-form ref="editForm" :model="formData" :rules="editRules"  label-position="right" label-width="80px" >
        <el-form-item label="文章id:"  prop="artId">
                 <el-input v-model.number="formData.artId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="详细:"  prop="detail"> 			<div id="vditor" />

       </el-form-item>
			</el-form>
			<div class="btn-save">
				<el-button class="el-btn-save" type="primary" @click="save">保存</el-button>
			</div>
	</div>
</template>

<script setup lang="ts">
  import { getToken, formatToken } from "@/utils/auth";
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
	createCmsDetail,
	updateCmsDetail,
	findCmsDetail
} from '@/api/cmsDetail'


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

   import Vditor from 'vditor';
   import 'vditor/dist/index.css';
   const vditor = ref(null);


  const treeOptions = ref([])
  //图片处理
  import FileListEdit from '@/components/mediaLib/fileListEdit.vue'

  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findCmsDetail({
      id: id.value
    })
    if (res.code === 200) {
      formData.value = res.data
      //formData.value.pidList = getTreeFullPath(treeOptions.value, formData.value.branchId);
    } else {
     //  message(res.msg, { type: "error" })
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
			//BeEditor   this.formData.detail = this.$refs.editor_detail.getContent();
      let res;
      if (id.value > 0) { //update
        res = await updateCmsDetail(formData.value)
      } else {
        formData.value.status = 1
        res = await createCmsDetail(formData.value)
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

  const myToken =ref("")

  const init = async () => {
    //console.log("props.data = ", props.data)
    console.log("id = ", id.value)
    getOptionsData()
     getTreeData()
    if (id.value > 0) {
      await  getData()
    }
    myToken.value = getToken().accessToken
    vditor.value = new Vditor('vditor', {
    		width:"100%",
    		mode: "wysiwyg",
    		outline: {
    		    "enable": true
    		},
    		toolbarConfig: {
    			pin: true,
    		},
    		cache: {
    			enable: false,
    		},
    		counter: {
    		    "enable": true,
    		    "max": 4000
    		},
    		 upload: {
    		      accept: 'image/*,.mp3,.wav,.rar',
    		      token: myToken.value,   //'x-token': userStore.token,  X-Upload-Token
    		      url: 'api/commFile/upload',
    		      linkToImgUrl: '',
    			   fieldName:"file",
    			   success:(_, msg)=>{
    				  let res = JSON.parse(msg)
    				  if (res && res.code == 200 ){
    					  let path =res.data.path
    					  path = path.replace(".jpg", "_src.jpg");
    					  path = path.replace(".png", "_src.png");
    					  vditor.value.insertValue(`![${res.data.name}](${path})`)
    				  }else {
    					  console.log("上传出错 ,",msg)
    				  }
    				  //vditor.value.insertValue(`![${JSON.parse(msg).data.name}](${JSON.parse(msg).data.url})`)
    			  }
    		    },
    		after: () => {
    			// vditor.value is a instance of Vditor now and thus can be safely used here
    			 // vditor.value.setValue('Vue Composition API + Vditor + TypeScript Minimal Example');
    			vditor.value.setValue(formData.value.detail)
    		},
    	});
  }

  onMounted(() => {
    init()
  })



 const formData = ref({
		   id:0,
         artId: 0,
           detail: '',
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

