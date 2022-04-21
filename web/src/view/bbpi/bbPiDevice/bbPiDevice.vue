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
                <el-form-item label="上传" prop="status">                
                    <el-select v-model="searchInfo.status" placeholder="请选择" clearable>
                      <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value"></el-option>
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
          <!--status  BeQuickEdit -->
        <el-table-column label="上传" prop="status" width="120"  sortable="custom" >
        <template #default="scope">  
        <el-popover trigger="click" placement="top"  width = "280">  
              <el-select v-model="scope.row.status" placeholder="请选择"  @change="quickEdit_do('status',scope.row.ID,scope.row.status,scope)">
                  <el-option v-for="(item,key) in status_upOptions" :key="key" :label="item.label" :value="item.value"></el-option>
              </el-select> 
              <template #reference>
                  <div class="quickEdit" > {{filterDict(scope.row.status,"status_up")}} </div>
              </template>
            </el-popover>
        </template>  
        </el-table-column>
          <!--jgdm  BeQuickEdit -->  
        <el-table-column label="机构标识" prop="jgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jgdm',scope.row.ID,scope.row.jgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sbdh  BeQuickEdit -->  
        <el-table-column label="设备代号" prop="sbdh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sbdh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sbdh',scope.row.ID,scope.row.sbdh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sbdh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sbmc  BeQuickEdit -->  
        <el-table-column label="设备名称" prop="sbmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sbmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sbmc',scope.row.ID,scope.row.sbmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sbmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tpsbts  BeQuickEdit -->  
        <el-table-column label="同批设备台数" prop="tpsbts" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tpsbts"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('tpsbts',scope.row.ID,scope.row.tpsbts,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tpsbts}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cd  BeQuickEdit -->  
        <el-table-column label="产地" prop="cd" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cd"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cd',scope.row.ID,scope.row.cd,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cd}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sccj  BeQuickEdit -->  
        <el-table-column label="生产厂家" prop="sccj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sccj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sccj',scope.row.ID,scope.row.sccj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sccj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sbxh  BeQuickEdit -->  
        <el-table-column label="设备型号" prop="sbxh" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sbxh"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sbxh',scope.row.ID,scope.row.sbxh,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sbxh}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sbcs  BeQuickEdit -->  
        <el-table-column label="设备参数" prop="sbcs" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sbcs"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sbcs',scope.row.ID,scope.row.sbcs,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sbcs}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sblx  BeQuickEdit -->  
        <el-table-column label="设备类型" prop="sblx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sblx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sblx',scope.row.ID,scope.row.sblx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sblx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gmrq  BeQuickEdit -->  
        <el-table-column label="购买日期" prop="gmrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gmrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gmrq',scope.row.ID,scope.row.gmrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gmrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--yt  BeQuickEdit -->  
        <el-table-column label="用途" prop="yt" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.yt"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('yt',scope.row.ID,scope.row.yt,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.yt}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sbjzje  BeQuickEdit -->  
        <el-table-column label="设备价值金额" prop="sbjzje" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sbjzje"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sbjzje',scope.row.ID,scope.row.sbjzje,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sbjzje}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--gjsxqk  BeQuickEdit -->  
        <el-table-column label="购进时新旧情况" prop="gjsxqk" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.gjsxqk"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('gjsxqk',scope.row.ID,scope.row.gjsxqk,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.gjsxqk}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--llsjsm  BeQuickEdit -->  
        <el-table-column label="理论设计寿命" prop="llsjsm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.llsjsm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('llsjsm',scope.row.ID,scope.row.llsjsm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.llsjsm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--syqk  BeQuickEdit -->  
        <el-table-column label="使用情况" prop="syqk" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.syqk"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('syqk',scope.row.ID,scope.row.syqk,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.syqk}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--bz  BeQuickEdit -->  
        <el-table-column label="备注" prop="bz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.bz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('bz',scope.row.ID,scope.row.bz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.bz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sjscsj  BeQuickEdit -->  
        <el-table-column label="数据生成时间" prop="sjscsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sjscsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sjscsj',scope.row.ID,scope.row.sjscsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sjscsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--tbrqsj  BeQuickEdit -->  
        <el-table-column label="填报日期" prop="tbrqsj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.tbrqsj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('tbrqsj',scope.row.ID,scope.row.tbrqsj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.tbrqsj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--cxbz  BeQuickEdit -->  
        <el-table-column label="撤销标志" prop="cxbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.cxbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('cxbz',scope.row.ID,scope.row.cxbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.cxbz}} </div>
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
  createBbPiDevice, 
  deleteBbPiDeviceByIds,
  updateBbPiDevice,
  findBbPiDevice,
  getBbPiDeviceList,
  quickEdit,
  excelList
} from '@/api/bbPiDevice' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiDevice',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiDeviceList,   
      excelListApi: excelList,
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
    await this.getDict('status_up')
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
        //this.deleteBbPiDevice(row)
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
     const res = await deleteBbPiDeviceByIds({ ids })
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
			this.$router.push({ name:'bbPiDeviceForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiDeviceForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiDevice({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiDevice 
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
          res = await createBbPiDevice(this.formData);
          break
        case "update": 
          res = await updateBbPiDevice(this.formData);
          break
        default: 
          res = await createBbPiDevice(this.formData);
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
