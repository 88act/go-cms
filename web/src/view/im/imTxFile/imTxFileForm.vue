<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="消息id:">
                 <el-input v-model.number="formData.msgId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="远程地址:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="本地路径:"> 
              <el-input v-model="formData.local" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下载状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_downloadOptions" :key="key" :label="item.label" :value="item.value" />
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
  createImTxFile,
  updateImTxFile,
  findImTxFile
} from '@/api/imTxFile' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ImTxFile',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      media_typeOptions: [],
      status_downloadOptions: [],
      formData: {
            msgId: 0,
            mediaType: 0,
           url: '',
           local: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findImTxFile({ID:id})
      if (res.code === 200) {
        this.formData = res.data.imTxFile
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('media_type')
    await this.getDict('status_download') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createImTxFile(this.formData)
          break
        case 'update':
          res = await updateImTxFile(this.formData)
          break
        default:
          res = await createImTxFile(this.formData)
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
