<template>
  <div class="gocms-form-box bg-bg_color">
    <el-form ref="editForm" :model="formData" :rules="editRules" label-position="right" label-width="80px">
      <el-form-item label="文章:" prop="artId">
        <el-select v-model="formData.artId" placeholder="请选择"  :filterable="true" clearable style="width: 450px;" >
          <el-option v-for="(item,key) in art_options" :key="key" :label="item.label" :value="item.value" />
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
    createCmsCatArt,
    updateCmsCatArt,
    findCmsCatArt
  } from '@/api/cmsCatArt'


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
  const art_options = ref([])

  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findCmsCatArt({
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
       formData.value.catId= id.value
       formData.value.status = 1
       let  res = await createCmsCatArt(formData.value)
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

    let dictReq = {
      table: "cms_art",
      pidField: "id",
      nameField: "title"
    }
    art_options.value = await getPidData(dictReq)
  }
  const init = async () => {
   //console.log("props.data = ", props.data)
    console.log("id = ", id.value)
    getOptionsData()
    // if (id.value > 0) {
    //   getData()
    // }
  }
  onMounted(() => {
    init()
  })
  const formData = ref({
    id: 0,
    userId: 0,
    catId: 0,
    artId: 0,
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
