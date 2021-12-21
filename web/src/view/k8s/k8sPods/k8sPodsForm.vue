<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="容器名称:"> 
              <el-input v-model="formData.podName" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="容器IP:"> 
              <el-input v-model="formData.podIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主机IP:"> 
              <el-input v-model="formData.hostIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:"> 
              <el-input v-model="formData.status" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="启动时间:"> 
              <el-input v-model="formData.startTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="重启次数:">
                 <el-input v-model.number="formData.restartCount" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="命名空间:"> 
              <el-input v-model="formData.nameSpace" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="line:">
                 <el-input v-model.number="formData.line" clearable placeholder="请输入" />
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
  createK8sPods,
  updateK8sPods,
  findK8sPods
} from '@/api/k8sPods' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑K8sPods',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      formData: {
           podName: '',
           podIp: '',
           hostIp: '',
           status: '',
           startTime: '',
            restartCount: 0,
           nameSpace: '',
            line: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findK8sPods({ID:id})
      if (res.code === 200) {
        this.formData = res.data.k8sPods
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
          res = await createK8sPods(this.formData)
          break
        case 'update':
          res = await updateK8sPods(this.formData)
          break
        default:
          res = await createK8sPods(this.formData)
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
