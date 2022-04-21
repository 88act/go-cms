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
        <el-form-item label="服务网点代码:"> 
              <el-input v-model="formData.fwwddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊记录:"> 
              <el-input v-model="formData.zzjlid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊号:"> 
              <el-input v-model="formData.mzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:"> 
              <el-input v-model="formData.xb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄岁:"> 
              <el-input v-model="formData.nls" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄月:"> 
              <el-input v-model="formData.nly" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责医生工号:"> 
              <el-input v-model="formData.fzysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责医生姓名:"> 
              <el-input v-model="formData.fzysxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责科室编码:"> 
              <el-input v-model="formData.fzksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责科室名称:"> 
              <el-input v-model="formData.fzksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊原因:"> 
              <el-input v-model="formData.zzyy" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊时间:"> 
              <el-input v-model="formData.zzsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向机构代码:"> 
              <el-input v-model="formData.zxjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向机构名称:"> 
              <el-input v-model="formData.zxjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向科室代码:"> 
              <el-input v-model="formData.zxksdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向科室名称:"> 
              <el-input v-model="formData.zxksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向医生工号:"> 
              <el-input v-model="formData.zxysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转向医生姓名:"> 
              <el-input v-model="formData.zxysxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主要既往史:"> 
              <el-input v-model="formData.zyjws" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="治疗经过:"> 
              <el-input v-model="formData.zljg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下一步治疗方案:"> 
              <el-input v-model="formData.xybzlfa" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊标志:"> 
              <el-input v-model="formData.zzbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
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
  createBbPiTreatmentReferral,
  updateBbPiTreatmentReferral,
  findBbPiTreatmentReferral
} from '@/api/bbPiTreatmentReferral' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiTreatmentReferral',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           zzjlid: '',
           kh: '',
           klx: '',
           mzh: '',
           xm: '',
           xb: '',
           csrq: '',
           nls: '',
           nly: '',
           jzrqsj: '',
           fzysgh: '',
           fzysxm: '',
           fzksbm: '',
           fzksmc: '',
           zzyy: '',
           zzsj: '',
           zxjgdm: '',
           zxjgmc: '',
           zxksdm: '',
           zxksmc: '',
           zxysgh: '',
           zxysxm: '',
           zyjws: '',
           zljg: '',
           xybzlfa: '',
           zzbz: '',
           sjscsj: '',
           tbrqsj: '',
           mj: '',
           cxbz: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiTreatmentReferral({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiTreatmentReferral
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
          res = await createBbPiTreatmentReferral(this.formData)
          break
        case 'update':
          res = await updateBbPiTreatmentReferral(this.formData)
          break
        default:
          res = await createBbPiTreatmentReferral(this.formData)
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
