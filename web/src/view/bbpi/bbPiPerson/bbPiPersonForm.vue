<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="userid:">
                 <el-input v-model.number="formData.userid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="上传:">
                 <el-select v-model="formData.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value" />
                 </el-select>
       </el-form-item>
        <el-form-item label="机构标识:"> 
              <el-input v-model="formData.jgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件号码:"> 
              <el-input v-model="formData.zjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证件类别:"> 
              <el-input v-model="formData.zjlbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别代码:"> 
              <el-input v-model="formData.xbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别名称:"> 
              <el-input v-model="formData.xbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="患者属性:"> 
              <el-input v-model="formData.hzsx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="婚姻状况代码:"> 
              <el-input v-model="formData.hyztdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="婚姻状态名称:"> 
              <el-input v-model="formData.hyztmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族代码:"> 
              <el-input v-model="formData.mzdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="民族名称:"> 
              <el-input v-model="formData.mzmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国籍代码:"> 
              <el-input v-model="formData.gjdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="国籍名称:"> 
              <el-input v-model="formData.gjmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机构内部档案号:"> 
              <el-input v-model="formData.jgnbdah" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="固定电话:"> 
              <el-input v-model="formData.gddhhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="手机号码:"> 
              <el-input v-model="formData.sjhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="电子邮件:"> 
              <el-input v-model="formData.dzyj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文化程度代码:"> 
              <el-input v-model="formData.whcddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="文化程度名称:"> 
              <el-input v-model="formData.whcdmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职业类别代码:"> 
              <el-input v-model="formData.zylbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="职业类别名称:"> 
              <el-input v-model="formData.zylbmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生地行政区划码:"> 
              <el-input v-model="formData.csdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生地:"> 
              <el-input v-model="formData.csd" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="居住地行政区划码:"> 
              <el-input v-model="formData.jzdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="居住地址:"> 
              <el-input v-model="formData.jzdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="户口地行政区划码:"> 
              <el-input v-model="formData.hkdxzqhm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="户口地址:"> 
              <el-input v-model="formData.hkdz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人姓名:"> 
              <el-input v-model="formData.lxrxm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="联系人电话:"> 
              <el-input v-model="formData.lxrdh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="ABO 血型:"> 
              <el-input v-model="formData.abo" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="RH 血型:"> 
              <el-input v-model="formData.rh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="参保类别代码:"> 
              <el-input v-model="formData.cblbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="个人档案ID:"> 
              <el-input v-model="formData.grdaid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留一:"> 
              <el-input v-model="formData.yl1" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="预留二:"> 
              <el-input v-model="formData.yl2" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="数据生成时间:"> 
              <el-input v-model="formData.sjscsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="填报日期:"> 
              <el-input v-model="formData.tbrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
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
  createBbPiPerson,
  updateBbPiPerson,
  findBbPiPerson
} from '@/api/bbPiPerson' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiPerson',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            userid: 0,
            status: 0,
           jgdm: '',
           kh: '',
           klx: '',
           zjhm: '',
           zjlbdm: '',
           xm: '',
           xbdm: '',
           xbmc: '',
           hzsx: '',
           hyztdm: '',
           hyztmc: '',
           csrq: '',
           mzdm: '',
           mzmc: '',
           gjdm: '',
           gjmc: '',
           jgnbdah: '',
           gddhhm: '',
           sjhm: '',
           dzyj: '',
           whcddm: '',
           whcdmc: '',
           zylbdm: '',
           zylbmc: '',
           csdxzqhm: '',
           csd: '',
           jzdxzqhm: '',
           jzdz: '',
           hkdxzqhm: '',
           hkdz: '',
           lxrxm: '',
           lxrdh: '',
           abo: '',
           rh: '',
           cblbdm: '',
           grdaid: '',
           yl1: '',
           yl2: '',
           sjscsj: '',
           tbrqsj: '',
           mj: '',
           cxbz: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiPerson({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiPerson
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
          res = await createBbPiPerson(this.formData)
          break
        case 'update':
          res = await updateBbPiPerson(this.formData)
          break
        default:
          res = await createBbPiPerson(this.formData)
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
