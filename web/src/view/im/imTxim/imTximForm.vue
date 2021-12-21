<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:"> 
              <el-input v-model="formData.name" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="appid:"> 
              <el-input v-model="formData.appId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="管理员帐号:"> 
              <el-input v-model="formData.identifier" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="签名:"> 
              <el-input v-model="formData.userSig" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="运行次数:">
                 <el-input v-model.number="formData.runTimes" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开始时间:"> 
              <el-input v-model="formData.beginTime" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="当前时间:"> 
              <el-input v-model="formData.nowTime" clearable placeholder="请输入" />
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
  createImTxim,
  updateImTxim,
  findImTxim
} from '@/api/imTxim' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑ImTxim',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
           name: '',
           appId: '',
           identifier: '',
           userSig: '',
            runTimes: 0,
           beginTime: '',
           nowTime: '',
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findImTxim({ID:id})
      if (res.code === 200) {
        this.formData = res.data.imTxim
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
          res = await createImTxim(this.formData)
          break
        case 'update':
          res = await updateImTxim(this.formData)
          break
        default:
          res = await createImTxim(this.formData)
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
