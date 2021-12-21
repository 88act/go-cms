<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="节点名称:"> 
              <el-input v-model="formData.nodeName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ip地址:"> 
              <el-input v-model="formData.ip" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点状态:"> 
              <el-input v-model="formData.status" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点角色:"> 
              <el-input v-model="formData.roles" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建时间:"> 
              <el-input v-model="formData.createTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点版本:"> 
              <el-input v-model="formData.version" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点资源:"> 
              <el-input v-model="formData.resource" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="节点标签:"> 
              <el-input v-model="formData.label" clearable placeholder="请输入" />
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
  createK8sNodes,
  updateK8sNodes,
  findK8sNodes
} from '@/api/k8sNodes' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑K8sNodes',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      formData: {
           nodeName: '',
           ip: '',
           status: '',
           roles: '',
           createTime: '',
           version: '',
           resource: '',
           label: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findK8sNodes({ID:id})
      if (res.code === 200) {
        this.formData = res.data.k8sNodes
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
          res = await createK8sNodes(this.formData)
          break
        case 'update':
          res = await updateK8sNodes(this.formData)
          break
        default:
          res = await createK8sNodes(this.formData)
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
