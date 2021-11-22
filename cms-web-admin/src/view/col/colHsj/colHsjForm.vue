<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="采集id:">
                 <el-input v-model.number="formData.colId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="原URL:"> 
              <el-input v-model="formData.url" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标题:"> 
              <el-input v-model="formData.title" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内容:">
              <editor ref="editor_content" :value="formData.content" placeholder="请输入内容" />
       </el-form-item>
        <el-form-item label="发布时间:">
               <el-date-picker type="datetimerange" v-model="formData.pubTime" format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss" :style="{width: '100%'}" start-placeholder="开始日期"
                  end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
       </el-form-item>
        <el-form-item label="发布单位:"> 
              <el-input v-model="formData.pubUnit" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经度:"> 
              <el-input v-model="formData.log" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="纬度:"> 
              <el-input v-model="formData.lat" clearable placeholder="请输入" />
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
  createColHsj,
  updateColHsj,
  findColHsj
} from '@/api/colHsj' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ColHsj',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
            colId: 0,
           url: '',
           title: '',
           content: '',
            pubTime: new Date(),
           pubUnit: '',
           log: '',
           lat: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findColHsj({ID:id})
      if (res.code === 200) {
        this.formData = res.data.colHsj
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status') 
  },
  methods: {
    async save() { 
      this.formData.content = this.$refs.editor_content.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createColHsj(this.formData)
          break
        case 'update':
          res = await updateColHsj(this.formData)
          break
        default:
          res = await createColHsj(this.formData)
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
