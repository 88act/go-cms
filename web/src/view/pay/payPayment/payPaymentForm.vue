<template>
  <div>
    <div class="gocms-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="订单号:"> 
              <el-input v-model="formData.sn" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="用户id:">
                 <el-input v-model.number="formData.userId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付方式:"> 
              <el-input v-model="formData.payMode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="三方支付类型:"> 
              <el-input v-model="formData.tradeType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="三方交易状态:"> 
              <el-input v-model="formData.tradeState" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总金额:">
                 <el-input v-model.number="formData.price" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="三方支付单号:"> 
              <el-input v-model="formData.transactionId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付状态:"> 
              <el-input v-model="formData.tradeStateDesc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务单号:"> 
              <el-input v-model="formData.orderSn" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="业务类型:"> 
              <el-input v-model="formData.serviceType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付状态:">
                 <el-input v-model.number="formData.statusPay" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付时间:">
                <el-date-picker v-model="formData.payTime" type="datetime" style="width:100%" placeholder="选择时间日期" clearable />
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
  createPayPayment,
  updatePayPayment,
  findPayPayment
} from '@/api/payPayment' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑PayPayment',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
           sn: '',
            userId: 0,
           payMode: '',
           tradeType: '',
           tradeState: '',
            price: 0,
           transactionId: '',
           tradeStateDesc: '',
           orderSn: '',
           serviceType: '',
            statusPay: 0,
            payTime: new Date(),
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findPayPayment({ID:id})
      if (res.code === 200) {
        this.formData = res.data.payPayment
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
          res = await createPayPayment(this.formData)
          break
        case 'update':
          res = await updatePayPayment(this.formData)
          break
        default:
          res = await createPayPayment(this.formData)
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
