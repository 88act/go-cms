<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="是否上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="表名:"> 
              <el-input v-model="formData.tabName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表名cn:"> 
              <el-input v-model="formData.tabNameCn" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="字段名:"> 
              <el-input v-model="formData.tabField" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="字段名cn:"> 
              <el-input v-model="formData.tabFieldCn" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
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
  createBbPiUpField,
  updateBbPiUpField,
  findBbPiUpField
} from '@/api/bbPiUpField' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiUpField',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           tabName: '',
           tabNameCn: '',
           tabField: '',
           tabFieldCn: '',
            sort: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiUpField({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiUpField
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
          res = await createBbPiUpField(this.formData)
          break
        case 'update':
          res = await updateBbPiUpField(this.formData)
          break
        default:
          res = await createBbPiUpField(this.formData)
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
