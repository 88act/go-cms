<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="机构标识:"> 
              <el-input v-model="formData.jgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备代号:"> 
              <el-input v-model="formData.sbdh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备名称:"> 
              <el-input v-model="formData.sbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="同批设备台数:"> 
              <el-input v-model="formData.tpsbts" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="产地:"> 
              <el-input v-model="formData.cd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="生产厂家:"> 
              <el-input v-model="formData.sccj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备型号:"> 
              <el-input v-model="formData.sbxh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备参数:"> 
              <el-input v-model="formData.sbcs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备类型:"> 
              <el-input v-model="formData.sblx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="购买日期:"> 
              <el-input v-model="formData.gmrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用途:"> 
              <el-input v-model="formData.yt" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="设备价值金额:"> 
              <el-input v-model="formData.sbjzje" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="购进时新旧情况:"> 
              <el-input v-model="formData.gjsxqk" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="理论设计寿命:"> 
              <el-input v-model="formData.llsjsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="使用情况:"> 
              <el-input v-model="formData.syqk" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.bz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="撤销标志:"> 
              <el-input v-model="formData.cxbz" clearable placeholder="请输入" />
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
  createBbPiDevice,
  updateBbPiDevice,
  findBbPiDevice
} from '@/api/bbPiDevice' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiDevice',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           sbdh: '',
           sbmc: '',
           tpsbts: '',
           cd: '',
           sccj: '',
           sbxh: '',
           sbcs: '',
           sblx: '',
           gmrq: '',
           yt: '',
           sbjzje: '',
           gjsxqk: '',
           llsjsm: '',
           syqk: '',
           bz: '',
           sjscsj: '',
           tbrqsj: '',
           cxbz: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiDevice({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiDevice
        this.editType = 'update'
      }
    } else {
      this.editType = 'create'
    }
    await this.getDict('status_up') 
  },
  methods: {
    async save() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case 'create':
          res = await createBbPiDevice(this.formData)
          break
        case 'update':
          res = await updateBbPiDevice(this.formData)
          break
        default:
          res = await createBbPiDevice(this.formData)
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
