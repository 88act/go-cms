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
        <el-form-item label="月份:"> 
              <el-input v-model="formData.nf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="派出机构数量:"> 
              <el-input v-model="formData.pcjgsl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职工总数:"> 
              <el-input v-model="formData.zgzs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客户服务人员总数:"> 
              <el-input v-model="formData.khffryzs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="日均坐诊医生数:"> 
              <el-input v-model="formData.rjzzyspbs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务用房面积:"> 
              <el-input v-model="formData.ywyfmj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总收入:"> 
              <el-input v-model="formData.zsr" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总支出:"> 
              <el-input v-model="formData.zzc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总资产:"> 
              <el-input v-model="formData.zzch" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="流动资产:"> 
              <el-input v-model="formData.ldzc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对外投资:"> 
              <el-input v-model="formData.dwtz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="固定资产:"> 
              <el-input v-model="formData.gdzc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="无形资产及开办费:"> 
              <el-input v-model="formData.wxzcjkbf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负债:"> 
              <el-input v-model="formData.fz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="净资产:"> 
              <el-input v-model="formData.jzc" clearable placeholder="请输入" />
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
  createBbPiInstitutionBusiness,
  updateBbPiInstitutionBusiness,
  findBbPiInstitutionBusiness
} from '@/api/bbPiInstitutionBusiness' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiInstitutionBusiness',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           nf: '',
           pcjgsl: '',
           zgzs: '',
           khffryzs: '',
           rjzzyspbs: '',
           ywyfmj: '',
           zsr: '',
           zzc: '',
           zzch: '',
           ldzc: '',
           dwtz: '',
           gdzc: '',
           wxzcjkbf: '',
           fz: '',
           jzc: '',
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
      const res = await findBbPiInstitutionBusiness({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiInstitutionBusiness
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
          res = await createBbPiInstitutionBusiness(this.formData)
          break
        case 'update':
          res = await updateBbPiInstitutionBusiness(this.formData)
          break
        default:
          res = await createBbPiInstitutionBusiness(this.formData)
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
