<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="机构标识:"> 
              <el-input v-model="formData.jgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医护人员账号:"> 
              <el-input v-model="formData.yhrygh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医护人员姓名:"> 
              <el-input v-model="formData.yhryxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:"> 
              <el-input v-model="formData.xb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证号:"> 
              <el-input v-model="formData.sfzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件类别代码:"> 
              <el-input v-model="formData.zjlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="所属科室代码:"> 
              <el-input v-model="formData.ksdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专业技术职务代码:"> 
              <el-input v-model="formData.zyjszwdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专业技术职务类别代码:"> 
              <el-input v-model="formData.zyjszwlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资质类别名称:"> 
              <el-input v-model="formData.zzlbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格证书编号:"> 
              <el-input v-model="formData.zgzsbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格获得时间:"> 
              <el-input v-model="formData.zghdsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业证书编码:"> 
              <el-input v-model="formData.zyzsbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发证日期:"> 
              <el-input v-model="formData.fzrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业地点:"> 
              <el-input v-model="formData.zydd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业范围:"> 
              <el-input v-model="formData.zyfw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要执业医疗机构代码:"> 
              <el-input v-model="formData.zyzyyljgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要执业医院名称:"> 
              <el-input v-model="formData.zyzyyymc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="全科医生标志:"> 
              <el-input v-model="formData.qkysbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号码:"> 
              <el-input v-model="formData.sjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="专长:"> 
              <el-input v-model="formData.zc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族:"> 
              <el-input v-model="formData.mz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参加工作日期:"> 
              <el-input v-model="formData.cjgzrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注册日期时间:"> 
              <el-input v-model="formData.zcsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="学历:"> 
              <el-input v-model="formData.xl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="个人照片存放地址:"> 
              <el-input v-model="formData.grzpcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="资格证存放地址:"> 
              <el-input v-model="formData.zgzcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执业证存放地址:"> 
              <el-input v-model="formData.zyzcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成日期时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="撤销标志:"> 
              <el-input v-model="formData.cxbz" clearable placeholder="请输入" />
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
  createBbPiStaff,
  updateBbPiStaff,
  findBbPiStaff
} from '@/api/bbPiStaff' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiStaff',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           yhrygh: '',
           yhryxm: '',
           xb: '',
           csrq: '',
           sfzh: '',
           zjlbdm: '',
           ksdm: '',
           zyjszwdm: '',
           zyjszwlbdm: '',
           zzlbmc: '',
           zgzsbh: '',
           zghdsj: '',
           zyzsbm: '',
           fzrq: '',
           zydd: '',
           zyfw: '',
           zyzyyljgdm: '',
           zyzyyymc: '',
           qkysbz: '',
           sjhm: '',
           zc: '',
           mz: '',
           cjgzrq: '',
           zcsj: '',
           xl: '',
           grzpcfdz: '',
           zgzcfdz: '',
           zyzcfdz: '',
           sjscsj: '',
           tbrqsj: '',
           cxbz: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiStaff({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiStaff
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status_up') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createBbPiStaff(this.formData)
          break
        case 'update':
          res = await updateBbPiStaff(this.formData)
          break
        default:
          res = await createBbPiStaff(this.formData)
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
