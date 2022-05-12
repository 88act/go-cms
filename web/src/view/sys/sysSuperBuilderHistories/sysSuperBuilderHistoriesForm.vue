<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="表名:"> 
              <el-input v-model="formData.tabName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表单:">
              <editor ref="editor_requestMeta" :value="formData.requestMeta" placeholder="请输入表单" />
       </el-form-item>
        <el-form-item label="路径:">
              <editor ref="editor_superBuilderPath" :value="formData.superBuilderPath" placeholder="请输入路径" />
       </el-form-item>
        <el-form-item label="注射:">
              <editor ref="editor_injectionMeta" :value="formData.injectionMeta" placeholder="请输入注射" />
       </el-form-item>
        <el-form-item label="struct名称:"> 
              <el-input v-model="formData.structName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="struct中文名:"> 
              <el-input v-model="formData.structCnName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="api:"> 
              <el-input v-model="formData.apiIds" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标记:">
                 <el-input v-model.number="formData.flag" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表名:"> 
              <el-input v-model="formData.tableName" clearable placeholder="请输入" />
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
  createSysSuperBuilderHistories,
  updateSysSuperBuilderHistories,
  findSysSuperBuilderHistories
} from '@/api/sysSuperBuilderHistories' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑SysSuperBuilderHistories',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      formData: {
           tabName: '',
           requestMeta: '',
           superBuilderPath: '',
           injectionMeta: '',
           structName: '',
           structCnName: '',
           apiIds: '',
            flag: 0,
           tableName: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findSysSuperBuilderHistories({ID:id})
      if (res.code === 200) {
        this.formData = res.data.sysSuperBuilderHistories
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    } 
  },
  methods: {
    async save() { 
      this.formData.requestMeta = this.$refs.editor_requestMeta.getContent(); 
      this.formData.superBuilderPath = this.$refs.editor_superBuilderPath.getContent(); 
      this.formData.injectionMeta = this.$refs.editor_injectionMeta.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createSysSuperBuilderHistories(this.formData)
          break
        case 'update':
          res = await updateSysSuperBuilderHistories(this.formData)
          break
        default:
          res = await createSysSuperBuilderHistories(this.formData)
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
