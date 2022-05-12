<template>
  <div>  
  <!----------查询form------------------ -->
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline"> 
      <el-form-item label="创建时间">
            <el-date-picker 
                  v-model="searchInfo.createdAtBetween" 
                  type="datetimerange"
                  format="YYYY-MM-DD HH:mm:ss"
                  :shortcuts="shortcuts"
                  range-separator="至"
                  start-placeholder="开始日期"
                  end-placeholder="结束日期"
                />
              </el-form-item> 
        <el-form-item label="ID">
            <el-input placeholder="搜索ID" v-model="searchInfo.ID" />
        </el-form-item> 
                <el-form-item label="订单号">
                  <el-input placeholder="搜索条件" v-model="searchInfo.sn" clearable />
                </el-form-item>
                  <el-form-item label="用户id">
                      <el-input placeholder="搜索条件" v-model="searchInfo.userId" clearable />
                  </el-form-item> 
                <el-form-item label="支付方式">
                  <el-input placeholder="搜索条件" v-model="searchInfo.payMode" clearable />
                </el-form-item> 
                <el-form-item label="三方交易状态">
                  <el-input placeholder="搜索条件" v-model="searchInfo.tradeState" clearable />
                </el-form-item>
                <el-form-item label="状态" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
                    </el-select>
                </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="el-icon-search" @click="onSearch">查询</el-button>
          <el-button size="mini" type="primary" icon="el-icon-plus" @click="goEditForm(0)">新增</el-button>
           <el-button size="mini" type="success"  @click="onExcel">导出当前</el-button>
         <el-button size="mini" type="success"  @click="onExcelAll">导出全部</el-button>
          <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
              <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
              <el-button icon="el-icon-delete" size="mini" type="danger" style="margin-left: 10px;">批量删除</el-button>
            </template>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
   <!----------数据表------------------ -->
    <el-table
      ref="multipleTable"
      border
      stripe
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      @selection-change="handleSelectionChange"
      @sort-change="sortChange" 
    >
      <el-table-column type="selection" width="55" />      
       <el-table-column label="ID" min-width="60" prop="ID" sortable="custom" />
          <!--sn  BeQuickEdit -->  
        <el-table-column label="订单号" prop="sn" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sn"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sn',scope.row.ID,scope.row.sn,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sn}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--userId  BeQuickEdit -->  
        <el-table-column label="用户id" prop="userId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.userId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('user_id',scope.row.ID,scope.row.userId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.userId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--payMode  BeQuickEdit -->  
        <el-table-column label="支付方式" prop="payMode" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.payMode"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('pay_mode',scope.row.ID,scope.row.payMode,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.payMode}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tradeType  BeQuickEdit -->  
        <el-table-column label="三方支付类型" prop="tradeType" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tradeType"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('trade_type',scope.row.ID,scope.row.tradeType,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tradeType}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tradeState  BeQuickEdit -->  
        <el-table-column label="三方交易状态" prop="tradeState" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tradeState"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('trade_state',scope.row.ID,scope.row.tradeState,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tradeState}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--price  BeQuickEdit -->  
        <el-table-column label="总金额" prop="price" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.price"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('price',scope.row.ID,scope.row.price,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.price}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--transactionId  BeQuickEdit -->  
        <el-table-column label="三方支付单号" prop="transactionId" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.transactionId"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('transaction_id',scope.row.ID,scope.row.transactionId,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.transactionId}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tradeStateDesc  BeQuickEdit -->  
        <el-table-column label="支付状态" prop="tradeStateDesc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tradeStateDesc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('trade_state_desc',scope.row.ID,scope.row.tradeStateDesc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tradeStateDesc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--orderSn  BeQuickEdit -->  
        <el-table-column label="业务单号" prop="orderSn" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.orderSn"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('order_sn',scope.row.ID,scope.row.orderSn,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.orderSn}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--serviceType  BeQuickEdit -->  
        <el-table-column label="业务类型" prop="serviceType" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.serviceType"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('service_type',scope.row.ID,scope.row.serviceType,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.serviceType}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--statusPay  BeQuickEdit -->  
        <el-table-column label="支付状态" prop="statusPay" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.statusPay"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('status_pay',scope.row.ID,scope.row.statusPay,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.statusPay}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--payTime  BeQuickEdit -->  
        <el-table-column label="支付时间" prop="payTime" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.payTime"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('pay_time',scope.row.ID,scope.row.payTime,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.payTime}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--status  BeQuickEdit -->
        <el-table-column label="状态" prop="status" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">
                  <el-option v-for="(item,key) in statusOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.status,"status")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column> 
      <el-table-column label="日期" width="180" prop="created_at" sortable="custom" >
        <template #default="scope">{{ formatDate(scope.row.CreatedAt)}}</template>
      </el-table-column>
      
      <el-table-column label="操作">
        <template #default="scope">
          <el-button plain size="mini" type="primary" icon="el-icon-edit" class="table-button" @click="goEditForm(scope.row.ID)">编辑</el-button>
          <el-button plain size="mini" type="danger" icon="el-icon-delete"  @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      layout="total, sizes, prev, pager, next, jumper"
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10,20,50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
    />
    <!---------- 编辑弹窗------------------ -->
    <el-dialog  v-if="dialogFormVisible"  :before-close="closeDialog" v-model="dialogFormVisible" title="编辑资料">
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
     </el-form>
      <div slot="footer" class="el-dialog__footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button type="primary" @click="saveEditForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  createPayPayment, 
  deletePayPaymentByIds,
  updatePayPayment,
  findPayPayment,
  getPayPaymentList,
  quickEdit,
  excelList
} from '@/api/payPayment' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'PayPayment',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getPayPaymentList,   
      excelListApi: excelList,
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
    await this.getDict('status')
    await this.getTableData()
  },
  methods: { 
    onSearch() {
      this.page = 1
      this.pageSize = 20 
      this.getTableData()
    },
    handleSelectionChange(val) {
      this.multipleSelection = val
    },
    deleteRow(row) {
      this.$confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
         const ids = [row.ID] 
         this.doDelete(ids); 
        //this.deletePayPayment(row)
      })
    },
    async onDelete() {
      const ids = []
      if (this.multipleSelection.length === 0) {
        this.$message({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      this.multipleSelection &&
        this.multipleSelection.map(item => {
          ids.push(item.ID)
        })
      this.doDelete(ids); 
    },
  	async doDelete(ids) { 
     const res = await deletePayPaymentByIds({ ids })
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '删除成功'
        })
        if (this.tableData.length === ids.length && this.page > 1) {
          this.page--
        }
        this.deleteVisible = false
        this.getTableData()
      } 
	},  
   async goEditForm(id) { 
	  if (this.beNewWindow) {
		  if (id >0) {
			this.$router.push({ name:'payPaymentForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'payPaymentForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findPayPayment({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.payPayment 
		 }else
		 {
			this.editType = 'create' 
		 }
		  this.dialogFormVisible = true
	  }
	}, 
    async saveEditForm() {  
      delete this.formData.mapData;
      delete this.formData.CreatedAt;
      delete this.formData.UpdatedAt;
      let res;
      switch (this.editType) {
        case "create":         
          res = await createPayPayment(this.formData);
          break
        case "update": 
          res = await updatePayPayment(this.formData);
          break
        default: 
          res = await createPayPayment(this.formData);
          break
      }
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
        this.closeDialog()
        this.getTableData()
      }
    },   
    quickEdit_do(field,id,value,scope) {    
      let value2 = value;
      if (typeof(value)==="boolean")
         value2 = value?"1":"0"
      value2 =  value2+"";   
      let obj = {field:field,id:id,value:value2}	 
      this.quickEdit_do2(obj);	  
	    // if (scope._self.$refs[`popover-${scope.$index}`])
		  // scope._self.$refs[`popover-${scope.$index}`].doClose();
    },
    async quickEdit_do2(obj) {  
      const res =  await quickEdit(obj)	  
      if (res.code === 200) {
        this.$message({
          type: 'success',
          message: '修改成功'
        }) 
        // this.getTableData()
      }
    },
   onExcel(){
        this.getExcelList(this.page,this.pageSize)    
    },
    onExcelAll(){
        this.getExcelList(1,1000)  
    }
  },
}
</script>
<style>
</style>
