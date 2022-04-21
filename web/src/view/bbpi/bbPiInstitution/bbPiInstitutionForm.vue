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
        <el-form-item label="统一信用代码:"> 
              <el-input v-model="formData.zzjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构名称:"> 
              <el-input v-model="formData.jgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="行政区划代码:"> 
              <el-input v-model="formData.xzqhdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构类型:"> 
              <el-input v-model="formData.jglx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构成立日期:"> 
              <el-input v-model="formData.jgclrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单位隶属关系代码:"> 
              <el-input v-model="formData.dwlsgxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构分类代码:"> 
              <el-input v-model="formData.jgflgllbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构分类代码:"> 
              <el-input v-model="formData.jgfldm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="经济类型代码:"> 
              <el-input v-model="formData.jjlxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="地址:"> 
              <el-input v-model="formData.dz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="组织机构代码:"> 
              <el-input v-model="formData.styyzzjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院名称:"> 
              <el-input v-model="formData.styymc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医疗机构级别:"> 
              <el-input v-model="formData.styljgjb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院机构等级:"> 
              <el-input v-model="formData.styljgdj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医院网址:"> 
              <el-input v-model="formData.hlwyywz" clearable placeholder="请输入" />
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
        <el-form-item label="信息安等级保护:"> 
              <el-input v-model="formData.xxaqdjbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="信息安等保编号:"> 
              <el-input v-model="formData.xxaqdjbhbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="开办资金金额数:"> 
              <el-input v-model="formData.kbzjjes" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="法人代表:"> 
              <el-input v-model="formData.frdb" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构民族自治标志:"> 
              <el-input v-model="formData.jgszdmzzzdfbz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否分支机构:"> 
              <el-input v-model="formData.sffzjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构第二名称:"> 
              <el-input v-model="formData.jgdemc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构描述:"> 
              <el-input v-model="formData.jgms" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="邮政编码:"> 
              <el-input v-model="formData.yzbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电话号码:"> 
              <el-input v-model="formData.dhhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子信箱:"> 
              <el-input v-model="formData.dwdzyx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第三方信用代码:"> 
              <el-input v-model="formData.dsfjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="第三方机构名称:"> 
              <el-input v-model="formData.dsfjgmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="协议有效期:"> 
              <el-input v-model="formData.xyyxq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="部署平台:"> 
              <el-input v-model="formData.bspt" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="架构描述:"> 
              <el-input v-model="formData.jgmsxx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机房面积:"> 
              <el-input v-model="formData.jfmj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="是否双路发电:"> 
              <el-input v-model="formData.sfslgd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.bz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成日期时间:"> 
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
  createBbPiInstitution,
  updateBbPiInstitution,
  findBbPiInstitution
} from '@/api/bbPiInstitution' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiInstitution',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           zzjgdm: '',
           jgmc: '',
           xzqhdm: '',
           jglx: '',
           jgclrq: '',
           dwlsgxdm: '',
           jgflgllbdm: '',
           jgfldm: '',
           jjlxdm: '',
           dz: '',
           styyzzjgdm: '',
           styymc: '',
           styljgjb: '',
           styljgdj: '',
           hlwyywz: '',
           xkzhm: '',
           xkxmmc: '',
           xkzyxq: '',
           xxaqdjbh: '',
           xxaqdjbhbh: '',
           kbzjjes: '',
           frdb: '',
           jgszdmzzzdfbz: '',
           sffzjg: '',
           jgdemc: '',
           jgms: '',
           yzbm: '',
           dhhm: '',
           dwdzyx: '',
           dsfjgdm: '',
           dsfjgmc: '',
           xyyxq: '',
           bspt: '',
           jgmsxx: '',
           jfmj: '',
           sfslgd: '',
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
      const res = await findBbPiInstitution({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiInstitution
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
          res = await createBbPiInstitution(this.formData)
          break
        case 'update':
          res = await updateBbPiInstitution(this.formData)
          break
        default:
          res = await createBbPiInstitution(this.formData)
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
