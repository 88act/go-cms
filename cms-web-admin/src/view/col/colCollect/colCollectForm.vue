<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="路径:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="翻页路径:"> 
              <el-input v-model="formData.urlPage" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对应表:"> 
              <el-input v-model="formData.toTable" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前id:"> 
              <el-input v-model="formData.nowId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前页码:">
                 <el-input v-model.number="formData.pageNow" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开始页码:">
                 <el-input v-model.number="formData.pageStart" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束页码:">
                 <el-input v-model.number="formData.pageEnd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运行状态:">
                 <el-select v-model="formData.statusRun" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_runOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
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
  createColCollect,
  updateColCollect,
  findColCollect
} from '@/api/colCollect' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ColCollect',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_runOptions: [],
      statusOptions: [],
      formData: {
            userid: 0,
           name: '',
           url: '',
           urlPage: '',
           toTable: '',
           nowId: '',
            pageNow: 0,
            pageStart: 0,
            pageEnd: 0,
            statusRun: 0,
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findColCollect({ID:id})
      if (res.code === 200) {
        this.formData = res.data.colCollect
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status_run')
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
          res = await createColCollect(this.formData)
          break
        case 'update':
          res = await updateColCollect(this.formData)
          break
        default:
          res = await createColCollect(this.formData)
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
