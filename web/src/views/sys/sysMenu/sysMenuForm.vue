<template>
  <div class="gocms-form-box bg-bg_color">
    <el-form ref="editForm" :model="formData" :rules="editRules" label-position="right" label-width="80px">
      <!--  <el-form-item label="商户:"  prop="cuId">
                 <el-input v-model.number="formData.cuId" clearable placeholder="请输入" />
       </el-form-item>
       <el-form-item label="父菜单ID:"  prop="pid">
                 <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
       </el-form-item> -->
      <el-form-item label="菜单标题:" prop="title">
        <el-input v-model="formData.title" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="上级菜单:" prop="pidList">
        <el-cascader v-model="formData.pidList" :options="treeOptions" :props="defaultProps" clearable  style="width:400px" />
      </el-form-item>
      <el-form-item label="name:" prop="name">
        <el-input v-model="formData.name" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="path:" prop="path">
        <el-input v-model="formData.path" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="前端路径:" prop="component">
        <el-input v-model="formData.component" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="是否隐藏:" prop="showLink">
        <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
          v-model="formData.showLink" clearable></el-switch>
      </el-form-item>
      <el-form-item label="参数:" prop="param">
        <el-input v-model="formData.param" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="排序:" prop="sort">
        <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
      </el-form-item>
    <!--  <el-form-item label="保持连接:" prop="keepAlive">
        <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否"
          v-model="formData.keepAlive" clearable></el-switch>
      </el-form-item> -->
      <el-form-item label="图标:" prop="icon">
        <el-input v-model="formData.icon" clearable placeholder="请输入" />
      </el-form-item>
      <!--   <el-form-item label="活动链接:"  prop="activePath">
                <el-input v-model="formData.activePath" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="外链:"  prop="frameSrc">
                <el-input v-model="formData.frameSrc" clearable placeholder="请输入" />
       </el-form-item>
      <el-form-item label="等级:"  prop="level">
                 <el-input v-model.number="formData.level" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="按钮权限:"  prop="auths">
                <el-input v-model="formData.auths" clearable placeholder="请输入" />
       </el-form-item>    <el-form-item label="状态:"  prop="status">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item> -->
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
    createSysMenu,
    deleteSysMenuByIds,
    updateSysMenu,
    findSysMenu,
    getSysMenuList,
    quickEdit,
    excelList
  } from '@/api/sysMenu'

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
    index: number,
    options?: DialogOptions;
  }
  // 声明 props 默认值
  // 推荐阅读：https://cn.vuejs.org/guide/typescript/composition-api.html#typing-component-props
  const props = withDefaults(defineProps<FormProps>(), {
    editId: () => 0,
    beChange: () => false,
    index: ()=>0
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
    const res = await findSysMenu({
      id: id.value
    })
    if (res.code === 200) {
      formData.value = res.data
      formData.value.pidList = getTreeFullPath(treeOptions.value, formData.value.pid);
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

      if (formData.value.pidList && formData.value.pidList.length > 0)
        formData.value.pid = formData.value.pidList[formData.value.pidList.length - 1]
      else
        formData.value.pid = 0
      if (formData.value.id >0 &&  formData.value.id == formData.value.pid){
         message("不能选自己作为上级菜单",{"type":"error"})
         return
      }
      let res;
      if (id.value > 0) { //update
        res = await updateSysMenu(formData.value)
      } else {
        formData.value.status = 1
        res = await createSysMenu(formData.value)
      }
      if (res.code === 200) {
        console.log(res)
        beChange.value = true;
        if (res.data && res.data.id){
           id.value = res.data.id
        }
        message(res.msg, { type:"success"})
        closeDialog(props.options,props.index)
      } else {
        message(res.msg, { type: "error" })
      }
    }
  }

  const getOptionsData = async () => {
    status_options.value = await getDict("status")
  }
  const getTreeData = async () => {
   let treeDataReq = {
     table: "sys_menu",
     pidField: "id",
     nameField: "title",
     pidValue: 0,
     where:"be_sys=1"
   }
   treeOptions.value = await getPidTreeData(treeDataReq)
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

  const defaultProps = ref({
  	checkStrictly: true,
  	expandTrigger: "hover"
  })

  const formData = ref({
    id: 0,
    cuId: 0,
    pid: 0,
    name: '',
    path: '',
    title: '',
    level: 0,
    auths: '',
    showLink: false,
    activePath: '',
    frameSrc: '',
    component: '',
    param: '',
    sort: 0,
    keepAlive: false,
    icon: '',
    status: 0,
    pidList:[]
  })
</script>
