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
        <el-form-item label="科室编码:"> 
              <el-input v-model="formData.ksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="科室名称:"> 
              <el-input v-model="formData.ksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标准科室代码:"> 
              <el-input v-model="formData.bzksdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医保局代码:"> 
              <el-input v-model="formData.ybjdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="科室简介:"> 
              <el-input v-model="formData.ksjs" clearable placeholder="请输入" />
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
  createBbPiDepartment,
  updateBbPiDepartment,
  findBbPiDepartment
} from '@/api/bbPiDepartment' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiDepartment',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           ksbm: '',
           ksmc: '',
           bzksdm: '',
           ybjdm: '',
           ksjs: '',
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
      const res = await findBbPiDepartment({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiDepartment
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
          res = await createBbPiDepartment(this.formData)
          break
        case 'update':
          res = await updateBbPiDepartment(this.formData)
          break
        default:
          res = await createBbPiDepartment(this.formData)
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
