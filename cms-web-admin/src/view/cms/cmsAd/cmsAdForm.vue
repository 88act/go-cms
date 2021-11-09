<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="广告位置:">
                 <el-input v-model.number="formData.seatId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="是否新窗:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beTarget" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="广告图片:">
               <ImageView ref="imageView_image" be-edit :url="getMapData(formData.image,formData.mapData)" :guid="formData.image" />
       </el-form-item>
        <el-form-item label="广告链接:"> 
              <el-input v-model="formData.link" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="广告内容:">
              <editor ref="editor_detail" :value="formData.detail" placeholder="请输入广告内容" />
       </el-form-item>
        <el-form-item label="过期内容:">
              <editor ref="editor_expDetail" :value="formData.expDetail" placeholder="请输入过期内容" />
       </el-form-item>
        <el-form-item label="投放时间:">
               <el-date-picker type="datetimerange" v-model="formData.startTime" format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss" :style="{width: '100%'}" start-placeholder="开始日期"
                  end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
       </el-form-item>
        <el-form-item label="结束时间:">
               <el-date-picker type="datetimerange" v-model="formData.endTime" format="yyyy-MM-dd HH:mm:ss"
                  value-format="yyyy-MM-dd HH:mm:ss" :style="{width: '100%'}" start-placeholder="开始日期"
                  end-placeholder="结束日期" range-separator="至" clearable></el-date-picker>
       </el-form-item>
        <el-form-item label="点击量:">
                 <el-input v-model.number="formData.totalClick" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
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
  createCmsAd,
  updateCmsAd,
  findCmsAd
} from '@/api/cmsAd' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑CmsAd',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      media_typeOptions: [],
      statusOptions: [],
      formData: {
            seatId: 0,
           name: '',
            mediaType: 0,
           beTarget: false,
            image: "",
           link: '',
           detail: '',
           expDetail: '',
            startTime: new Date(),
            endTime: new Date(),
            totalClick: 0,
            sort: 0,
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findCmsAd({ID:id})
      if (res.code === 200) {
        this.formData = res.data.cmsAd
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
      this.formData.image = this.$refs.imageView_image.myGuid; 
      this.formData.detail = this.$refs.editor_detail.getContent(); 
      this.formData.expDetail = this.$refs.editor_expDetail.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createCmsAd(this.formData)
          break
        case 'update':
          res = await updateCmsAd(this.formData)
          break
        default:
          res = await createCmsAd(this.formData)
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
