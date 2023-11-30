<template>
      <div class="gocms-form-box bg-bg_color" style="display: flex;  justify-content: flex-start; ">

          <div style="display: flex; flex-wrap: wrap; justify-content: center; width:350px">
             <ReQrcode :text="qrcodeStr" :width="300" />
              <div style="vertical-align: bottom;">
                <span>二维码有效期: 5分钟 </span>
                 <el-button  class="el-btn-save" type="primary" @click="refQrcode()" > 刷 新</el-button>
             </div>
           </div>


	 			    <div style="padding-top: 30px; display: flex; flex-wrap: wrap; justify-content: flex-start;  width:450px">

            <el-form ref="editForm" :model="formData" :rules="editRules"  :inline="true"  label-width="100px">

              <el-form-item  label="用户名:" >
	 							 <div style="width: 180px;">  {{useUserStoreHook().username}}</div>
	 						</el-form-item>
              <el-form-item  label="旧密码:" prop="password">
	 							<el-input style="width: 180px;" v-model="formData.password" show-password />
	 						</el-form-item>
	 						<el-form-item  label="新密码:" prop="newPassword">
	 							<el-input style="width: 180px;" v-model="formData.newPassword" show-password />
	 						</el-form-item>
              <el-form-item  label="确认密码:" prop="newPassword2">
              	<el-input style="width: 180px;" v-model="formData.newPassword2" show-password />
              </el-form-item>
	 					</el-form>
	 					<el-button style="margin-left: 50px;" class="el-btn-save" type="primary"
            @click="savePassword">修改密码</el-button>
          </div>

    </div>
</template>

<script setup lang="tsx">

   import {
     ref,
     onMounted,
     computed,
     h,
     createVNode,
     type CSSProperties,
   } from "vue";
   import {
     message
   } from "@/utils/message";
   import {
     storageLocal
   } from "@pureadmin/utils";
   import {
     useUserStoreHook
   } from "@/store/modules/user";
   import {
     usePermissionStoreHook
   } from "@/store/modules/permission";
   import {
     useRenderIcon
   } from "@/components/ReIcon/src/hooks";
   import {
     ElMessage,
     ElMessageBox
   } from 'element-plus'

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
     addDialog,
     closeDialog,
     updateDialog,
     closeAllDialog
   } from "@/components/ReDialog";

 import {
    required,
    username,
    password,
    range,
    len,
    okChat
  } from "@/utils/formRules"

 import {
	changePwdMy,
  showQrcode
 } from '@/api/base'

 import ReQrcode from "@/components/ReQrcode";

 const editId = ref(0)
 const editForm = ref(null)
 const formData = ref({
    userId: 0,
    username:"",
    password:'',
    newPassword: '',
    newPassword2: ''
  })


	const editRules = ({
    password: [required,password],
		newPassword: [required,password],
		newPassword2: [required,password],
	})


  const savePassword = async () => {

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
     if (formData.value.newPassword  != formData.value.newPassword2 ) {
         message("两次密码不一样", { type: "error" })
         return
     }

      let res = await changePwdMy(formData.value)
      if (res.code == 200) {
        message(res.msg, { type: "success" })
      } else {
        message(res.msg, { type: "error" })
      }
    }
  }
  const qrcodeStr =ref("")
  async function refQrcode() {
   let data = {
    "id":useUserStoreHook().userId
   }
   let res = await showQrcode(data)
   if (res.code == 200) {
     // qrcodeShow.value = true
      qrcodeStr.value = res.data
     //  ElMessage.success(res.msg)
   } else  message(res.msg,{type:"error"})
}


	//编辑
	const goEditForm = (id) => {
	   editId.value = id
	 //   addDialog({
		// 	title: "编 辑",
		// 	fullscreenIcon: true,
		// 	hideFooter: true,
		// 	contentRenderer: ({ options, index }) =>
		// 		h(MemUserForm, {
		// 		editId: editId.value,
		// 		beChange: beChange.value,
		// 		index: index,
		// 		options: options,
		// 		"onUpdate:editId": val => (editId.value = val),
		// 		"onUpdate:beChange": val => (beChange.value = val)
		// 		}),
		// 	closeCallBack: () => {
		// 		if (beChange.value) {
		// 		    getTableData()
		// 		}
		// 	}
		// });
	}
	const init = () => {
     refQrcode()
	}
	onMounted(() => {
     init()
   })
</script>


<style lang="scss" scoped>
	.div-info {
		height: 280px;
		padding: 10px;
		margin: 20px;
		width:90%;
	}

	.box-userInfo {}
</style>

