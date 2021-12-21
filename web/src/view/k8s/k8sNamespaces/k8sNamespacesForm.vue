<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="命名空间:"> 
              <el-input v-model="formData.namespace" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:"> 
              <el-input v-model="formData.status" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:"> 
              <el-input v-model="formData.createTime" clearable placeholder="请输入" />
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
  createK8sNamespaces,
  updateK8sNamespaces,
  findK8sNamespaces
} from '@/api/k8sNamespaces' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑K8sNamespaces',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      formData: {
           namespace: '',
           status: '',
           createTime: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findK8sNamespaces({ID:id})
      if (res.code === 200) {
        this.formData = res.data.k8sNamespaces
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    } 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createK8sNamespaces(this.formData)
          break
        case 'update':
          res = await updateK8sNamespaces(this.formData)
          break
        default:
          res = await createK8sNamespaces(this.formData)
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
