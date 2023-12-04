<template>
		<div class="gocms-form-box bg-bg_color">
			<el-form ref="editForm" :model="formData" :rules="editRules"  label-position="right" label-width="80px" >

        <el-form-item label="文章标题:"  prop="title">
                <el-input v-model="formData.title" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章摘要:"  prop="desc">
                <el-input ref="editor_desc" type="textarea" rows="3" :value="formData.desc" placeholder="请输入文章摘要" />
       </el-form-item>
        <el-form-item label="标签列表:"  prop="tagList">
                <el-input v-model="formData.tagList" clearable placeholder="请输入" />
       </el-form-item>

        <el-form-item label="插图:"  prop="image">
						 <FileListEdit ref="fileListEdit_image" :multi="false" :imgSize="120" :mediaType="1" :max="2048" :objList="getFileByGuidStr(formData.image,formData.fileObjList)" />
       </el-form-item>
        <el-form-item label="链接地址:"  prop="link">
                <el-input v-model="formData.link" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="热门:"  prop="beTop">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beTop" clearable ></el-switch>

       </el-form-item>
        <el-form-item label="综合指数:"  prop="totalWhole">
                 <el-input v-model.number="formData.totalWhole" clearable placeholder="请输入" />
       </el-form-item>

        <el-form-item label="状态:"  prop="status">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_options" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="审核信息:"  prop="verifyMsg">
                <el-input v-model="formData.verifyMsg" clearable placeholder="请输入" />
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
	createCmsArt,
	updateCmsArt,
	findCmsArt
} from '@/api/cmsArt'


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
		const art_type_options = ref([])
		const status_options = ref([])

  const treeOptions = ref([])
  //图片处理
  import FileListEdit from '@/components/mediaLib/fileListEdit.vue'
    const fileListEdit_image = ref(null)
    const fileListEdit_fileList = ref(null)

  // // 查询
  const getData = async () => {
    if (id.value <= 0) {
      console.log("getData id.value <=0")
      return;
    }
    const res = await findCmsArt({
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
			//BeEditor   this.formData.desc = this.$refs.editor_desc.getContent();
		let guidList_image = fileListEdit_image.value.getGuidList()
		formData.value.image =  guidList_image.length >0? guidList_image.join(","):""
		// let guidList_fileList = fileListEdit_fileList.value.getGuidList()
		// formData.value.fileList =  guidList_fileList.length >0? guidList_fileList.join(","):""
      let res;
      if (id.value > 0) { //update
        res = await updateCmsArt(formData.value)
      } else {
        formData.value.status = 1
        res = await createCmsArt(formData.value)
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
			art_type_options.value = await getDict("art_type")
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
            pid: 0,
            userId: 0,
            catId: 0,
            catIdSys: 0,
            type: 0,
           title: '',
           desc: '',
           tagList: '',
           source: '',
            image: "",
            fileList: "",
           link: '',
           beTop: false,
            totalWhole: 0,
            totalShare: 0,
            totalFav: 0,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            totalGood: 0,
            totalPoor: 0,
            status: 0,
           verifyMsg: '',
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

