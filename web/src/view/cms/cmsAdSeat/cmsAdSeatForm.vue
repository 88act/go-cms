<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="位置名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告位宽度:">
                 <el-input v-model.number="formData.width" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告位高度:">
                 <el-input v-model.number="formData.height" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告描述:">
              <editor ref="editor_desc" :value="formData.desc" placeholder="请输入广告描述" />
       </el-form-item>
        <el-form-item label="模板:"> 
              <el-input v-model="formData.style" clearable placeholder="请输入" />
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
  createCmsAdSeat,
  updateCmsAdSeat,
  findCmsAdSeat
} from '@/api/cmsAdSeat' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑CmsAdSeat',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
           name: '',
            width: 0,
            height: 0,
           desc: '',
           style: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findCmsAdSeat({ID:id})
      if (res.code === 200) {
        this.formData = res.data.cmsAdSeat
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status') 
  },
  methods: {
    async save() { 
      this.formData.desc = this.$refs.editor_desc.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createCmsAdSeat(this.formData)
          break
        case 'update':
          res = await updateCmsAdSeat(this.formData)
          break
        default:
          res = await createCmsAdSeat(this.formData)
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
