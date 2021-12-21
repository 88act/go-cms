<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="唯一id:"> 
              <el-input v-model="formData.guid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件名:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="模块名:">
                 <el-select v-model="formData.module" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in moduleOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="媒体类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="存储位置:">
                 <el-select v-model="formData.driver" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in driverOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="文件路径:"> 
              <el-input v-model="formData.path" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件类型:"> 
              <el-input v-model="formData.ext" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文件大小[k]:">
                 <el-input v-model.number="formData.size" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="md5值:"> 
              <el-input v-model="formData.md5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="sha散列值:"> 
              <el-input v-model="formData.sha1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="下载次数:">
                 <el-input v-model.number="formData.download" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用次数:">
                 <el-input v-model.number="formData.usedTime" clearable placeholder="请输入" />
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
  createBasicFile,
  updateBasicFile,
  findBasicFile
} from '@/api/basicFile' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BasicFile',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      moduleOptions: [],
      media_typeOptions: [],
      driverOptions: [],
      formData: {
           guid: '',
            userId: 0,
           name: '',
            module: 0,
            mediaType: 0,
            driver: 0,
           path: '',
           ext: '',
            size: 0,
           md5: '',
           sha1: '',
            sort: 0,
            download: 0,
            usedTime: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBasicFile({ID:id})
      if (res.code === 200) {
        this.formData = res.data.basicFile
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('module')
    await this.getDict('media_type')
    await this.getDict('driver') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createBasicFile(this.formData)
          break
        case 'update':
          res = await updateBasicFile(this.formData)
          break
        default:
          res = await createBasicFile(this.formData)
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
