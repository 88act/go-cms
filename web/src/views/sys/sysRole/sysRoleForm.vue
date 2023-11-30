<template>
   <div class="gocms-form-box bg-bg_color">
      <el-form ref="editForm" :model="formData" :rules="editRules" label-position="right" label-width="80px">
        <el-form-item label="角色名:" prop="roleName">
          <el-input v-model="formData.roleName" clearable placeholder="请输入" />
        </el-form-item>
     <!--   <el-form-item label="角色编码:" prop="roleCode">
          <el-input v-model="formData.roleCode" clearable placeholder="请输入" />
        </el-form-item> -->
        <el-form-item label="菜单:" prop="menuList">
          <el-tree-select v-model="formData.menuList" :data="menuListOptions" multiple
        	:render-after-expand="false" show-checkbox check-strictly check-on-click-node />
          </el-form-item>
        <el-form-item label="默认菜单:" prop="defaultMenu">
          <el-tree-select v-model="formData.defaultMenu" :data="menuListOptions"
          :render-after-expand="false" show-checkbox check-strictly check-on-click-node />
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
    createSysRole,
    deleteSysRoleByIds,
    updateSysRole,
    findSysRole,
    getSysRoleList,
    quickEdit,
    excelList
  } from '@/api/sysRole'

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
  // 声明 props 类型
  export interface FormProps {
    editId : number;
    beChange : boolean,
    index : number,
    options ?: DialogOptions;
  }
  // 声明 props 默认值
  // 推荐阅读：https://cn.vuejs.org/guide/typescript/composition-api.html#typing-component-props
  const props = withDefaults(defineProps<FormProps>(), {
    editId: () => 0,
    beChange: () => false,
    index: () => 0
  });
  // 使用 vueuse 的双向绑定工具
  const emit = defineEmits(["update:editId", "update:beChange"]);
  const id = useVModel(props, "editId", emit);
  const beChange = useVModel(props, "beChange", emit);
  // 字典
  const status_options = ref([])
  const treeOptions = ref([])
  //import FileListEdit from '@/components/mediaLib/fileListEdit.vue'
  const editForm = ref(null)
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
  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findSysRole({
      id: id.value
    })
    if (res.code === 200) {
      formData.value = res.data
     	formData.value.menuList = JSON.parse("["+formData.value.menuList+"]")
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
      formData.value.menuList = formData.value.menuList.join(",")//split(","); //JSON.stringify(formData.value.menuList)
      //图片
      let res;
      if (id.value > 0) { //update
        res = await updateSysRole(formData.value)
      } else {
        formData.value.status = 1
        res = await createSysRole(formData.value)
      }
      if (res.code === 200) {
        console.log(res)
        beChange.value = true;
        if (res.data && res.data.id) {
          id.value = res.data.id
        }
        message(res.msg, { type: "success" })
        editForm.value.resetFields()
        closeDialog(props.options, props.index)
      } else {
        message(res.msg, { type: "error" })
      }
    }
  }

  const userName_options = ref([])
  const getOptionsData = async () => {
    status_options.value = await getDict('status')
    let dictReq = {
    	table: "mem_user",
    	pidField: "id",
    	nameField: "realname"
    }
    userName_options.value = await getPidData(dictReq)
  }
  const menuListOptions = ref([])
  const getTreeData = async () => {
    let treeDataReq = {
      table: "sys_menu",
      pidField: "id",
      nameField: "title",
      pidValue: 0
    }
    menuListOptions.value = await getPidTreeData(treeDataReq)
  }
  const init = async () => {
    //console.log("props.data = ", props.data)
    console.log("id = ", id.value)
    getOptionsData()
    getTreeData()
    if (id.value > 0) {
      getData()
    }
  }
  onMounted(() => {
    init()
  })
  const formData = ref({
    id: 0,
    cuId: 0,
    pid: 0,
    roleName: '',
    roleCode: '',
    defaultMenu: 0,
    status: 0,
  })
</script>
