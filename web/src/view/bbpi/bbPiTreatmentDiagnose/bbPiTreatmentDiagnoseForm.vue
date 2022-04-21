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
        <el-form-item label="处方编号:"> 
              <el-input v-model="formData.cfbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细ID:"> 
              <el-input v-model="formData.cfmxid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断信息 ID:"> 
              <el-input v-model="formData.zdxxid" clearable placeholder="请输入" />
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
        <el-form-item label="就诊时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断类型编码:"> 
              <el-input v-model="formData.zdlxbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="西医诊断编码:"> 
              <el-input v-model="formData.xyzdbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="西医诊断名称:"> 
              <el-input v-model="formData.xyzdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="西医诊断编码类型:"> 
              <el-input v-model="formData.xyzdbmlx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中医诊断编码:"> 
              <el-input v-model="formData.zyzdbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中医诊断名称:"> 
              <el-input v-model="formData.zyzdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中医诊断编码类型:"> 
              <el-input v-model="formData.zyzdbmlx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断说明:"> 
              <el-input v-model="formData.zdsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断标志:"> 
              <el-input v-model="formData.zdbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断医生工号:"> 
              <el-input v-model="formData.zdysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断医生姓名:"> 
              <el-input v-model="formData.zdysxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="诊断时间:"> 
              <el-input v-model="formData.zdsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留一:"> 
              <el-input v-model="formData.yl1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="撤销标志:"> 
              <el-input v-model="formData.cxbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
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
  createBbPiTreatmentDiagnose,
  updateBbPiTreatmentDiagnose,
  findBbPiTreatmentDiagnose
} from '@/api/bbPiTreatmentDiagnose' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiTreatmentDiagnose',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           cfbh: '',
           cfmxid: '',
           zdxxid: '',
           kh: '',
           klx: '',
           mzh: '',
           xm: '',
           xbdm: '',
           csrq: '',
           nls: '',
           nly: '',
           jzrqsj: '',
           zdlxbm: '',
           xyzdbm: '',
           xyzdmc: '',
           xyzdbmlx: '',
           zyzdbm: '',
           zyzdmc: '',
           zyzdbmlx: '',
           zdsm: '',
           zdbz: '',
           zdysgh: '',
           zdysxm: '',
           zdsj: '',
           yl1: '',
           sjscsj: '',
           tbrqsj: '',
           cxbz: '',
           mj: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiTreatmentDiagnose({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiTreatmentDiagnose
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
          res = await createBbPiTreatmentDiagnose(this.formData)
          break
        case 'update':
          res = await updateBbPiTreatmentDiagnose(this.formData)
          break
        default:
          res = await createBbPiTreatmentDiagnose(this.formData)
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
