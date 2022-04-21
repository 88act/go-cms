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
        <el-form-item label="统一社会信用代码:"> 
              <el-input v-model="formData.zzjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务网点代码:"> 
              <el-input v-model="formData.fwwddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点名称:"> 
              <el-input v-model="formData.fwdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="行政区划代码:"> 
              <el-input v-model="formData.xzqhdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点类型:"> 
              <el-input v-model="formData.fwdlx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点成立日期:"> 
              <el-input v-model="formData.fwdclrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位隶属关系代码:"> 
              <el-input v-model="formData.dwlsgxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点分类管理类别代码:"> 
              <el-input v-model="formData.fwdflgllbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点分类代码:"> 
              <el-input v-model="formData.fwdfldm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经济类型代码:"> 
              <el-input v-model="formData.jjlxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.dz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点医院级别:"> 
              <el-input v-model="formData.fwdyyjb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点医院等级:"> 
              <el-input v-model="formData.fwdyydj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可证号码:"> 
              <el-input v-model="formData.xkzhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可项目名称:"> 
              <el-input v-model="formData.xkxmmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="许可证有效期:"> 
              <el-input v-model="formData.xkzyxq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开办资金金额数:"> 
              <el-input v-model="formData.kbzjjes" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="法人代表:"> 
              <el-input v-model="formData.frdb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务点地方标志:"> 
              <el-input v-model="formData.fwdszdmzzzdfbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否分支机构:"> 
              <el-input v-model="formData.sffzjg" clearable placeholder="请输入" />
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
  createBbPiServicePoint,
  updateBbPiServicePoint,
  findBbPiServicePoint
} from '@/api/bbPiServicePoint' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiServicePoint',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           zzjgdm: '',
           fwwddm: '',
           fwdmc: '',
           xzqhdm: '',
           fwdlx: '',
           fwdclrq: '',
           dwlsgxdm: '',
           fwdflgllbdm: '',
           fwdfldm: '',
           jjlxdm: '',
           dz: '',
           fwdyyjb: '',
           fwdyydj: '',
           xkzhm: '',
           xkxmmc: '',
           xkzyxq: '',
           kbzjjes: '',
           frdb: '',
           fwdszdmzzzdfbz: '',
           sffzjg: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiServicePoint({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiServicePoint
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
          res = await createBbPiServicePoint(this.formData)
          break
        case 'update':
          res = await updateBbPiServicePoint(this.formData)
          break
        default:
          res = await createBbPiServicePoint(this.formData)
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
