<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类别ID:">
                 <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统类别ID:">
                 <el-input v-model.number="formData.catIdSys" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章类型:">
                 <el-select v-model="formData.mediaType" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in media_typeOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="文章标题:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章摘要:"> 
              <el-input v-model="formData.sketch" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文章内容:">
              <editor ref="editor_detail" :value="formData.detail" placeholder="请输入文章内容" />
       </el-form-item>
        <el-form-item label="文章作者:"> 
              <el-input v-model="formData.author" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标签列表:"> 
              <el-input v-model="formData.tagList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="作者邮箱:"> 
              <el-input v-model="formData.authorEmail" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="来源:"> 
              <el-input v-model="formData.referer" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="插图:">
               <ImageView ref="imageView_thumb" be-edit :url="getMapData(formData.thumb,formData.mapData)" :guid="formData.thumb" />
       </el-form-item>
        <el-form-item label="二维码图片:"> 
              <el-input v-model="formData.qrcode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="图片列表:"> 
              <el-input v-model="formData.imageList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体列表:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="链接地址:"> 
              <el-input v-model="formData.link" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="置顶:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="热门:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isHot" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="总评论:">
                 <el-input v-model.number="formData.totalDiscuss" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总点击:">
                 <el-input v-model.number="formData.totalClick" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总评:">
                 <el-input v-model.number="formData.totalStar" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评1:">
                 <el-input v-model.number="formData.totalStar1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评2:">
                 <el-input v-model.number="formData.totalStar2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评3:">
                 <el-input v-model.number="formData.totalStar3" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评4:">
                 <el-input v-model.number="formData.totalStar4" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评5:">
                 <el-input v-model.number="formData.totalStar5" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序:">
                 <el-input v-model.number="formData.sort" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="父id:">
                 <el-input v-model.number="formData.pid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="章节集合:"> 
              <el-input v-model="formData.chapter" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="审核信息:"> 
              <el-input v-model="formData.verifyMsg" clearable placeholder="请输入" />
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
  createCmsArticle,
  updateCmsArticle,
  findCmsArticle
} from '@/api/cmsArticle' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑CmsArticle',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      media_typeOptions: [],
      statusOptions: [],
      formData: {
            userId: 0,
            catId: 0,
            catIdSys: 0,
            mediaType: 0,
           name: '',
           sketch: '',
           detail: '',
           author: '',
           tagList: '',
           authorEmail: '',
           referer: '',
            thumb: "",
           qrcode: '',
           imageList: '',
           mediaList: '',
           link: '',
           isTop: false,
           isHot: false,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            totalStar1: 0,
            totalStar2: 0,
            totalStar3: 0,
            totalStar4: 0,
            totalStar5: 0,
            sort: 0,
            pid: 0,
           chapter: '',
            status: 0,
           verifyMsg: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findCmsArticle({ID:id})
      if (res.code === 200) {
        this.formData = res.data.cmsArticle
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
      this.formData.detail = this.$refs.editor_detail.getContent();
      this.formData.thumb = this.$refs.imageView_thumb.myGuid;  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createCmsArticle(this.formData)
          break
        case 'update':
          res = await updateCmsArticle(this.formData)
          break
        default:
          res = await createCmsArticle(this.formData)
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
