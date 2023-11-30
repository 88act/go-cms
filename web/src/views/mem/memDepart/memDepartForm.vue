<template>
  <div class="gocms-form-box">
    <el-form ref="editForm" :model="formData" :rules="editRules" label-position="right" label-width="80px">
      <el-form-item label="部门名称:" prop="name">
        <el-input v-model="formData.name" clearable placeholder="请输入" />
      </el-form-item>
     <el-form-item label="上级部门:">
				<el-cascader v-model="formData.pidList" :options="treeOptions" :props="defaultProps" clearable
					style="width:400px" />
			</el-form-item>
      <el-form-item label="排序:" prop="sort">
        <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="联系电话:" prop="phone">
        <el-input v-model="formData.phone" clearable placeholder="请输入" />
      </el-form-item>
    
      <el-form-item label="描述:" prop="desc">
        <el-input v-model="formData.desc" type="textarea" rows="3" clearable placeholder="请输入" />
      </el-form-item>
      <!-- <el-form-item label="地址:" prop="address">
        <el-input v-model="formData.address" clearable placeholder="请输入" />
      </el-form-item>
      <el-form-item label="编号:" prop="encode">
        <el-input v-model="formData.encode" clearable placeholder="请输入" />
      </el-form-item>
  <el-form-item label="邮箱:" prop="email">
        <el-input v-model="formData.email" clearable placeholder="请输入" />
      </el-form-item> 
      <el-form-item label="配图:" prop="image">
        <div style="display: block;">
          <div>
            <FileListEdit ref="fileListEdit_image" :multi="true" :mediaType="1" :max="1024"
              :objList="getFileByGuidStr(formData.image,formData.fileObjList)" />
          </div>
          <div>
            <span
              style="font-family: SourceHanSansCN-Regular; font-size: 12px; font-weight: normal; line-height: 20px;letter-spacing: -0.01px;">只能上传jpg/png文件，且不超过1024kb</span>
          </div>
        </div>
      </el-form-item>
      <el-form-item label="状态:" prop="status">
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
    createMemDepart,
    updateMemDepart,
    findMemDepart
  } from '@/api/memDepart'
import { error } from "@/router/enums";


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
  const status_options = ref([])

  const treeOptions = ref([])
  //图片处理
  //import FileListEdit from '@/components/mediaLib/fileListEdit.vue'
  const fileListEdit_image = ref(null)

  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findMemDepart({
      id: id.value
    })
    if (res.code === 200) {
      formData.value = res.data
      formData.value.pidList = getTreeFullPath(treeOptions.value, formData.value.pid);
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
      // let guidList_image = fileListEdit_image.value.getGuidList()
      // formData.value.image = guidList_image.length > 0 ? guidList_image.join(",") : ""
      if (formData.value.pidList && formData.value.pidList.length > 0)
      	  formData.value.pid = formData.value.pidList[formData.value.pidList.length - 1]
      else
         formData.value.pid = 0
      if (formData.value.id >0 &&  formData.value.id == formData.value.pid){
         message("不能选自己作为上级部门",{"type":"error"})
         return
      }
      let res;
      if (id.value > 0) { //update
        res = await updateMemDepart(formData.value)
      } else {
        formData.value.status = 1
        res = await createMemDepart(formData.value)
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
    status_options.value = await getDict("status")
  }
  const getTreeData = async () => {
    let treeDataReq = {
      table: "mem_depart",
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
    userId: 0,
    name: '',
    encode: '',
    desc: '',
    address: '',
    phone: '',
    email: '',
    sort: 0,
    image: "",
    status: 0,
    pidList: [],
  })

  const editRules = ({
  })
	const defaultProps = ref({
		checkStrictly:true,
		expandTrigger:"hover"
	})

	const el_tree_props = ref({
		children:'children',
		label:'label',
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
