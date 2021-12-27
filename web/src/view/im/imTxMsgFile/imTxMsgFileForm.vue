<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="消息类型:"> 
              <el-input v-model="formData.chatType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消息时间:"> 
              <el-input v-model="formData.msgTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下载路径:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="过期时间:">
                <el-date-picker v-model="formData.expireTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="文件大小:">
                 <el-input v-model.number="formData.fileSize" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件md5:"> 
              <el-input v-model="formData.fileMd5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="压缩大小:">
                 <el-input v-model.number="formData.gzipSize" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="压缩md5:"> 
              <el-input v-model="formData.gzipMd5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求code:"> 
              <el-input v-model="formData.errorCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求信息:"> 
              <el-input v-model="formData.errorInfo" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="请求状态:"> 
              <el-input v-model="formData.actionStatus" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="本地路径:"> 
              <el-input v-model="formData.localFile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
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
  createImTxMsgFile,
  updateImTxMsgFile,
  findImTxMsgFile
} from '@/api/imTxMsgFile' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ImTxMsgFile',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
           chatType: '',
           msgTime: '',
           url: '',
            expireTime: new Date(),
            fileSize: 0,
           fileMd5: '',
            gzipSize: 0,
           gzipMd5: '',
           errorCode: '',
           errorInfo: '',
           actionStatus: '',
           localFile: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findImTxMsgFile({ID:id})
      if (res.code === 200) {
        this.formData = res.data.imTxMsgFile
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createImTxMsgFile(this.formData)
          break
        case 'update':
          res = await updateImTxMsgFile(this.formData)
          break
        default:
          res = await createImTxMsgFile(this.formData)
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
