<template>

    <MdEditor v-model="formData.detail"  @onUploadImg="onUploadImg" />

    <div class="btn-save">
      <el-button class="el-btn-save" type="primary" @click="save">保存</el-button>
    </div>
</template>

<script setup lang="ts">
  import { MdEditor } from 'md-editor-v3';
  import 'md-editor-v3/lib/style.css';
  import axios from 'axios';

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
  const props = withDefaults(defineProps<FormProps>(), {
    editId: () => 0,
    beChange: () => false,
    index: () => 0
  });

  const emit = defineEmits(["update:editId", "update:beChange"]);
  const id = useVModel(props, "editId", emit);
  const beChange = useVModel(props, "beChange", emit);


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

    let res;
    formData.value.id = id.value
    res = await createCmsDetail(formData.value)
    if (res.code === 200) {
      console.log(res)
      beChange.value = false;
      if (res.data && res.data.id) {
        id.value = res.data.id
      }
      message(res.msg, { type: "success" })
      closeDialog(props.options, props.index)
    } else {
      message(res.msg, { type: "error" })
    }
  }

  const myToken = ref("")

  const init = async () => {
    myToken.value = getToken().accessToken
    if (id.value > 0) {
       await getData()
    }
  }

  onMounted(() => {
    //vditor.value.setValue("hello,Vditor+Vue!");
    init()
  })

const onUploadImg = async (files, callback) => {
  const res = await Promise.all(
    files.map((file) => {
      return new Promise((rev, rej) => {
        const form = new FormData();
        form.append('file', file);
        axios.post('api/commFile/upload', form, {
            headers: {
              'Content-Type': 'multipart/form-data',
              'Authorization': myToken.value
            }
          })
          .then((res) =>{
            console.log("res=====2==",res.data);
           rev(res.data)})
          .catch((error) => rej(error));
      });
    })
  );
    callback(res.map((item) =>  item.data.path));
    //  console.log("res===1====",res)
    //   let resObj =  res[0]
    //   console.log("res===2====",resObj)
    // // let resObj = JSON.parse(resObj2)
    
    //  if (resObj && resObj.code == 200 ){
    //    let path =resObj.data.path
    //    path = path.replace(".jpg", "_src.jpg");
    //    path = path.replace(".png", "_src.png");
    //    console.log("res==path===",path)
    //    callback((path)=>path)
    //  }else {
    //     message("上传出错"+resObj.msg,{'type':"error"})
    //  }
  }

  const formData = ref({
    id: 0,
    detail: '',
  })

</script>


 // myToken.value = getToken().accessToken
    // vditor.value = new Vditor('vditor', {
    // 		width:"100%",
    // 		mode: "wysiwyg",
    // 		outline: {
    // 		    "enable": true
    // 		},
    // 		toolbarConfig: {
    // 			pin: true,
    // 		},
    // 		cache: {
    // 			enable: false,
    // 		},
    // 		counter: {
    // 		    "enable": true,
    // 		    "max": 4000
    // 		},
    // 		 upload: {
    // 		      accept: 'image/*,.mp3,.wav,.rar',
    // 		      token: myToken.value,   //'x-token': userStore.token,  X-Upload-Token
    // 		      url: 'api/commFile/upload',
    // 		      linkToImgUrl: '',
    // 			   fieldName:"file",
    // 			   success:(_, msg)=>{
    // 				  let res = JSON.parse(msg)
    // 				  if (res && res.code == 200 ){
    // 					  let path =res.data.path
    // 					  path = path.replace(".jpg", "_src.jpg");
    // 					  path = path.replace(".png", "_src.png");
    // 					  vditor.value.insertValue(`![${res.data.name}](${path})`)
    // 				  }else {
    // 					  console.log("上传出错 ,",msg)
    // 				  }
    // 				  //vditor.value.insertValue(`![${JSON.parse(msg).data.name}](${JSON.parse(msg).data.url})`)
    // 			  }
    // 		    },
    // 		after: () => {
    // 			// vditor.value is a instance of Vditor now and thus can be safely used here
    // 			 vditor.value.setValue('Vue Composition API + Vditor + TypeScript Minimal Example');
    // 			//vditor.value.setValue(formData.value.detail)
    // 		},
    // 	});
