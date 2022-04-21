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
          <!--zzjgdm  BeQuickEdit -->  
        <el-table-column label="统一社会信用代码" prop="zzjgdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.zzjgdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('zzjgdm',scope.row.ID,scope.row.zzjgdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.zzjgdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwwddm  BeQuickEdit -->  
        <el-table-column label="服务网点代码" prop="fwwddm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwwddm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwwddm',scope.row.ID,scope.row.fwwddm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwwddm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdmc  BeQuickEdit -->  
        <el-table-column label="服务点名称" prop="fwdmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdmc',scope.row.ID,scope.row.fwdmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xzqhdm  BeQuickEdit -->  
        <el-table-column label="行政区划代码" prop="xzqhdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xzqhdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xzqhdm',scope.row.ID,scope.row.xzqhdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xzqhdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdlx  BeQuickEdit -->  
        <el-table-column label="服务点类型" prop="fwdlx" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdlx"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdlx',scope.row.ID,scope.row.fwdlx,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdlx}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdclrq  BeQuickEdit -->  
        <el-table-column label="服务点成立日期" prop="fwdclrq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdclrq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdclrq',scope.row.ID,scope.row.fwdclrq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdclrq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dwlsgxdm  BeQuickEdit -->  
        <el-table-column label="单位隶属关系代码" prop="dwlsgxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dwlsgxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dwlsgxdm',scope.row.ID,scope.row.dwlsgxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dwlsgxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdflgllbdm  BeQuickEdit -->  
        <el-table-column label="服务点分类管理类别代码" prop="fwdflgllbdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdflgllbdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdflgllbdm',scope.row.ID,scope.row.fwdflgllbdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdflgllbdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdfldm  BeQuickEdit -->  
        <el-table-column label="服务点分类代码" prop="fwdfldm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdfldm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdfldm',scope.row.ID,scope.row.fwdfldm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdfldm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--jjlxdm  BeQuickEdit -->  
        <el-table-column label="经济类型代码" prop="jjlxdm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.jjlxdm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('jjlxdm',scope.row.ID,scope.row.jjlxdm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.jjlxdm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--dz  BeQuickEdit -->  
        <el-table-column label="地址" prop="dz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.dz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('dz',scope.row.ID,scope.row.dz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.dz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdyyjb  BeQuickEdit -->  
        <el-table-column label="服务点医院级别" prop="fwdyyjb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdyyjb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdyyjb',scope.row.ID,scope.row.fwdyyjb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdyyjb}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdyydj  BeQuickEdit -->  
        <el-table-column label="服务点医院等级" prop="fwdyydj" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdyydj"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdyydj',scope.row.ID,scope.row.fwdyydj,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdyydj}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkzhm  BeQuickEdit -->  
        <el-table-column label="许可证号码" prop="xkzhm" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkzhm"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkzhm',scope.row.ID,scope.row.xkzhm,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkzhm}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkxmmc  BeQuickEdit -->  
        <el-table-column label="许可项目名称" prop="xkxmmc" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkxmmc"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkxmmc',scope.row.ID,scope.row.xkxmmc,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkxmmc}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--xkzyxq  BeQuickEdit -->  
        <el-table-column label="许可证有效期" prop="xkzyxq" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.xkzyxq"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('xkzyxq',scope.row.ID,scope.row.xkzyxq,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.xkzyxq}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--kbzjjes  BeQuickEdit -->  
        <el-table-column label="开办资金金额数" prop="kbzjjes" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.kbzjjes"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('kbzjjes',scope.row.ID,scope.row.kbzjjes,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.kbzjjes}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--frdb  BeQuickEdit -->  
        <el-table-column label="法人代表" prop="frdb" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.frdb"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('frdb',scope.row.ID,scope.row.frdb,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.frdb}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--fwdszdmzzzdfbz  BeQuickEdit -->  
        <el-table-column label="服务点地方标志" prop="fwdszdmzzzdfbz" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.fwdszdmzzzdfbz"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('fwdszdmzzzdfbz',scope.row.ID,scope.row.fwdszdmzzzdfbz,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.fwdszdmzzzdfbz}} </div>
              </template>
            </el-popover>
        </template>
          </el-table-column>
          <!--sffzjg  BeQuickEdit -->  
        <el-table-column label="是否分支机构" prop="sffzjg" width="120"   sortable="custom" >
        <template #default="scope">
            <el-popover trigger="click" placement="top"  width = "280">  
            <el-row :gutter="10">
              <el-col :span="16">  <el-input type="textarea" autosize placeholder="请输入内容" v-model="scope.row.sffzjg"></el-input></el-col>
              <el-col :span="4"> <el-button size="small" type="primary" icon="el-icon-edit" class="table-button" @click="quickEdit_do('sffzjg',scope.row.ID,scope.row.sffzjg,scope)">保存</el-button> </el-col> 
            </el-row>  
              <template #reference>
                <div  class="quickEditTxt"  > {{scope.row.sffzjg}} </div>
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
  createBbPiServicePoint, 
  deleteBbPiServicePointByIds,
  updateBbPiServicePoint,
  findBbPiServicePoint,
  getBbPiServicePointList,
  quickEdit,
  excelList
} from '@/api/bbPiServicePoint' //  此处请自行替换地址
import { formatTimeToStr } from '@/utils/date'
import infoList from '@/mixins/infoList' 
import tinymce from '@/mixins/tinymce' 
import editForm from '@/mixins/editForm' 
export default {
  name: 'BbPiServicePoint',
  mixins: [infoList,tinymce,editForm], 
  data() {
    return {
      beNewWindow:true,//是否在新窗口打开编辑器
      listApi: getBbPiServicePointList,   
      excelListApi: excelList,
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
        //this.deleteBbPiServicePoint(row)
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
     const res = await deleteBbPiServicePointByIds({ ids })
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
			this.$router.push({ name:'bbPiServicePointForm', params:{id:id}})
		  } else {
			 this.$router.push({ name:'bbPiServicePointForm',params:{id:id}})
		  }
	  }else
	  {
		 if (id >0) {
			  const res = await findBbPiServicePoint({ID:id})
			  //console.log(res.data)
			  this.editType = 'update'
			  if (res.code === 200) 
			     this.formData = res.data.bbPiServicePoint 
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
          res = await createBbPiServicePoint(this.formData);
          break
        case "update": 
          res = await updateBbPiServicePoint(this.formData);
          break
        default: 
          res = await createBbPiServicePoint(this.formData);
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
