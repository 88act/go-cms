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
        <el-form-item label="订单类型:">
                 <el-input v-model.number="formData.orderType" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="对象id:">
                 <el-input v-model.number="formData.objId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="实付价:">
                 <el-input v-model.number="formData.price" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="原价:">
                 <el-input v-model.number="formData.priceSrc" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="优惠券ID:">
                 <el-input v-model.number="formData.couponId" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="备注:"> 
              <el-input v-model="formData.remark" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="核销码:"> 
              <el-input v-model="formData.orderCode" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="支付状态:">
                 <el-input v-model.number="formData.statusPay" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="订单状态:">
                 <el-input v-model.number="formData.statusOrder" clearable placeholder="请输入" />
       </el-form-item>
        <el-form-item label="记录状态:">
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
  createOrderOrder,
  updateOrderOrder,
  findOrderOrder
} from '@/api/orderOrder' //  此处请自行替换地址
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm'
import { emitter } from '@/utils/bus.js' 
export default {
 name: '编辑OrderOrder',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      statusOptions: [],
      formData: {
           sn: '',
            userId: 0,
            orderType: 0,
            objId: 0,
            price: 0,
            priceSrc: 0,
            couponId: 0,
           remark: '',
           orderCode: '',
            statusPay: 0,
            statusOrder: 0,
            status: 0,
            mapData: {}
      }
    }
  },
  async created() {
    let id = this.$route.params.id;
    if (id && id >0) {
      const res = await findOrderOrder({ID:id})
      if (res.code === 200) {
        this.formData = res.data.orderOrder
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
          res = await createOrderOrder(this.formData)
          break
        case 'update':
          res = await updateOrderOrder(this.formData)
          break
        default:
          res = await createOrderOrder(this.formData)
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
