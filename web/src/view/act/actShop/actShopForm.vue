<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标题:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="简介:">
              <editor ref="editor_desc" :value="formData.desc" placeholder="请输入简介" />
       </el-form-item>
        <el-form-item label="详细内容:">
              <editor ref="editor_detail" :value="formData.detail" placeholder="请输入详细内容" />
       </el-form-item>
        <el-form-item label="缩略图:"> 
              <el-input v-model="formData.avater" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体列表:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.address" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地区id:">
                 <el-input v-model.number="formData.areaId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经度:">
                 <el-input v-model.number="formData.lng" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="纬度:">
                 <el-input v-model.number="formData.lat" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮件:"> 
              <el-input v-model="formData.email" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机:"> 
              <el-input v-model="formData.mobile" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip等级:">
                 <el-input v-model.number="formData.vipLev" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip结束时间:">
                <el-date-picker v-model="formData.vipTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="综合指数:">
                 <el-input v-model.number="formData.totalWhole" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总分享:">
                 <el-input v-model.number="formData.totalShare" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总收藏:">
                 <el-input v-model.number="formData.totalFav" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总报名人数:">
                 <el-input v-model.number="formData.totalJoin" clearable placeholder="请输入" />
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
  createActShop,
  updateActShop,
  findActShop
} from '@/api/actShop' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ActShop',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
            userId: 0,
           name: '',
           desc: '',
           detail: '',
           avater: '',
           mediaList: '',
           address: '',
            areaId: 0,
            lng: 0,
            lat: 0,
           email: '',
           mobile: '',
            vipLev: 0,
            vipTime: new Date(),
            totalWhole: 0,
            totalShare: 0,
            totalFav: 0,
            totalJoin: 0,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findActShop({ID:id})
      if (res.code === 200) {
        this.formData = res.data.actShop
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
      this.formData.detail = this.$refs.editor_detail.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createActShop(this.formData)
          break
        case 'update':
          res = await updateActShop(this.formData)
          break
        default:
          res = await createActShop(this.formData)
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
