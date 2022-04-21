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
        <el-form-item label="就诊类型:"> 
              <el-input v-model="formData.jzlx" clearable placeholder="请输入" />
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
        <el-form-item label="科室编码:"> 
              <el-input v-model="formData.ksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="科室名称:"> 
              <el-input v-model="formData.ksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别代码:"> 
              <el-input v-model="formData.xbdm" clearable placeholder="请输入" />
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
        <el-form-item label="过敏史标识:"> 
              <el-input v-model="formData.gmsbs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="过敏史:"> 
              <el-input v-model="formData.gms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参保类别:"> 
              <el-input v-model="formData.cblb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊日期时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊诊断说明:"> 
              <el-input v-model="formData.jzzdsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊标志代码:"> 
              <el-input v-model="formData.czbzdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主诉:"> 
              <el-input v-model="formData.zs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="现病史:"> 
              <el-input v-model="formData.xbs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="既往史:"> 
              <el-input v-model="formData.jws" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辅助检查项目:"> 
              <el-input v-model="formData.fzjcxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辅助检查结果:"> 
              <el-input v-model="formData.fzjcjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状-诊断代码:"> 
              <el-input v-model="formData.mzzzzddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状-诊断名称:"> 
              <el-input v-model="formData.mzzzzdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊症状诊断编码类型:"> 
              <el-input v-model="formData.mzzzzdbmlx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="症状描述:"> 
              <el-input v-model="formData.zzms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="辨证依据:"> 
              <el-input v-model="formData.bzyj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="治则治法:"> 
              <el-input v-model="formData.zzzf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="健康指导:"> 
              <el-input v-model="formData.jkzd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处置计划:"> 
              <el-input v-model="formData.czjh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应诊医生工号:"> 
              <el-input v-model="formData.yzysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="应诊医师签名:"> 
              <el-input v-model="formData.yzysqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊医疗卫生机构代码:"> 
              <el-input v-model="formData.czylwsjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="初诊机构名称:"> 
              <el-input v-model="formData.czylwsjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊结束时间:"> 
              <el-input v-model="formData.jzjssj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="转诊标志:"> 
              <el-input v-model="formData.zzbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者满意度:"> 
              <el-input v-model="formData.hzmyd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留一:"> 
              <el-input v-model="formData.yl1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留二:"> 
              <el-input v-model="formData.yl2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子处方地址:"> 
              <el-input v-model="formData.dzcfwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医生端视频地址:"> 
              <el-input v-model="formData.ysdspwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者端视频地址:"> 
              <el-input v-model="formData.hzdspwjcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="交流音频地址:"> 
              <el-input v-model="formData.jlypcfdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成日期时间:"> 
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
  createBbPiTreatmentRecord,
  updateBbPiTreatmentRecord,
  findBbPiTreatmentRecord
} from '@/api/bbPiTreatmentRecord' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiTreatmentRecord',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           jzlx: '',
           kh: '',
           klx: '',
           mzh: '',
           ksbm: '',
           ksmc: '',
           xm: '',
           xbdm: '',
           csrq: '',
           nls: '',
           nly: '',
           gmsbs: '',
           gms: '',
           cblb: '',
           jzrqsj: '',
           jzzdsm: '',
           czbzdm: '',
           zs: '',
           xbs: '',
           jws: '',
           fzjcxm: '',
           fzjcjg: '',
           mzzzzddm: '',
           mzzzzdmc: '',
           mzzzzdbmlx: '',
           zzms: '',
           bzyj: '',
           zzzf: '',
           jkzd: '',
           czjh: '',
           yzysgh: '',
           yzysqm: '',
           czylwsjgdm: '',
           czylwsjgmc: '',
           jzjssj: '',
           zzbz: '',
           hzmyd: '',
           yl1: '',
           yl2: '',
           dzcfwjcfdz: '',
           ysdspwjcfdz: '',
           hzdspwjcfdz: '',
           jlypcfdz: '',
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
      const res = await findBbPiTreatmentRecord({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiTreatmentRecord
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
          res = await createBbPiTreatmentRecord(this.formData)
          break
        case 'update':
          res = await updateBbPiTreatmentRecord(this.formData)
          break
        default:
          res = await createBbPiTreatmentRecord(this.formData)
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
