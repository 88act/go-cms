<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="系统栏目:">
                 <el-input v-model.number="formData.catId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户栏目:">
                 <el-input v-model.number="formData.catIdUser" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beOnline" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="活动类型:"> 
              <el-input v-model="formData.actType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="分期活动:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beMulti" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="分期:">
                 <el-input v-model.number="formData.phase" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标题:"> 
              <el-input v-model="formData.title" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="简介:"> 
              <el-input v-model="formData.desc" clearable placeholder="请输入" />
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
        <el-form-item label="开始时间:">
                <el-date-picker v-model="formData.beginTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="结束时间:">
                <el-date-picker v-model="formData.endTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="直播开始:">
                <el-date-picker v-model="formData.liveTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="直播结束:">
                <el-date-picker v-model="formData.liveEnd" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="主持/讲师:"> 
              <el-input v-model="formData.presenter" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="客服电话:"> 
              <el-input v-model="formData.phoneKefu" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="合作电话:"> 
              <el-input v-model="formData.phoneHezu" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="微信:"> 
              <el-input v-model="formData.wx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="QQ:"> 
              <el-input v-model="formData.qq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="群组id:"> 
              <el-input v-model="formData.groupId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="结束报名:">
                <el-date-picker v-model="formData.endJoinTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="票价:">
                 <el-input v-model.number="formData.price" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="vip票价:">
                 <el-input v-model.number="formData.priceVip" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="票价描述:"> 
              <el-input v-model="formData.priceDesc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退费类型:">
                 <el-input v-model.number="formData.refundType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="退费天:">
                 <el-input v-model.number="formData.refundDays" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否显示人数:">
                 <el-input v-model.number="formData.beShowjoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="活动状态:">
                 <el-input v-model.number="formData.statusAct" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="最大报名人数:">
                 <el-input v-model.number="formData.maxJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前报名人数:">
                 <el-input v-model.number="formData.nowJoin" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否精选:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beChosen" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="是否周末:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beWeek" clearable ></el-switch>
              
       </el-form-item>
        <el-form-item label="是否会员:">
             <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.beVip" clearable ></el-switch>
              
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
        <el-form-item label="总报名数:">
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
        <el-form-item label="总星评1:">
                 <el-input v-model.number="formData.totalStar1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评2:">
                 <el-input v-model.number="formData.totalStar2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总星评3:">
                 <el-input v-model.number="formData.totalStar3" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="状态:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="审核原因:"> 
              <el-input v-model="formData.statusMsg" clearable placeholder="请输入" />
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
  createActAct,
  updateActAct,
  findActAct
} from '@/api/actAct' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ActAct',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
            userId: 0,
            catId: 0,
            catIdUser: 0,
           beOnline: false,
           actType: '',
           beMulti: false,
            phase: 0,
           title: '',
           desc: '',
           avater: '',
           mediaList: '',
           address: '',
            areaId: 0,
            lng: 0,
            lat: 0,
            beginTime: new Date(),
            endTime: new Date(),
            liveTime: new Date(),
            liveEnd: new Date(),
           presenter: '',
           phoneKefu: '',
           phoneHezu: '',
           wx: '',
           qq: '',
           groupId: '',
            endJoinTime: new Date(),
            price: 0,
            priceVip: 0,
           priceDesc: '',
            refundType: 0,
            refundDays: 0,
            beShowjoin: 0,
            statusAct: 0,
            maxJoin: 0,
            nowJoin: 0,
           beChosen: false,
           beWeek: false,
           beVip: false,
            totalWhole: 0,
            totalShare: 0,
            totalFav: 0,
            totalJoin: 0,
            totalDiscuss: 0,
            totalClick: 0,
            totalStar: 0,
            totalStar1: 0,
            totalStar2: 0,
            totalStar3: 0,
            status: 0,
           statusMsg: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findActAct({ID:id})
      if (res.code === 200) {
        this.formData = res.data.actAct
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
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
          res = await createActAct(this.formData)
          break
        case 'update':
          res = await updateActAct(this.formData)
          break
        default:
          res = await createActAct(this.formData)
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
