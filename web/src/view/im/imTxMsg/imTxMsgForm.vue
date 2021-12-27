<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="消息类型:"> 
              <el-input v-model="formData.chatType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消息时间:"> 
              <el-input v-model="formData.msgTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发送人:"> 
              <el-input v-model="formData.fromAccount" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="接收人:"> 
              <el-input v-model="formData.toAccount" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="时间撮:">
                <el-date-picker v-model="formData.msgTimestamp" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
       </el-form-item>
        <el-form-item label="seq:">
                 <el-input v-model.number="formData.msgSeq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="random:">
                 <el-input v-model.number="formData.msgRandom" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="消息类型:"> 
              <el-input v-model="formData.msgType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内容:">
              <editor ref="editor_msgContent" :value="formData.msgContent" placeholder="请输入内容" />
       </el-form-item>
        <el-form-item label="文本:"> 
              <el-input v-model="formData.msgText" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体文件:"> 
              <el-input v-model="formData.mediaList" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体远程文件:"> 
              <el-input v-model="formData.mediaListTx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="媒体下载:">
                 <el-select v-model="formData.statusMedia" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_downloadOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="IP地址:"> 
              <el-input v-model="formData.clientIp" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="平台:"> 
              <el-input v-model="formData.msgFromPlatform" clearable placeholder="请输入" />
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
  createImTxMsg,
  updateImTxMsg,
  findImTxMsg
} from '@/api/imTxMsg' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ImTxMsg',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_downloadOptions: [],
      statusOptions: [],
      formData: {
           chatType: '',
           msgTime: '',
           fromAccount: '',
           toAccount: '',
            msgTimestamp: new Date(),
            msgSeq: 0,
            msgRandom: 0,
           msgType: '',
           msgContent: '',
           msgText: '',
           mediaList: '',
           mediaListTx: '',
            statusMedia: 0,
           clientIp: '',
           msgFromPlatform: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findImTxMsg({ID:id})
      if (res.code === 200) {
        this.formData = res.data.imTxMsg
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status_download')
    await this.getDict('status') 
  },
  methods: {
    async save() { 
      this.formData.msgContent = this.$refs.editor_msgContent.getContent();  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createImTxMsg(this.formData)
          break
        case 'update':
          res = await updateImTxMsg(this.formData)
          break
        default:
          res = await createImTxMsg(this.formData)
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
