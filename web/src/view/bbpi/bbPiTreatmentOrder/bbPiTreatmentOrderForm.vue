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
        <el-form-item label="服务网点代码:"> 
              <el-input v-model="formData.fwwddm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方编号:"> 
              <el-input v-model="formData.cfbh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细ID:"> 
              <el-input v-model="formData.cfmxid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡号:"> 
              <el-input v-model="formData.kh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="卡类型:"> 
              <el-input v-model="formData.klx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立日期:"> 
              <el-input v-model="formData.cfklsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方有效天数:"> 
              <el-input v-model="formData.cfyxts" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方科室编码:"> 
              <el-input v-model="formData.cfklksbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方科室名称:"> 
              <el-input v-model="formData.cfklksmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="门诊号:"> 
              <el-input v-model="formData.mzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="姓名:"> 
              <el-input v-model="formData.xm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别代码:"> 
              <el-input v-model="formData.xbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="出生日期:"> 
              <el-input v-model="formData.csrq" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄岁:"> 
              <el-input v-model="formData.nls" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="年龄月:"> 
              <el-input v-model="formData.nly" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="就诊日期时间:"> 
              <el-input v-model="formData.jzrqsj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医嘱说明:"> 
              <el-input v-model="formData.yzsm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="排序号:"> 
              <el-input v-model="formData.pxh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="医嘱项目类型代码:"> 
              <el-input v-model="formData.yzxmlxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药品处方属性:"> 
              <el-input v-model="formData.ypcfsx" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药类别代码:"> 
              <el-input v-model="formData.zylbdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细医保编码:"> 
              <el-input v-model="formData.cfmxxmbm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药监药品ID:"> 
              <el-input v-model="formData.ypid" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方明细名称:"> 
              <el-input v-model="formData.cfmxmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物名称:"> 
              <el-input v-model="formData.ywmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药品规格:"> 
              <el-input v-model="formData.ypgg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="DDD值:"> 
              <el-input v-model="formData.dddz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物剂型代码:"> 
              <el-input v-model="formData.ywjxdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用次剂量:"> 
              <el-input v-model="formData.ywsycjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用剂量单位:"> 
              <el-input v-model="formData.ywsyjldw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用频次代码:"> 
              <el-input v-model="formData.ywsypcdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用频次名称:"> 
              <el-input v-model="formData.ywsypcmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用药途径代码:"> 
              <el-input v-model="formData.yytjdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用药途径名称:"> 
              <el-input v-model="formData.yytjmc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="药物使用总剂量:"> 
              <el-input v-model="formData.ywsyzjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方药品组号:"> 
              <el-input v-model="formData.cfypzh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片处方:"> 
              <el-input v-model="formData.zyypcf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片剂数:"> 
              <el-input v-model="formData.zyypjs" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药饮片煎煮法:"> 
              <el-input v-model="formData.zyypjzf" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="中药用药方法:"> 
              <el-input v-model="formData.zyyyff" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发药剂量:"> 
              <el-input v-model="formData.fyjl" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="发药剂量单位:"> 
              <el-input v-model="formData.fyjldw" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="单价:"> 
              <el-input v-model="formData.dj" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总金额:"> 
              <el-input v-model="formData.zje" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立医师工号:"> 
              <el-input v-model="formData.cfklysgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方开立医师签名:"> 
              <el-input v-model="formData.cfklysqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方审核工号:"> 
              <el-input v-model="formData.cfshyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方审核签名:"> 
              <el-input v-model="formData.cfshyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方调配工号:"> 
              <el-input v-model="formData.cftpyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方调配签名:"> 
              <el-input v-model="formData.cftpyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方核对工号:"> 
              <el-input v-model="formData.cfhdyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方核对签名:"> 
              <el-input v-model="formData.cfhdyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方发药工号:"> 
              <el-input v-model="formData.cffyyjsgh" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="处方发药签名:"> 
              <el-input v-model="formData.cffyyjsqm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="执行结果:"> 
              <el-input v-model="formData.zxjg" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.bz" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取药机构代码:"> 
              <el-input v-model="formData.qyjgdm" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="取药机构名称:"> 
              <el-input v-model="formData.qyjgmc" clearable placeholder="请输入" />
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
        <el-form-item label="密级:"> 
              <el-input v-model="formData.mj" clearable placeholder="请输入" />
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
  createBbPiTreatmentOrder,
  updateBbPiTreatmentOrder,
  findBbPiTreatmentOrder
} from '@/api/bbPiTreatmentOrder' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑BbPiTreatmentOrder',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      status_upOptions: [],
      formData: {
            status: 0,
           jgdm: '',
           fwwddm: '',
           cfbh: '',
           cfmxid: '',
           kh: '',
           klx: '',
           cfklsj: '',
           cfyxts: '',
           cfklksbm: '',
           cfklksmc: '',
           mzh: '',
           xm: '',
           xbdm: '',
           csrq: '',
           nls: '',
           nly: '',
           jzrqsj: '',
           yzsm: '',
           pxh: '',
           yzxmlxdm: '',
           ypcfsx: '',
           zylbdm: '',
           cfmxxmbm: '',
           ypid: '',
           cfmxmc: '',
           ywmc: '',
           ypgg: '',
           dddz: '',
           ywjxdm: '',
           ywsycjl: '',
           ywsyjldw: '',
           ywsypcdm: '',
           ywsypcmc: '',
           yytjdm: '',
           yytjmc: '',
           ywsyzjl: '',
           cfypzh: '',
           zyypcf: '',
           zyypjs: '',
           zyypjzf: '',
           zyyyff: '',
           fyjl: '',
           fyjldw: '',
           dj: '',
           zje: '',
           cfklysgh: '',
           cfklysqm: '',
           cfshyjsgh: '',
           cfshyjsqm: '',
           cftpyjsgh: '',
           cftpyjsqm: '',
           cfhdyjsgh: '',
           cfhdyjsqm: '',
           cffyyjsgh: '',
           cffyyjsqm: '',
           zxjg: '',
           bz: '',
           qyjgdm: '',
           qyjgmc: '',
           sjscsj: '',
           tbrqsj: '',
           cxbz: '',
           mj: '',
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findBbPiTreatmentOrder({ID:id})
      if (res.code === 200) {
        this.formData = res.data.bbPiTreatmentOrder
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
          res = await createBbPiTreatmentOrder(this.formData)
          break
        case 'update':
          res = await updateBbPiTreatmentOrder(this.formData)
          break
        default:
          res = await createBbPiTreatmentOrder(this.formData)
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
