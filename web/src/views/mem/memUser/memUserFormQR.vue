<template>
  <div class="gocms-form-box bg-bg_color"  style="display: flex;  justify-content: center;" >
     <div style="display: flex; flex-wrap: wrap; justify-content: center; width:350px">
        <ReQrcode :text="qrcodeStr" :width="300" />
         <div style="vertical-align: bottom;">
           <span>二维码有效期: 5分钟 </span>
            <el-button  class="el-btn-save" type="primary" @click="refQrcode()" > 刷 新</el-button>
        </div>
      </div>
  </div>
</template>

<script setup lang="ts">
  import { message } from "@/utils/message";
  import { useVModel } from "@vueuse/core";
  import { ref, onMounted, computed, type CSSProperties, } from "vue";
  import { isString,cloneDeep,isAllEmpty,storageLocal } from "@pureadmin/utils";
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
    mobile,
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
  showQrcode
 } from '@/api/base'
 import ReQrcode from "@/components/ReQrcode";

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

   const qrcodeStr =ref("")
     async function refQrcode() {
      let data = {
       "id":id.value
      }
      let res = await showQrcode(data)
      if (res.code == 200) {
        // qrcodeShow.value = true
         qrcodeStr.value = res.data
        //  ElMessage.success(res.msg)
      } else  message(res.msg,{type:"error"})
   }

  const init = async () => {
   refQrcode()

  }
  onMounted(() => {
    init()
  })

</script>
