<template>
	<div>
			<el-form ref="editForm" :model="formData" :rules="editRules"  label-position="right" label-width="80px" >
        <el-form-item label="客户id:"  prop="cuId">
                 <el-input v-model.number="formData.cuId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:"  prop="userId">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="log类型:"  prop="logType">
             <el-input v-model.number="formData.logType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求方法:"  prop="method">
                <el-input v-model="formData.method" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求路径:"  prop="path">
                <el-input v-model="formData.path" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="延迟:"  prop="latency">
                 <el-input v-model.number="formData.latency" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="代理:"  prop="agent">
                <el-input v-model="formData.agent" type="textarea" rows="1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="错误信息:"  prop="errorMessage">
                <el-input v-model="formData.errorMessage" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求Body:"  prop="body">
                <el-input v-model="formData.body" type="textarea" rows="5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="响应Body:"  prop="resp">
                <el-input v-model="formData.resp" type="textarea" rows="5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求ip:"  prop="ip">
                <el-input v-model="formData.ip" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"  prop="ipAddr">
                <el-input v-model="formData.ipAddr" clearable placeholder="请输入" />
       </el-form-item>
       <!-- <el-form-item label="请求状态:"  prop="status">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item> -->
			</el-form>
		<!-- 	<div class="btn-save">
				<el-button class="el-btn-save" type="primary" @click="save">保存</el-button>
			</div> -->
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
    createSysLogs,
    deleteSysLogsByIds,
    updateSysLogs,
    findSysLogs,
    getSysLogsList,
    quickEdit,
    excelList
  } from '@/api/SysLogs'

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
  //const id = ref(0)

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
    const res = await findSysLogs({
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
        res = await updateSysLogs(formData.value)
      } else {
        formData.value.status = 1
        res = await createSysLogs(formData.value)
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
//getOptionsData()
    //getTreeData()
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
  })
</script>

