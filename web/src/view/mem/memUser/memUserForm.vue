<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户名:"> 
              <el-input v-model="formData.username" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码:"> 
              <el-input v-model="formData.password" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密码盐:"> 
              <el-input v-model="formData.passwordSlat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮件:"> 
              <el-input v-model="formData.email" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:"> 
              <el-input v-model="formData.mobile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="昵称:"> 
              <el-input v-model="formData.nickname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="真实名:"> 
              <el-input v-model="formData.realname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证:"> 
              <el-input v-model="formData.cardId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:">
                 <el-select v-model="formData.sex" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in sexOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="生日:">
                <el-date-picker v-model="formData.birthday" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="头像:"> 
              <el-input v-model="formData.avatar" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="验证手机:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.mobileValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="验证邮箱:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.emailValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="验证实名:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.cardidValidated" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.info" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="推荐码16位（自己的）:"> 
              <el-input v-model="formData.recommendCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="安全状态:">
                 <el-select v-model="formData.statusSafe" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in usersafe_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="注册ip:">
                 <el-input v-model.number="formData.regIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登录ip:">
                 <el-input v-model.number="formData.loginIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="登录次数:">
                 <el-input v-model.number="formData.loginTotal" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最后登录时间:">
                <el-date-picker v-model="formData.loginTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import {
  createMemUser,
  updateMemUser,
  findMemUser
} from '@/api/memUser' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑MemUser',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      sexOptions: [],
      statusOptions: [],
      usersafe_typeOptions: [],
      formData: {
           username: '',
           password: '',
           passwordSlat: '',
           email: '',
           mobile: '',
           nickname: '',
           realname: '',
           cardId: '',
            sex: 0,
            birthday: new Date(),
           avatar: '',
           mobileValidated: false,
           emailValidated: false,
           cardidValidated: false,
           info: '',
           recommendCode: '',
            status: 0,
            statusSafe: 0,
            regIp: 0,
            loginIp: 0,
            loginTotal: 0,
            loginTime: new Date(),
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findMemUser({ID:id})
      if (res.code === 200) {
        this.formData = res.data.memUser
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('sex')
    await this.getDict('status')
    await this.getDict('usersafe_type') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createMemUser(this.formData)
          break
        case 'update':
          res = await updateMemUser(this.formData)
          break
        default:
          res = await createMemUser(this.formData)
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
         emitter.emit('closeThisPage') 
      }
    },
    back() {
      this.$router.go(-1)
      emitter.emit('closeThisPage') 
    }
  }
}
</script>
<style>
</style>
