<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:"> 
              <el-input v-model="formData.areaname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父栏目:">
                 <el-input v-model.number="formData.parentid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="简称:"> 
              <el-input v-model="formData.shortname" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电话区号:"> 
              <el-input v-model="formData.areacode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮编:"> 
              <el-input v-model="formData.zipcode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拼音:"> 
              <el-input v-model="formData.pinyin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="拼音简写:"> 
              <el-input v-model="formData.py" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经度:"> 
              <el-input v-model="formData.lng" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="纬度:"> 
              <el-input v-model="formData.lat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="级别:">
                 <el-input v-model.number="formData.level" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="position字段:"> 
              <el-input v-model="formData.position" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="组合名称:"> 
              <el-input v-model="formData.mergername" clearable placeholder="请输入" />
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
  createBasicArea,
  updateBasicArea,
  findBasicArea
} from '@/api/basicArea' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BasicArea',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      formData: {
           areaname: '',
            parentid: 0,
           shortname: '',
           areacode: '',
           zipcode: '',
           pinyin: '',
           py: '',
           lng: '',
           lat: '',
            level: 0,
           position: '',
           mergername: '',
            sort: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBasicArea({ID:id})
      if (res.code === 200) {
        this.formData = res.data.basicArea
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
          res = await createBasicArea(this.formData)
          break
        case 'update':
          res = await updateBasicArea(this.formData)
          break
        default:
          res = await createBasicArea(this.formData)
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
