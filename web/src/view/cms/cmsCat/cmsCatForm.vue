<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="父ID:">
                 <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统分类:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beSys" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="群组id:">
                 <el-input v-model.number="formData.groupId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="配图:">
               <ImageView ref="imageView_thumb" be-edit :url="getMapData(formData.thumb,formData.mapData)" :guid="formData.thumb" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否导航:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beNav" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="描述:">
              <editor ref="editor_desc" :value="formData.desc" placeholder="请输入描述" />
       </el-form-item>
        <el-form-item label="关键词:"> 
              <el-input v-model="formData.keywords" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="别名:"> 
              <el-input v-model="formData.alias" clearable placeholder="请输入" />
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
  createCmsCat,
  updateCmsCat,
  findCmsCat
} from '@/api/cmsCat' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑CmsCat',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      media_typeOptions: [],
      statusOptions: [],
      formData: {
            pid: 0,
           beSys: false,
            groupId: 0,
            mediaType: 0,
           name: '',
            thumb: "",
            sort: 0,
           beNav: false,
           desc: '',
           keywords: '',
           alias: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findCmsCat({ID:id})
      if (res.code === 200) {
        this.formData = res.data.cmsCat
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('media_type')
    await this.getDict('status') 
  },
  methods: {
    async save() {
      this.formData.thumb = this.$refs.imageView_thumb.myGuid; 
      this.formData.desc = this.$refs.editor_desc.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createCmsCat(this.formData)
          break
        case 'update':
          res = await updateCmsCat(this.formData)
          break
        default:
          res = await createCmsCat(this.formData)
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
